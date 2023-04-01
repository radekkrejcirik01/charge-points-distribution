package statuses

import (
	"fmt"

	"gorm.io/gorm"
)

// Get all groups records from groups table
func getGroups(db *gorm.DB) ([]Group, error) {
	var groups []Group
	if err := db.Table("groups").Find(&groups).Error; err != nil {
		return []Group{}, err
	}

	// If length of groups is 0 return error message
	if len(groups) == 0 {
		return groups, fmt.Errorf("no groups found 😔")
	}

	return groups, nil
}

// Get all charge points records from charge_points table
func getChargePoints(db *gorm.DB) ([]ChargePoint, error) {
	var chargePoints []ChargePoint
	if err := db.Table("charge_points").Find(&chargePoints).Error; err != nil {
		return []ChargePoint{}, err
	}

	// If length of charge points is 0 return error message
	if len(chargePoints) == 0 {
		return chargePoints, fmt.Errorf("no charge points found 😔")
	}

	return chargePoints, nil
}

// Get all charge points connectors records from charge_points_connectors table
func getChargePointsConnectors(db *gorm.DB) ([]ChargePointConnector, error) {
	var chargePointConnectors []ChargePointConnector
	if err := db.Table("charge_point_connectors").Find(&chargePointConnectors).Error; err != nil {
		return []ChargePointConnector{}, err
	}

	// If length of charge point connectors is 0 return error message
	if len(chargePointConnectors) == 0 {
		return chargePointConnectors, fmt.Errorf("no connectors for charge points found 😔")
	}

	return chargePointConnectors, nil
}
