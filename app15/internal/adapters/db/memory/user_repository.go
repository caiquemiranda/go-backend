package memory

import (
	"context"
	"errors"
	"sync"

	"app15/internal/domain"
)

// Erros específicos do repositório
var (
	ErrUserNotFound      = errors.New("usuário não encontrado")
	ErrEmailAlreadyExists = errors.New("email já cadastrado")
	ErrUsernameAlreadyExists = errors.New("nome de usuário já cadastrado")
)

// UserRepository implementa o repositório de usuários em memória
type UserRepository struct {
	users map[string]*domain.User
	mu    sync.RWMutex
}

// NewUserRepository cria uma nova instância do repositório de usuários em memória
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*domain.User),
	}
}

// Create adiciona um novo usuário ao repositório
func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verificar se email já existe
	for _, u := range r.users {
		if u.Email == user.Email {
			return ErrEmailAlreadyExists
		}
		if u.Username == user.Username {
			return ErrUsernameAlreadyExists
		}
	}

	// Adicionar usuário
	r.users[user.ID] = user

	return nil
}

// GetByID busca um usuário pelo ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

// GetByEmail busca um usuário pelo email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, ErrUserNotFound
}

// GetByUsername busca um usuário pelo nome de usuário
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, ErrUserNotFound
}

// Update atualiza os dados de um usuário
func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verificar se usuário existe
	_, exists := r.users[user.ID]
	if !exists {
		return ErrUserNotFound
	}

	// Verificar se o novo email já está em uso por outro usuário
	for id, u := range r.users {
		if u.Email == user.Email && id != user.ID {
			return ErrEmailAlreadyExists
		}
		if u.Username == user.Username && id != user.ID {
			return ErrUsernameAlreadyExists
		}
	}

	// Atualizar usuário
	r.users[user.ID] = user

	return nil
}

// Delete remove um usuário do repositório
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verificar se usuário existe
	_, exists := r.users[id]
	if !exists {
		return ErrUserNotFound
	}

	// Remover usuário
	delete(r.users, id)

	return nil
}

// List retorna uma lista paginada de usuários
func (r *UserRepository) List(ctx context.Context, page, pageSize int) ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Calcular índices de paginação
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	// Copiar todos os usuários para um slice
	users := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	// Aplicar paginação
	if startIndex >= len(users) {
		return []*domain.User{}, nil
	}

	if endIndex > len(users) {
		endIndex = len(users)
	}

	return users[startIndex:endIndex], nil
} 