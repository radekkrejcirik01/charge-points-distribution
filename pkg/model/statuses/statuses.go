package statuses

import (
	"sort"

	"gorm.io/gorm"
)

const chargingStatus = "Charging"

type Group struct {
	Id         uint
	MaxCurrent float32
}

type ChargePoint struct {
	Id      uint
	GroupId uint
}

type ChargePointConnector struct {
	Id            uint
	ChargePointId uint
	Status        string
}

type Output struct {
	ChargePointId uint    `json:"chargePointId"`
	Current       float32 `json:"current"`
}

func GetOutput(db *gorm.DB) ([]Output, error) {
	// Get all groups from DB
	groups, err := getGroups(db)
	if err != nil {
		return []Output{}, err
	}

	// Get all charge points from DB
	chargePoints, err := getChargePoints(db)
	if err != nil || len(chargePoints) == 0 {
		return []Output{}, err
	}

	// Get all charge points connectors from DB
	chargePointsConnectors, err := getChargePointsConnectors(db)
	if err != nil {
		return []Output{}, err
	}

	return getOutput(groups, chargePoints, chargePointsConnectors)
}

// Get list of maps with charge point id and it's current
func getOutput(
	groups []Group,
	chargePoints []ChargePoint,
	chargePointsConnectors []ChargePointConnector,
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

	// Order output values by ChargePointId
	sort.SliceStable(output, func(i, y int) bool {
		return output[i].ChargePointId < output[y].ChargePointId
	})

	return output, nil
}

// Get charge points from one group by id
func getChargePointsByGroupId(groupId uint, chargePoints []ChargePoint) []ChargePoint {
	result := make([]ChargePoint, 0)

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
	chargePointsConnectors []ChargePointConnector,
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

	// Devide the maxinal currebt by number of charge points
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
