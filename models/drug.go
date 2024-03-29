package models

import (
	"time"

	"gorm.io/gorm"
)

type DrugModel struct {
	gorm.Model
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255);"`
	Approved    bool      `gorm:"type:boolean;"`
	MinDose     int       `gorm:"type:integer;"`
	MaxDose     int       `gorm:"type:integer;"`
	AvailableAt time.Time `gorm:"type:timestamp;"`
}
