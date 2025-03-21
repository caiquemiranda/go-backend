package repositories

import (
	"context"

	"app15/internal/domain"
)

// PostRepository define a interface para operações de persistência de posts
type PostRepository interface {
	// Create cria um novo post no repositório
	Create(ctx context.Context, post *domain.Post) error
	
	// GetByID busca um post pelo seu ID
	GetByID(ctx context.Context, id string) (*domain.Post, error)
	
	// Update atualiza os dados de um post existente
	Update(ctx context.Context, post *domain.Post) error
	
	// Delete remove um post do repositório
	Delete(ctx context.Context, id string) error
	
	// List retorna uma lista paginada de posts
	List(ctx context.Context, page, pageSize int) ([]*domain.Post, error)
	
	// ListByAuthor retorna uma lista paginada de posts de um autor específico
	ListByAuthor(ctx context.Context, authorID string, page, pageSize int) ([]*domain.Post, error)
} 