package handlers

import (
	"encoding/json"
	"net/http"

	"app15/internal/application"
)

// AuthHandler manipula as requisições de autenticação
type AuthHandler struct {
	authService *application.AuthService
}

// NewAuthHandler cria uma nova instância do AuthHandler
func NewAuthHandler(authService *application.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register registra um novo usuário
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req application.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Formato de requisição inválido", http.StatusBadRequest)
		return
	}

	resp, err := h.authService.Register(r.Context(), req)
	if err != nil {
		switch err {
		case application.ErrUserAlreadyExists:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// Login autentica um usuário
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req application.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Formato de requisição inválido", http.StatusBadRequest)
		return
	}

	resp, err := h.authService.Login(r.Context(), req)
	if err != nil {
		switch err {
		case application.ErrInvalidCredentials:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		default:
			http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Profile retorna o perfil do usuário autenticado
func (h *AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {
	// Obter ID do usuário do contexto (definido pelo middleware de autenticação)
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "Usuário não autenticado", http.StatusUnauthorized)
		return
	}

	user, err := h.authService.GetUserByID(r.Context(), userID)
	if err != nil {
		switch err {
		case application.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
} 