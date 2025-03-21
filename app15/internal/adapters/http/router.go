package http

import (
	"net/http"
	"time"

	"app15/internal/adapters/db/memory"
	"app15/internal/adapters/http/handlers"
	"app15/internal/adapters/http/middleware"
	"app15/internal/application"
	"app15/internal/ports/repositories"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// SetupRouter configura as rotas e middlewares da aplicação
func SetupRouter() http.Handler {
	// Inicializar repositórios
	userRepo := memory.NewUserRepository()

	// Inicializar serviços de aplicação
	jwtKey := "sua-chave-secreta-aqui" // Deve ser configurado via ambiente em produção
	jwtExp := 24 * time.Hour
	authService := application.NewAuthService(userRepo, jwtKey, jwtExp)

	// Inicializar handlers HTTP
	authHandler := handlers.NewAuthHandler(authService)

	// Inicializar middlewares
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Configurar router
	r := chi.NewRouter()

	// Middlewares básicos
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.Timeout(60 * time.Second))
	
	// Configuração de CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Em produção, especificar origens permitidas
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not entirely supported by all browsers
	}))

	// Rota para verificar saúde
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Rotas de API
	r.Route("/api", func(r chi.Router) {
		// Rotas de autenticação (públicas)
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
			r.Post("/login", authHandler.Login)
			
			// Rotas protegidas por autenticação
			r.Group(func(r chi.Router) {
				r.Use(authMiddleware.Authenticate)
				r.Get("/profile", authHandler.Profile)
			})
		})

		// TODO: Adicionar outras rotas (posts, comentários, etc.)
	})

	return r
}

// NewRepositories cria e retorna instâncias de todos os repositórios
func NewRepositories() (repositories.UserRepository, repositories.PostRepository, repositories.CommentRepository) {
	userRepo := memory.NewUserRepository()
	// TODO: Implementar e inicializar os repositórios de posts e comentários
	var postRepo repositories.PostRepository
	var commentRepo repositories.CommentRepository
	
	return userRepo, postRepo, commentRepo
} 