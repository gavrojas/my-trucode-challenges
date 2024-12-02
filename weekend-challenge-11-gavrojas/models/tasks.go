package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name          string
	Description   string
	EstimateHours int
	ProjectID     uint
	Project       Project `gorm:"constraint:OnDelete:CASCADE;"`
}
