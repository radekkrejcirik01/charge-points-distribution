package service

import (
	"sort"

	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"gorm.io/gorm"
)

const chargingStatus = "Charging"
const availableStatus = "Available"

type Output struct {
	ChargePointId uint    `json:"chargePointId"`
	Current       float32 `json:"current"`
}

func GetOutput(db *gorm.DB) ([]Output, error) {
	// Get all groups from DB
	groups, err := database.GetGroups(db)
	if err != nil {
		return []Output{}, err
	}

	// Get all charge points from DB
	chPoints, err := database.GetChargePoints(db)
	if err != nil {
		return []Output{}, err
	}

	// Get all charge points connectors from DB
	chPointsConnectors, err := database.GetChargePointsConnectors(db)
	if err != nil {
		return []Output{}, err
	}

	return getOutput(groups, chPoints, chPointsConnectors)
}

// Get list of maps with charge point id and it's current
func getOutput(
	groups []database.Group,
	chargePoints []database.ChargePoint,
	chargePointsConnectors []database.ChargePointConnector,
) ([]Output, error) {
	output := make([]Output, 0)

	for _, group := range groups {
		availableChPoints := make([]database.ChargePoint, 0)

		// Get all charge points per group
		chPoints := getChargePointsByGroupId(group.Id, chargePoints)
		for _, chPoint := range chPoints {
			// Check if charge point has at least one charging status
			if hasChargePointChargingStatus(chPoint.Id, chargePointsConnectors) {
				availableChPoints = append(availableChPoints, chPoint)
			} else {
				// Allocate current 0 if there is no connector with charging status
				output = append(output, Output{
					ChargePointId: chPoint.Id,
					Current:       0,
				})
			}
		}

		// Distribute the maximal current between available charge points
		currents := distributeCurrent(
			availableChPoints,
			group.MaxCurrent,
		)

		output = append(output, currents...)
	}

	// Order output values by charge point id
	sort.SliceStable(output, func(i, y int) bool {
		return output[i].ChargePointId < output[y].ChargePointId
	})

	return output, nil
}

// Get charge points per group id
func getChargePointsByGroupId(
	groupId uint,
	chargePoints []database.ChargePoint,
) []database.ChargePoint {
	result := make([]database.ChargePoint, 0)

	for _, chargePoint := range chargePoints {
		if chargePoint.GroupId == groupId {
			result = append(result, chargePoint)
		}
	}

	return result
}

// Check if charge point has at least one charging status
func hasChargePointChargingStatus(
	chargePointId uint,
	chargePointsConnectors []database.ChargePointConnector,
) bool {
	for _, connector := range chargePointsConnectors {
		if connector.ChargePointId == chargePointId && connector.Status == chargingStatus {
			return true
		}
	}

	return false
}

// Distribute the maximal current between available charge points based on priority
func distributeCurrent(
	availableChPoints []database.ChargePoint,
	maxCurrent float32,
) []Output {
	output := make([]Output, 0)

	// Get summary of all priorities
	var prioritySum int
	for _, chPoint := range availableChPoints {
		prioritySum += chPoint.Priority
	}

	// Get one unit by dividing the maximal current by summary of priorities
	currentUnit := maxCurrent / float32(prioritySum)

	for _, chPoint := range availableChPoints {
		// Calculate the current, multiply one current unit by priority
		current := currentUnit * float32(chPoint.Priority)

		output = append(output, Output{
			ChargePointId: chPoint.Id,
			Current:       current,
		})
	}

	return output
}
