package repositories

import (
	"context"

	"app15/internal/domain"
)

// UserRepository define a interface para operações de persistência de usuários
type UserRepository interface {
	// Create cria um novo usuário no repositório
	Create(ctx context.Context, user *domain.User) error
	
	// GetByID busca um usuário pelo seu ID
	GetByID(ctx context.Context, id string) (*domain.User, error)
	
	// GetByEmail busca um usuário pelo seu email
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	
	// GetByUsername busca um usuário pelo seu nome de usuário
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	
	// Update atualiza os dados de um usuário existente
	Update(ctx context.Context, user *domain.User) error
	
	// Delete remove um usuário do repositório
	Delete(ctx context.Context, id string) error
	
	// List retorna uma lista paginada de usuários
	List(ctx context.Context, page, pageSize int) ([]*domain.User, error)
} 