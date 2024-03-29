package repository

import (
	"context"
	"testing"

	"github.com/jordanlanch/ionix-test/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := NewUserRepository(db)

	user := &domain.User{
		Name:  nil, // Opcional, puedes omitirlo o poner un valor
		Email: "test@test.com",
	}

	createdUser, err := repo.Create(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser.ID)
	assert.Equal(t, user.Email, createdUser.Email)
}

func TestFindByID(t *testing.T) {
	repo := NewUserRepository(db)

	// Primero, creamos un usuario para tener un ID válido para buscar.
	user := &domain.User{
		Name:  nil, // Opcional
		Email: "test@test.com",
	}
	db.Create(user) // Asumimos que Create ajusta el ID automáticamente.

	// Ahora podemos buscar por el ID del usuario recién creado.
	foundUser, err := repo.FindByID(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, foundUser.ID)

	// Intentamos buscar un usuario con un ID que probablemente no exista.
	_, err = repo.FindByID(context.Background(), -1) // Usamos un ID improbable.
	assert.Error(t, err)
}

func TestFindByEmail(t *testing.T) {
	repo := NewUserRepository(db)

	email := "test@test.com"
	user := &domain.User{
		Email: email,
		Name:  nil, // Opcional
	}
	db.Create(user)

	foundUser, err := repo.FindByEmail(context.Background(), email)
	assert.NoError(t, err)
	assert.Equal(t, email, foundUser.Email)

	// Intentamos buscar un usuario con un email que no existe.
	_, err = repo.FindByEmail(context.Background(), "noexiste@test.com")
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	repo := NewUserRepository(db)

	// Creamos un usuario primero para tener algo que actualizar.
	user := &domain.User{
		Email: "test@test.com",
		Name:  nil, // Opcional
	}
	db.Create(user)

	// Actualizamos el correo electrónico del usuario.
	user.Email = "updated@test.com"
	updatedUser, err := repo.Update(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, "updated@test.com", updatedUser.Email)
}
