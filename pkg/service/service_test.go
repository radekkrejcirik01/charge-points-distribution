package service

import (
	"testing"

	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestGetChargePointsByGroupId(t *testing.T) {
	chargePoints := []database.ChargePoint{
		{Id: 1, GroupId: 2},
		{Id: 2, GroupId: 2},
		{Id: 3, GroupId: 2},
	}

	tests := []struct {
		name     string
		groupId  uint
		provided []database.ChargePoint
		expected []database.ChargePoint
	}{
		{
			name:     "No charge points provided",
			groupId:  1,
			provided: []database.ChargePoint{},
			expected: []database.ChargePoint{},
		},
		{
			name:     "No charge points provided for group id 1",
			groupId:  1,
			provided: chargePoints,
			expected: []database.ChargePoint{},
		},
		{
			name:     "Charge points happy path",
			groupId:  2,
			provided: chargePoints,
			expected: chargePoints,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, getChargePointsByGroupId(test.groupId, test.provided))
		})
	}
}

func TestHasChargePointChargingStatus(t *testing.T) {
	chargePointConnectors := []database.ChargePointConnector{
		{ChargePointId: 1, Status: availableStatus},
		{ChargePointId: 1, Status: availableStatus},
		{ChargePointId: 2, Status: chargingStatus},
		{ChargePointId: 2, Status: availableStatus},
		{ChargePointId: 3, Status: chargingStatus},
		{ChargePointId: 3, Status: chargingStatus},
	}

	tests := []struct {
		name          string
		chargePointId uint
		provided      []database.ChargePointConnector
		expected      bool
	}{
		{
			name:          "Has no charging status",
			chargePointId: 1,
			provided:      chargePointConnectors,
			expected:      false,
		},
		{
			name:          "Has one charging status",
			chargePointId: 2,
			provided:      chargePointConnectors,
			expected:      true,
		},
		{
			name:          "Has mutliple charging status",
			chargePointId: 3,
			provided:      chargePointConnectors,
			expected:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, hasChargePointChargingStatus(test.chargePointId, test.provided))
		})
	}
}

func TestDistributeCurrent(t *testing.T) {
	availableChPoints := []database.ChargePoint{
		{Id: 1, Priority: 2},
		{Id: 2, Priority: 3},
		{Id: 3, Priority: 1}}

	tests := []struct {
		name       string
		maxCurrent float32
		provided   []database.ChargePoint
		expected   []Output
	}{
		{
			name:       "Pass zero as maximal current",
			maxCurrent: 0,
			provided:   availableChPoints,
			expected: []Output{
				{ChargePointId: 1, Current: 0},
				{ChargePointId: 2, Current: 0},
				{ChargePointId: 3, Current: 0},
			},
		},
		{
			name:       "Pass decimal number as maximal current",
			maxCurrent: 12.75,
			provided:   availableChPoints,
			expected: []Output{
				{ChargePointId: 1, Current: 4.25},
				{ChargePointId: 2, Current: 6.375},
				{ChargePointId: 3, Current: 2.125},
			},
		},
		{
			name:       "Pass whole number as maximal current",
			maxCurrent: 24,
			provided:   availableChPoints,
			expected: []Output{
				{ChargePointId: 1, Current: 8},
				{ChargePointId: 2, Current: 12},
				{ChargePointId: 3, Current: 4},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, distributeCurrent(test.provided, test.maxCurrent))
		})
	}
}
