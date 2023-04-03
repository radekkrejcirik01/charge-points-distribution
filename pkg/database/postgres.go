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
	Id       uint
	GroupId  uint
	Priority int
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
		return groups, fmt.Errorf("no groups found")
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
		return chPoints, fmt.Errorf("no charge points found")
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
		return chPointConnectors, fmt.Errorf("no connectors for charge points found")
	}

	return chPointConnectors, nil
}

// Create a new group with maximal current value in groups table
func CreateGroup(db *gorm.DB, t *Group) error {
	return db.Table("groups").Create(t).Error
}

// Add a new charge point with group id and priority to charge_points table
func AddChargePoint(db *gorm.DB, t *ChargePoint) error {
	return db.Table("charge_points").Create(t).Error
}

// Get count of charge points in group from charge_points table
func GetGroupChargePointsCount(db *gorm.DB, t *ChargePoint) (int64, error) {
	var count int64
	if err := db.Table("charge_points").Where("group_id = ?", t.GroupId).Count(&count).Error; err != nil {
		return 0, err
	}

	// If count of charge points is 0 return error message
	if count == 0 {
		return 0, fmt.Errorf("no charge points found")
	}

	return count, nil
}

// Add a new charge point with group id and priority to charge_points table
func GetChargePointsByGroupId(db *gorm.DB, t *ChargePoint) error {
	return db.Table("charge_points").Create(t).Error
}

// Update priority of charge point by id in charge_points table
func UpdateChargePoint(db *gorm.DB, t *ChargePoint) error {
	return db.Table("charge_points").
		Where("id = ?", t.Id).
		Update("priority", t.Priority).
		Error
}

// Add a new charge point connector with charge point id and status to
// charge_point_connectors table
func AddChargePointConnector(db *gorm.DB, t *ChargePointConnector) error {
	return db.Table("charge_point_connectors").Create(t).Error
}

// Update status of charge point connector by id in charge_point_connectors table
func UpdateChargePointConnector(db *gorm.DB, t *ChargePointConnector) error {
	return db.Table("charge_point_connectors").
		Where("id = ?", t.Id).
		Update("status", t.Status).
		Error
}
