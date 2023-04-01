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
		availableChPointIds := make([]uint, 0)

		// Get all charge points per group
		chPoints := getChargePointsByGroupId(group.Id, chargePoints)
		for _, chPoint := range chPoints {
			// Check if charge point has at least one charging status
			if hasChargePointChargingStatus(chPoint.Id, chargePointsConnectors) {
				availableChPointIds = append(availableChPointIds, chPoint.Id)
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
			availableChPointIds,
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

// Distribute the maximal current between available charge points
func distributeCurrent(availableChPointIds []uint, maxCurrent float32) []Output {
	output := make([]Output, 0)

	// Devide the maxinal current by number of charge points
	chargePointsNumber := len(availableChPointIds)
	current := maxCurrent / float32(chargePointsNumber)

	for _, chargePointId := range availableChPointIds {
		output = append(output, Output{
			ChargePointId: chargePointId,
			Current:       current,
		})
	}

	return output
}
