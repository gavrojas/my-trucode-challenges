package models

import "gorm.io/gorm"

type Entry struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	UserID     uint   `json:"userId"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	User       User   `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}
