package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jordanlanch/ionix-test/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateDrug(t *testing.T) {
	repo := NewDrugRepository(db)

	drug := &domain.Drug{
		Name:        "Paracetamol",
		Approved:    true,
		MinDose:     500,
		MaxDose:     1000,
		AvailableAt: time.Now(),
	}

	err := repo.Create(context.Background(), drug)
	assert.NoError(t, err)
	assert.NotZero(t, drug.ID) // Asegurarse de que se haya generado un ID.
}

func TestUpdateDrug(t *testing.T) {
	repo := NewDrugRepository(db)

	// Crear un registro de medicamento inicial para actualizarlo después.
	drug := domain.Drug{
		Name:        "Ibuprofeno",
		Approved:    true,
		MinDose:     200,
		MaxDose:     800,
		AvailableAt: time.Now(),
	}
	err := repo.Create(context.Background(), &drug)
	assert.NoError(t, err)

	// Actualizar el registro de medicamento.
	drug.Name = "Ibuprofeno Actualizado"
	drug.MaxDose = 1200
	err = repo.Update(context.Background(), drug.ID, &drug)
	assert.NoError(t, err)

	// Verificar que la actualización fue exitosa.
	updatedDrug, err := repo.FindByID(context.Background(), drug.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Ibuprofeno Actualizado", updatedDrug.Name)
	assert.Equal(t, 1200, updatedDrug.MaxDose)
}

func TestFindAllDrugs(t *testing.T) {
	repo := NewDrugRepository(db)

	// Crear registros de medicamentos para probar la paginación.
	for i := 0; i < 5; i++ {
		drug := domain.Drug{
			Name:        fmt.Sprintf("Medicamento %d", i),
			Approved:    i%2 == 0,
			MinDose:     100 + i*100,
			MaxDose:     500 + i*100,
			AvailableAt: time.Now(),
		}
		err := repo.Create(context.Background(), &drug)
		assert.NoError(t, err)
	}

	// Probar la obtención de todos los registros con paginación.
	pagination := domain.Pagination{Limit: IntPointer(2), Offset: IntPointer(0)}
	drugs, err := repo.FindAll(context.Background(), &pagination)
	assert.NoError(t, err)
	assert.Len(t, drugs, 2) // Verificar que solo se obtienen 2 registros debido a la paginación.
}

func TestDeleteDrug(t *testing.T) {
	repo := NewDrugRepository(db)

	// Crear un registro de medicamento para eliminarlo.
	drug := domain.Drug{
		Name:        "Aspirina",
		Approved:    true,
		MinDose:     300,
		MaxDose:     600,
		AvailableAt: time.Now(),
	}
	err := repo.Create(context.Background(), &drug)
	assert.NoError(t, err)

	// Eliminar el registro de medicamento creado.
	err = repo.Delete(context.Background(), drug.ID)
	assert.NoError(t, err)

	// Verificar que el registro haya sido eliminado.
	_, err = repo.FindByID(context.Background(), drug.ID)
	assert.Error(t, err) // Se espera un error al intentar encontrar un registro eliminado.
}
