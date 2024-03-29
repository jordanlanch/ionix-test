package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jordanlanch/ionix-test/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreateVaccination(t *testing.T) {
	repo := NewVaccinationRepository(db)

	vac := domain.Vaccination{
		Name:   "John Doe",
		DrugID: 1,
		Dose:   100,
		Date:   time.Now(),
	}

	err := repo.Create(context.Background(), &vac)
	assert.NoError(t, err)
	assert.NotZero(t, vac.ID) // Asegurarse de que se haya generado un ID
}

func TestFindByIDVaccination(t *testing.T) {
	repo := NewVaccinationRepository(db)

	// Primero, necesitas crear una vacunación para tener un ID válido para buscar.
	vac := domain.Vaccination{
		Name:   "John Doe",
		DrugID: 1,
		Dose:   100,
		Date:   time.Now(),
	}
	err := repo.Create(context.Background(), &vac)
	assert.NoError(t, err)

	// Ahora puedes buscar por el ID de la vacunación recién creada.
	foundVac, err := repo.FindByID(context.Background(), vac.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundVac)
	assert.Equal(t, vac.Name, foundVac.Name)
}

func TestUpdateVaccination(t *testing.T) {
	repo := NewVaccinationRepository(db)

	// Crear un registro de vacunación inicial para actualizarlo después.
	vac := domain.Vaccination{
		Name:   "John Doe",
		DrugID: 1,
		Dose:   100,
		Date:   time.Now(),
	}
	err := repo.Create(context.Background(), &vac)
	assert.NoError(t, err)

	// Actualizar el registro de vacunación.
	vac.Name = "Jane Doe"
	vac.Dose = 150
	err = repo.Update(context.Background(), vac.ID, &vac)
	assert.NoError(t, err)

	// Verificar que la actualización fue exitosa.
	updatedVac, err := repo.FindByID(context.Background(), vac.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", updatedVac.Name)
	assert.Equal(t, 150, updatedVac.Dose)
}

func TestFindAllVaccinations(t *testing.T) {
	repo := NewVaccinationRepository(db)

	// Crear registros de vacunación para probar la paginación.
	for i := 0; i < 5; i++ {
		vac := domain.Vaccination{
			Name:   fmt.Sprintf("Person %d", i),
			DrugID: 1,
			Dose:   100 + i,
			Date:   time.Now(),
		}
		err := repo.Create(context.Background(), &vac)
		assert.NoError(t, err)
	}

	// Probar la obtención de todos los registros con paginación.
	pagination := domain.Pagination{Limit: IntPointer(2), Offset: IntPointer(0)}
	vacs, err := repo.FindAll(context.Background(), &pagination)
	assert.NoError(t, err)
	assert.Len(t, vacs, 2) // Verificar que solo se obtienen 2 registros debido a la paginación.
}

// util.IntPointer es una función de ayuda para obtener un puntero a un int.
func IntPointer(i int) *int {
	return &i
}

func TestDeleteVaccination(t *testing.T) {
	repo := NewVaccinationRepository(db)

	// Crear un registro de vacunación para eliminarlo.
	vac := domain.Vaccination{
		Name:   "John Doe",
		DrugID: 1,
		Dose:   100,
		Date:   time.Now(),
	}
	err := repo.Create(context.Background(), &vac)
	assert.NoError(t, err)

	// Eliminar el registro de vacunación creado.
	err = repo.Delete(context.Background(), vac.ID)
	assert.NoError(t, err)

	// Verificar que el registro haya sido eliminado.
	_, err = repo.FindByID(context.Background(), vac.ID)
	assert.Error(t, err) // Se espera un error al intentar encontrar un registro eliminado.
}
