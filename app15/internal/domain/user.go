package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User representa a entidade de usuário no domínio
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // O campo senha não será serializado para JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser cria uma nova instância de User
func NewUser(username, email, password string) (*User, error) {
	if username == "" {
		return nil, errors.New("username não pode ser vazio")
	}
	
	if email == "" {
		return nil, errors.New("email não pode ser vazio")
	}
	
	if password == "" {
		return nil, errors.New("senha não pode ser vazia")
	}
	
	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	
	now := time.Now()
	
	return &User{
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// ValidatePassword verifica se a senha fornecida corresponde à senha armazenada
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// UpdatePassword atualiza a senha do usuário
func (u *User) UpdatePassword(password string) error {
	if password == "" {
		return errors.New("nova senha não pode ser vazia")
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	u.Password = string(hashedPassword)
	u.UpdatedAt = time.Now()
	return nil
}

// UpdateEmail atualiza o email do usuário
func (u *User) UpdateEmail(email string) error {
	if email == "" {
		return errors.New("email não pode ser vazio")
	}
	
	u.Email = email
	u.UpdatedAt = time.Now()
	return nil
}

// UpdateUsername atualiza o nome de usuário
func (u *User) UpdateUsername(username string) error {
	if username == "" {
		return errors.New("username não pode ser vazio")
	}
	
	u.Username = username
	u.UpdatedAt = time.Now()
	return nil
} 