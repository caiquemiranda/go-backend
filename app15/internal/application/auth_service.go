package application

import (
	"context"
	"errors"
	"time"

	"app15/internal/domain"
	"app15/internal/ports/repositories"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Errors específicos do serviço de autenticação
var (
	ErrInvalidCredentials = errors.New("credenciais inválidas")
	ErrUserAlreadyExists  = errors.New("usuário já existe")
	ErrUserNotFound       = errors.New("usuário não encontrado")
)

// AuthService representa o serviço de autenticação da aplicação
type AuthService struct {
	userRepo repositories.UserRepository
	jwtKey   []byte
	jwtExp   time.Duration
}

// NewAuthService cria uma nova instância do serviço de autenticação
func NewAuthService(userRepo repositories.UserRepository, jwtKey string, jwtExp time.Duration) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtKey:   []byte(jwtKey),
		jwtExp:   jwtExp,
	}
}

// LoginRequest representa a estrutura de dados para requisição de login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest representa a estrutura de dados para requisição de registro
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse representa a estrutura de dados de resposta após autenticação
type AuthResponse struct {
	Token string      `json:"token"`
	User  *domain.User `json:"user"`
}

// Login autentica um usuário e retorna um token JWT
func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !user.ValidatePassword(req.Password) {
		return nil, ErrInvalidCredentials
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

// Register registra um novo usuário no sistema
func (s *AuthService) Register(ctx context.Context, req RegisterRequest) (*AuthResponse, error) {
	// Verificar se o email já está em uso
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// Verificar se o username já está em uso
	existingUser, err = s.userRepo.GetByUsername(ctx, req.Username)
	if err == nil && existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// Criar novo usuário
	user, err := domain.NewUser(req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	// Gerar ID único para o usuário
	user.ID = uuid.New().String()

	// Salvar usuário no repositório
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// Gerar token JWT
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

// GetUserByID busca um usuário pelo ID
func (s *AuthService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// ValidateToken valida um token JWT e retorna o ID do usuário
func (s *AuthService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inesperado")
		}
		return s.jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.New("token inválido: campo user_id ausente")
		}
		return userID, nil
	}

	return "", errors.New("token inválido")
}

// generateToken gera um novo token JWT para o usuário
func (s *AuthService) generateToken(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(s.jwtExp)

	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"username":  user.Username,
		"exp":       expirationTime.Unix(),
		"issued_at": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
} 