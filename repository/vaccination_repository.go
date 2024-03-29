package repository

import (
	"context"

	"github.com/jordanlanch/ionix-test/domain"
	"gorm.io/gorm"
)

type vaccinationRepository struct {
	db *gorm.DB
}

func NewVaccinationRepository(db *gorm.DB) domain.VaccinationRepository {
	return &vaccinationRepository{db}
}

func (r *vaccinationRepository) Create(ctx context.Context, vac *domain.Vaccination) error {
	return r.db.WithContext(ctx).Create(vac).Error
}

func (r *vaccinationRepository) Update(ctx context.Context, id int, vac *domain.Vaccination) error {
	return r.db.WithContext(ctx).Model(&domain.Vaccination{}).Where("id = ?", id).Updates(vac).Error
}

func (r *vaccinationRepository) FindByID(ctx context.Context, id int) (*domain.Vaccination, error) {
	var vac domain.Vaccination
	if err := r.db.WithContext(ctx).First(&vac, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &vac, nil
}

func (r *vaccinationRepository) FindAll(ctx context.Context, pagination *domain.Pagination) ([]*domain.Vaccination, error) {
	var vacs []*domain.Vaccination
	db := r.db.WithContext(ctx)
	db = applyPagination(db, pagination)

	if err := db.Find(&vacs).Error; err != nil {
		return nil, err
	}
	return vacs, nil
}

func (r *vaccinationRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&domain.Vaccination{}, id).Error
}

// applyPagination aplica la paginación a la consulta de base de datos.
func applyPagination(db *gorm.DB, pagination *domain.Pagination) *gorm.DB {
	if pagination != nil {
		if pagination.Limit != nil {
			db = db.Limit(*pagination.Limit)
		}
		if pagination.Offset != nil {
			db = db.Offset(*pagination.Offset)
		}
	}
	return db
}
