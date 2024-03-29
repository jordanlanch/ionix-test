package repository

import (
	"context"

	"github.com/jordanlanch/ionix-test/domain"
	"gorm.io/gorm"
)

type drugRepository struct {
	db    *gorm.DB
}

// NewDrugRepository crea una nueva instancia de DrugRepository
func NewDrugRepository(db *gorm.DB) domain.DrugRepository {
	return &drugRepository{db}
}

// Create añade una nueva instancia de Drug a la base de datos
func (r *drugRepository) Create(ctx context.Context, drug *domain.Drug) error {
	result := r.db.WithContext(ctx).Create(drug)
	return result.Error
}

// Update modifica una instancia existente de Drug en la base de datos
func (r *drugRepository) Update(ctx context.Context, id int, drug *domain.Drug) error {
	result := r.db.WithContext(ctx).Model(&domain.Drug{}).Where("id = ?", id).Updates(drug)
	return result.Error
}

// FindByID busca una instancia de Drug por su ID
func (r *drugRepository) FindByID(ctx context.Context, id int) (*domain.Drug, error) {
	var drug domain.Drug
	result := r.db.WithContext(ctx).First(&drug, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &drug, nil
}

// FindAll retorna todas las instancias de Drug de la base de datos con paginación
func (r *drugRepository) FindAll(ctx context.Context, pagination *domain.Pagination) ([]*domain.Drug, error) {
	var drugs []*domain.Drug
	db := r.db.WithContext(ctx)
	db = applyPagination(db, pagination)

	if err := db.Find(&drugs).Error; err != nil {
		return nil, err
	}
	return drugs, nil
}

// Delete elimina una instancia de Drug de la base de datos
func (r *drugRepository) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&domain.Drug{}, id)
	return result.Error
}
