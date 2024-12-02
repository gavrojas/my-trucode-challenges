package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name         string
	DeliveryDate string
	Engineers    []User `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;"`
	Tasks        []Task `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;"`
}
