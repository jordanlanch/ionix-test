package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jordanlanch/ionix-test/domain"
	"github.com/jordanlanch/ionix-test/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateVaccination(t *testing.T) {
	mockVaccinationRepo := new(mocks.VaccinationRepository)
	vu := NewVaccinationUsecase(mockVaccinationRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
		mockVac := domain.Vaccination{Name: "COVID-19 Vaccine", DrugID: 1, Dose: 100, Date: time.Now()}
		mockVaccinationRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Vaccination")).Return(nil).Once()

		err := vu.CreateVaccination(context.Background(), &mockVac)

		assert.NoError(t, err)
		mockVaccinationRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockVac := domain.Vaccination{Name: "COVID-19 Vaccine", DrugID: 1, Dose: 100, Date: time.Now()}
		mockVaccinationRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Vaccination")).Return(errors.New("Unexpected error")).Once()

		err := vu.CreateVaccination(context.Background(), &mockVac)

		assert.Error(t, err)
		mockVaccinationRepo.AssertExpectations(t)
	})
}

func TestUpdateVaccination(t *testing.T) {
	mockVaccinationRepo := new(mocks.VaccinationRepository)
	vu := NewVaccinationUsecase(mockVaccinationRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
			mockVaccination := domain.Vaccination{Name: "Covid-19 Vaccine", DrugID: 1, Dose: 1, Date: time.Now()}
			mockVaccinationRepo.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.MatchedBy(func(vac *domain.Vaccination) bool {
				return vac.Name == "Covid-19 Vaccine" && vac.DrugID == 1
			})).Return(nil).Once()


			err := vu.UpdateVaccination(context.Background(), mockVaccination.ID, &mockVaccination)

			assert.NoError(t, err)
			mockVaccinationRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
			mockVaccination := domain.Vaccination{ID: 1, Name: "Covid-19 Vaccine", DrugID: 1, Dose: 2, Date: time.Now()}
			mockVaccinationRepo.On("Update", mock.Anything, mock.AnythingOfType("int"), &mockVaccination).Return(errors.New("Unexpected error")).Once()

			err := vu.UpdateVaccination(context.Background(), mockVaccination.ID, &mockVaccination)

			assert.Error(t, err)
			mockVaccinationRepo.AssertExpectations(t)
	})
}

func TestGetAllVaccinations(t *testing.T) {
	mockVaccinationRepo := new(mocks.VaccinationRepository)
	vu := NewVaccinationUsecase(mockVaccinationRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
			mockVaccinations := []*domain.Vaccination{
				{ID: 1, Name: "Covid-19 Vaccine", DrugID: 1, Dose: 1, Date: time.Now()},
				{ID: 2, Name: "Flu Vaccine", DrugID: 2, Dose: 1, Date: time.Now()},
			}
			limit := 2
			offset := 0
			pagination := domain.Pagination{Limit: &limit, Offset: &offset}
			mockVaccinationRepo.On("FindAll", mock.Anything, &pagination).Return(mockVaccinations, nil).Once()

			result, err := vu.GetAllVaccinations(context.Background(), &pagination)

			assert.NoError(t, err)
			assert.Len(t, result, 2)
			mockVaccinationRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
			limit := 2
			offset := 0
			pagination := domain.Pagination{Limit: &limit, Offset: &offset}
			mockVaccinationRepo.On("FindAll", mock.Anything, &pagination).Return(nil, errors.New("Unexpected error")).Once()

			result, err := vu.GetAllVaccinations(context.Background(), &pagination)

			assert.Error(t, err)
			assert.Nil(t, result)
			mockVaccinationRepo.AssertExpectations(t)
	})
}



func TestDeleteVaccination(t *testing.T) {
	mockVaccinationRepo := new(mocks.VaccinationRepository)
	vu := NewVaccinationUsecase(mockVaccinationRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
		mockVaccinationRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := vu.DeleteVaccination(context.Background(), 1)

		assert.NoError(t, err)
		mockVaccinationRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockVaccinationRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("Unexpected error")).Once()

		err := vu.DeleteVaccination(context.Background(), 1)

		assert.Error(t, err)
		mockVaccinationRepo.AssertExpectations(t)
	})
}
