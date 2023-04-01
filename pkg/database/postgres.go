package database

import (
	"fmt"

	"gorm.io/gorm"
)

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

// Get all groups records from groups table
func GetGroups(d *gorm.DB) ([]Group, error) {
	var groups []Group
	if err := d.Table("groups").Find(&groups).Error; err != nil {
		return []Group{}, err
	}

	// If length of groups is 0 return error message
	if len(groups) == 0 {
		return groups, fmt.Errorf("no groups found 😔")
	}

	return groups, nil
}

// Get all charge points records from charge_points table
func GetChargePoints(d *gorm.DB) ([]ChargePoint, error) {
	var chargePoints []ChargePoint
	if err := d.Table("charge_points").Find(&chargePoints).Error; err != nil {
		return []ChargePoint{}, err
	}

	// If length of charge points is 0 return error message
	if len(chargePoints) == 0 {
		return chargePoints, fmt.Errorf("no charge points found 😔")
	}

	return chargePoints, nil
}

// Get all charge points connectors records from charge_points_connectors table
func GetChargePointsConnectors(d *gorm.DB) ([]ChargePointConnector, error) {
	var chargePointConnectors []ChargePointConnector
	if err := d.Table("charge_point_connectors").Find(&chargePointConnectors).Error; err != nil {
		return []ChargePointConnector{}, err
	}

	// If length of charge point connectors is 0 return error message
	if len(chargePointConnectors) == 0 {
		return chargePointConnectors, fmt.Errorf("no connectors for charge points found 😔")
	}

	return chargePointConnectors, nil
}
