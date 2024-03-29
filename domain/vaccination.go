package domain

import (
	"context"
	"time"
)


type Vaccination struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	DrugID   int       `json:"drug_id"`
	Dose     int       `json:"dose"`
	Date     time.Time `json:"date"`
}

type VaccinationRepository interface {
	Create(ctx context.Context, vac *Vaccination) error
	Update(ctx context.Context, id int, vac *Vaccination) error
	FindByID(ctx context.Context, id int) (*Vaccination, error)
	FindAll(ctx context.Context, pagination *Pagination) ([]*Vaccination, error)
	Delete(ctx context.Context, id int) error
}

type VaccinationUsecase interface {
	CreateVaccination(ctx context.Context, vac *Vaccination) error
	UpdateVaccination(ctx context.Context, id int, vac *Vaccination) error
	GetVaccinationById(ctx context.Context, id int) (*Vaccination, error)
	GetAllVaccinations(ctx context.Context, pagination *Pagination) ([]*Vaccination, error)
	DeleteVaccination(ctx context.Context, id int) error
}
