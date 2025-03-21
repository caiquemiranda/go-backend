package application

import (
	"context"
	"errors"

	"app15/internal/domain"
	"app15/internal/ports/repositories"

	"github.com/google/uuid"
)

// Errors específicos do serviço de comentários
var (
	ErrCommentNotFound    = errors.New("comentário não encontrado")
	ErrInvalidCommentData = errors.New("dados do comentário inválidos")
)

// CommentService representa o serviço de comentários da aplicação
type CommentService struct {
	commentRepo repositories.CommentRepository
	postRepo    repositories.PostRepository
	userRepo    repositories.UserRepository
}

// NewCommentService cria uma nova instância do serviço de comentários
func NewCommentService(
	commentRepo repositories.CommentRepository,
	postRepo repositories.PostRepository,
	userRepo repositories.UserRepository,
) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
		userRepo:    userRepo,
	}
}

// CommentRequest representa a estrutura de dados para criação/atualização de comentários
type CommentRequest struct {
	Content string `json:"content"`
}

// CreateComment cria um novo comentário
func (s *CommentService) CreateComment(ctx context.Context, postID string, req CommentRequest, authorID string) (*domain.Comment, error) {
	// Verificar se o post existe
	_, err := s.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, ErrPostNotFound
	}

	// Verificar se o autor existe
	_, err = s.userRepo.GetByID(ctx, authorID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Criar novo comentário
	comment, err := domain.NewComment(req.Content, postID, authorID)
	if err != nil {
		return nil, ErrInvalidCommentData
	}

	// Gerar ID único para o comentário
	comment.ID = uuid.New().String()

	// Salvar comentário no repositório
	err = s.commentRepo.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentByID busca um comentário pelo ID
func (s *CommentService) GetCommentByID(ctx context.Context, id string) (*domain.Comment, error) {
	comment, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrCommentNotFound
	}
	return comment, nil
}

// UpdateComment atualiza um comentário existente
func (s *CommentService) UpdateComment(ctx context.Context, id string, req CommentRequest, userID string) (*domain.Comment, error) {
	// Buscar o comentário existente
	comment, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrCommentNotFound
	}

	// Verificar se o usuário é o autor do comentário
	if !comment.IsAuthor(userID) {
		return nil, ErrNotAuthorized
	}

	// Atualizar o conteúdo do comentário
	if err := comment.UpdateContent(req.Content); err != nil {
		return nil, ErrInvalidCommentData
	}

	// Salvar as alterações no repositório
	err = s.commentRepo.Update(ctx, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// DeleteComment remove um comentário
func (s *CommentService) DeleteComment(ctx context.Context, id string, userID string) error {
	// Buscar o comentário existente
	comment, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		return ErrCommentNotFound
	}

	// Verificar se o usuário é o autor do comentário
	if !comment.IsAuthor(userID) {
		return ErrNotAuthorized
	}

	// Remover o comentário do repositório
	return s.commentRepo.Delete(ctx, id)
}

// ListCommentsByPost retorna uma lista paginada de comentários de um post específico
func (s *CommentService) ListCommentsByPost(ctx context.Context, postID string, page, pageSize int) ([]*domain.Comment, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Verificar se o post existe
	_, err := s.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, ErrPostNotFound
	}

	return s.commentRepo.ListByPost(ctx, postID, page, pageSize)
}

// ListCommentsByAuthor retorna uma lista paginada de comentários de um autor específico
func (s *CommentService) ListCommentsByAuthor(ctx context.Context, authorID string, page, pageSize int) ([]*domain.Comment, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Verificar se o autor existe
	_, err := s.userRepo.GetByID(ctx, authorID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return s.commentRepo.ListByAuthor(ctx, authorID, page, pageSize)
} 