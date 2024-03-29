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

func TestCreateDrug(t *testing.T) {
	mockDrugRepo := new(mocks.DrugRepository)
	du := NewDrugUsecase(mockDrugRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
		mockDrug := domain.Drug{Name: "Aspirin", Approved: true}
		mockDrugRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Drug")).Return(nil).Once()

		err := du.CreateDrug(context.Background(), &mockDrug)

		assert.NoError(t, err)
		mockDrugRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDrug := domain.Drug{Name: "Aspirin", Approved: true}
		mockDrugRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Drug")).Return(errors.New("Unexpected error")).Once()

		err := du.CreateDrug(context.Background(), &mockDrug)

		assert.Error(t, err)
		mockDrugRepo.AssertExpectations(t)
	})
}
func TestUpdateDrug(t *testing.T) {
	mockDrugRepo := new(mocks.DrugRepository)
	du := NewDrugUsecase(mockDrugRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
		mockDrug := domain.Drug{ID: 1, Name: "Aspirin", Approved: true}
		// Nota: El ID no se incluye como argumento a actualizar, solo al final para identificar el registro
		mockDrugRepo.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.MatchedBy(func(drug *domain.Drug) bool {
			return drug.Name == "Aspirin" && drug.Approved == true
		})).Return(nil).Once()

		err := du.UpdateDrug(context.Background(), mockDrug.ID, &mockDrug)

		assert.NoError(t, err)
		mockDrugRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDrug := domain.Drug{ID: 1, Name: "Aspirin", Approved: true}
		mockDrugRepo.On("Update", mock.Anything, mock.AnythingOfType("int"), &mockDrug).Return(errors.New("Unexpected error")).Once()

		err := du.UpdateDrug(context.Background(), mockDrug.ID, &mockDrug)

		assert.Error(t, err)
		mockDrugRepo.AssertExpectations(t)
	})
}

func TestGetDrugById(t *testing.T) {
	mockDrugRepo := new(mocks.DrugRepository)
	du := NewDrugUsecase(mockDrugRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
		mockDrug := domain.Drug{ID: 1, Name: "Aspirin", Approved: true}
		mockDrugRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(&mockDrug, nil).Once()

		result, err := du.GetDrugById(context.Background(), mockDrug.ID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		mockDrugRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDrugRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(nil, errors.New("Not found")).Once()

		result, err := du.GetDrugById(context.Background(), 1)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockDrugRepo.AssertExpectations(t)
	})
}
func TestGetAllDrugs(t *testing.T) {
	mockDrugRepo := new(mocks.DrugRepository)
	du := NewDrugUsecase(mockDrugRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
		mockDrugs := []*domain.Drug{
			{ID: 1, Name: "Aspirin", Approved: true},
			{ID: 2, Name: "Ibuprofen", Approved: true},
		}
		limit := 2
		offset := 0
		pagination := domain.Pagination{Limit: &limit, Offset: &offset}
		mockDrugRepo.On("FindAll", mock.Anything, &pagination).Return(mockDrugs, nil).Once()

		result, err := du.GetAllDrugs(context.Background(), &pagination)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		mockDrugRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		limit := 2
		offset := 2
		pagination := domain.Pagination{Limit: &limit, Offset: &offset}
		mockDrugRepo.On("FindAll", mock.Anything, &pagination).Return(nil, errors.New("Unexpected error")).Once()

		result, err := du.GetAllDrugs(context.Background(), &pagination)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockDrugRepo.AssertExpectations(t)
	})
}
func TestDeleteDrug(t *testing.T) {
	mockDrugRepo := new(mocks.DrugRepository)
	du := NewDrugUsecase(mockDrugRepo, time.Second*2)

	t.Run("success", func(t *testing.T) {
		mockDrugRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := du.DeleteDrug(context.Background(), 1)

		assert.NoError(t, err)
		mockDrugRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDrugRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("Unexpected error")).Once()

		err := du.DeleteDrug(context.Background(), 1)

		assert.Error(t, err)
		mockDrugRepo.AssertExpectations(t)
	})
}
