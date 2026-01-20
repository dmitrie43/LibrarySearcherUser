package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"size:255"`
	Code              string `gorm:"size:255"`
	IsAllowAdminPanel bool
}
