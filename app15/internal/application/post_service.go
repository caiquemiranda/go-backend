package application

import (
	"context"
	"errors"

	"app15/internal/domain"
	"app15/internal/ports/repositories"

	"github.com/google/uuid"
)

// Errors específicos do serviço de posts
var (
	ErrPostNotFound      = errors.New("post não encontrado")
	ErrNotAuthorized     = errors.New("não autorizado para esta ação")
	ErrInvalidPostData   = errors.New("dados do post inválidos")
)

// PostService representa o serviço de posts da aplicação
type PostService struct {
	postRepo repositories.PostRepository
	userRepo repositories.UserRepository
}

// NewPostService cria uma nova instância do serviço de posts
func NewPostService(postRepo repositories.PostRepository, userRepo repositories.UserRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

// PostRequest representa a estrutura de dados para criação/atualização de posts
type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreatePost cria um novo post
func (s *PostService) CreatePost(ctx context.Context, req PostRequest, authorID string) (*domain.Post, error) {
	// Verificar se o autor existe
	_, err := s.userRepo.GetByID(ctx, authorID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Criar novo post
	post, err := domain.NewPost(req.Title, req.Content, authorID)
	if err != nil {
		return nil, ErrInvalidPostData
	}

	// Gerar ID único para o post
	post.ID = uuid.New().String()

	// Salvar post no repositório
	err = s.postRepo.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetPostByID busca um post pelo ID
func (s *PostService) GetPostByID(ctx context.Context, id string) (*domain.Post, error) {
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrPostNotFound
	}
	return post, nil
}

// UpdatePost atualiza um post existente
func (s *PostService) UpdatePost(ctx context.Context, id string, req PostRequest, userID string) (*domain.Post, error) {
	// Buscar o post existente
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrPostNotFound
	}

	// Verificar se o usuário é o autor do post
	if !post.IsAuthor(userID) {
		return nil, ErrNotAuthorized
	}

	// Atualizar os campos do post
	if err := post.UpdateTitle(req.Title); err != nil {
		return nil, ErrInvalidPostData
	}

	if err := post.UpdateContent(req.Content); err != nil {
		return nil, ErrInvalidPostData
	}

	// Salvar as alterações no repositório
	err = s.postRepo.Update(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// DeletePost remove um post
func (s *PostService) DeletePost(ctx context.Context, id string, userID string) error {
	// Buscar o post existente
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return ErrPostNotFound
	}

	// Verificar se o usuário é o autor do post
	if !post.IsAuthor(userID) {
		return ErrNotAuthorized
	}

	// Remover o post do repositório
	return s.postRepo.Delete(ctx, id)
}

// ListPosts retorna uma lista paginada de posts
func (s *PostService) ListPosts(ctx context.Context, page, pageSize int) ([]*domain.Post, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.postRepo.List(ctx, page, pageSize)
}

// ListPostsByAuthor retorna uma lista paginada de posts de um autor específico
func (s *PostService) ListPostsByAuthor(ctx context.Context, authorID string, page, pageSize int) ([]*domain.Post, error) {
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

	return s.postRepo.ListByAuthor(ctx, authorID, page, pageSize)
} 