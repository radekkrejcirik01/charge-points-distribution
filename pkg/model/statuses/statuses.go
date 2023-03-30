package statuses

import (
	"gorm.io/gorm"
)

type Statuses struct {
	Id              uint
	ConnectorStatus string
}

func GetStatusById(db *gorm.DB, t *Statuses, id string) error {
	return db.Table("statuses").Where("id = ?", id).First(&t).Error
}
