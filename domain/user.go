package domain

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password,omitempty"`
}

// SetPassword encripta la contraseña y la asigna al usuario.
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ComparePassword compara la contraseña proporcionada con la contraseña almacenada del usuario.
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// UserRepository define los métodos necesarios para interactuar con la base de datos de registros de usuarios.
type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
}

// UserUseCase define los métodos necesarios para interactuar con la lógica de negocio de registros de usuarios.
type UserUseCase interface {
	GetByID(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}
