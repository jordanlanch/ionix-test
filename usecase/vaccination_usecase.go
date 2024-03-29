package usecase

import (
	"context"
	"time"

	"github.com/jordanlanch/ionix-test/domain"
)

type vaccinationUsecase struct {
	vaccinationRepo domain.VaccinationRepository
	timeout         time.Duration
}

func NewVaccinationUsecase(vr domain.VaccinationRepository, timeout time.Duration) domain.VaccinationUsecase {
	return &vaccinationUsecase{
		vaccinationRepo: vr,
		timeout:         timeout,
	}
}

func (vu *vaccinationUsecase) CreateVaccination(ctx context.Context, vac *domain.Vaccination) error {
	ctx, cancel := context.WithTimeout(ctx, vu.timeout)
	defer cancel()
	return vu.vaccinationRepo.Create(ctx, vac)
}

func (vu *vaccinationUsecase) UpdateVaccination(ctx context.Context, id int, vac *domain.Vaccination) error {
	ctx, cancel := context.WithTimeout(ctx, vu.timeout)
	defer cancel()
	return vu.vaccinationRepo.Update(ctx, id, vac)
}

func (vu *vaccinationUsecase) GetVaccinationById(ctx context.Context, id int) (*domain.Vaccination, error) {
	ctx, cancel := context.WithTimeout(ctx, vu.timeout)
	defer cancel()
	return vu.vaccinationRepo.FindByID(ctx, id)
}

func (vu *vaccinationUsecase) GetAllVaccinations(ctx context.Context, pagination *domain.Pagination) ([]*domain.Vaccination, error) {
	ctx, cancel := context.WithTimeout(ctx, vu.timeout)
	defer cancel()
	return vu.vaccinationRepo.FindAll(ctx, pagination)
}

func (vu *vaccinationUsecase) DeleteVaccination(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, vu.timeout)
	defer cancel()
	return vu.vaccinationRepo.Delete(ctx, id)
}
