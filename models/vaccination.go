package models

import (
	"time"

	"gorm.io/gorm"
)

type VaccinationModel struct {
	gorm.Model
	ID     int       `gorm:"primaryKey;autoIncrement"`
	Name   string    `gorm:"type:varchar(255);"`
	DrugID int       `gorm:"type:integer;not null"`
	Drug   DrugModel `gorm:"foreignKey:DrugID"`
	Dose   int       `gorm:"type:integer;"`
	Date   time.Time `gorm:"type:timestamp;"`
}
