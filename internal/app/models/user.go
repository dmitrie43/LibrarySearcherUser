package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"size:255"`
	Email         string `gorm:"type:varchar(100);unique_index"`
	Password      string `gorm:"size:255"`
	Role          Role   `gorm:"embedded"`
	Avatar        sql.NullString
	RememberToken sql.NullString
}
