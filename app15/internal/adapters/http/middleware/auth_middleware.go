package middleware

import (
	"context"
	"net/http"
	"strings"

	"app15/internal/application"
)

// AuthMiddleware é um middleware para autenticação de usuários
type AuthMiddleware struct {
	authService *application.AuthService
}

// NewAuthMiddleware cria uma nova instância do AuthMiddleware
func NewAuthMiddleware(authService *application.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// Authenticate valida o token JWT e injeta o ID do usuário no contexto
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrair token do cabeçalho Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token de autenticação não fornecido", http.StatusUnauthorized)
			return
		}

		// Verificar formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato de token inválido", http.StatusUnauthorized)
			return
		}

		// Validar token
		userID, err := m.authService.ValidateToken(parts[1])
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Adicionar ID do usuário ao contexto
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Optional é semelhante ao Authenticate, mas não retorna erro se o token não for fornecido
func (m *AuthMiddleware) Optional(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrair token do cabeçalho Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// Continuar sem autenticação
			next.ServeHTTP(w, r)
			return
		}

		// Verificar formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// Continuar sem autenticação
			next.ServeHTTP(w, r)
			return
		}

		// Validar token
		userID, err := m.authService.ValidateToken(parts[1])
		if err != nil {
			// Continuar sem autenticação
			next.ServeHTTP(w, r)
			return
		}

		// Adicionar ID do usuário ao contexto
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
} 