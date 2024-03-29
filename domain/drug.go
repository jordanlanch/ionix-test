package domain

import (
	"context"
	"time"
)

type Drug struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Approved    bool      `json:"approved"`
	MinDose     int       `json:"min_dose"`
	MaxDose     int       `json:"max_dose"`
	AvailableAt time.Time `json:"available_at"`
}

type DrugRepository interface {
	Create(ctx context.Context, drug *Drug) error
	Update(ctx context.Context, id int, drug *Drug) error
	FindByID(ctx context.Context, id int) (*Drug, error)
	FindAll(ctx context.Context, pagination *Pagination) ([]*Drug, error)
	Delete(ctx context.Context, id int) error
}

type DrugUsecase interface {
	CreateDrug(ctx context.Context, drug *Drug) error
	UpdateDrug(ctx context.Context, id int, drug *Drug) error
	GetDrugById(ctx context.Context, id int) (*Drug, error)
	GetAllDrugs(ctx context.Context, pagination *Pagination) ([]*Drug, error)
	DeleteDrug(ctx context.Context, id int) error
}
