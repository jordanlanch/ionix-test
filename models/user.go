package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID       int     `gorm:"primaryKey;autoIncrement"`
	Name     *string `gorm:"type:varchar(255);"`
	Email    string  `gorm:"type:varchar(255);unique"`
	Password string  `gorm:"type:varchar(255);"`
}
