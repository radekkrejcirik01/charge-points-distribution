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

	// If number of groups is 0 return error message
	if len(groups) == 0 {
		return groups, fmt.Errorf("no groups found 😔")
	}

	return groups, nil
}

// Get all charge points records from charge_points table
func GetChargePoints(d *gorm.DB) ([]ChargePoint, error) {
	var chPoints []ChargePoint
	if err := d.Table("charge_points").Find(&chPoints).Error; err != nil {
		return []ChargePoint{}, err
	}

	// If number of charge points is 0 return error message
	if len(chPoints) == 0 {
		return chPoints, fmt.Errorf("no charge points found 😔")
	}

	return chPoints, nil
}

// Get all charge points connectors records from charge_points_connectors table
func GetChargePointsConnectors(d *gorm.DB) ([]ChargePointConnector, error) {
	var chPointConnectors []ChargePointConnector
	if err := d.Table("charge_point_connectors").Find(&chPointConnectors).Error; err != nil {
		return []ChargePointConnector{}, err
	}

	// If number of charge point connectors is 0 return error message
	if len(chPointConnectors) == 0 {
		return chPointConnectors, fmt.Errorf("no connectors for charge points found 😔")
	}

	return chPointConnectors, nil
}

func CreateGroup(db *gorm.DB, maxCurrent float32) error {
	return db.Table("groups").Create(&Group{MaxCurrent: maxCurrent}).Error
}

func AddChargePoint(db *gorm.DB, groupId uint) error {
	return db.Table("charge_points").Create(&ChargePoint{GroupId: groupId}).Error
}

func AddChargePointConnector(db *gorm.DB, chPointId uint, status string) error {
	return db.Table("charge_point_connectors").Create(&ChargePointConnector{
		ChargePointId: chPointId,
		Status:        status,
	}).Error
}
