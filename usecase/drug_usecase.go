package usecase

import (
	"context"
	"time"

	"github.com/jordanlanch/ionix-test/domain"
)

type DrugUsecase struct {
	drugRepo       domain.DrugRepository
	contextTimeout time.Duration
}

func NewDrugUsecase(drugRepo domain.DrugRepository, timeout time.Duration) domain.DrugUsecase {
	return &DrugUsecase{
		drugRepo:       drugRepo,
		contextTimeout: timeout,
	}
}

func (du *DrugUsecase) CreateDrug(ctx context.Context, drug *domain.Drug) error {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	return du.drugRepo.Create(ctx, drug)
}

func (du *DrugUsecase) UpdateDrug(ctx context.Context, id int, drug *domain.Drug) error {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	return du.drugRepo.Update(ctx, id, drug)
}

func (du *DrugUsecase) GetDrugById(ctx context.Context, id int) (*domain.Drug, error) {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	return du.drugRepo.FindByID(ctx, id)
}

func (du *DrugUsecase) GetAllDrugs(ctx context.Context, pagination *domain.Pagination) ([]*domain.Drug, error) {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	return du.drugRepo.FindAll(ctx, pagination)
}

func (du *DrugUsecase) DeleteDrug(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	return du.drugRepo.Delete(ctx, id)
}
