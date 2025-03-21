package repositories

import (
	"context"

	"app15/internal/domain"
)

// CommentRepository define a interface para operações de persistência de comentários
type CommentRepository interface {
	// Create cria um novo comentário no repositório
	Create(ctx context.Context, comment *domain.Comment) error
	
	// GetByID busca um comentário pelo seu ID
	GetByID(ctx context.Context, id string) (*domain.Comment, error)
	
	// Update atualiza os dados de um comentário existente
	Update(ctx context.Context, comment *domain.Comment) error
	
	// Delete remove um comentário do repositório
	Delete(ctx context.Context, id string) error
	
	// ListByPost retorna uma lista paginada de comentários de um post específico
	ListByPost(ctx context.Context, postID string, page, pageSize int) ([]*domain.Comment, error)
	
	// ListByAuthor retorna uma lista paginada de comentários de um autor específico
	ListByAuthor(ctx context.Context, authorID string, page, pageSize int) ([]*domain.Comment, error)
} 