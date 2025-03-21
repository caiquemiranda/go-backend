package api

import (
	"net/http"
	"strings"

	"app14/internal/database"
	"app14/internal/handlers"
	"app14/internal/middleware"
)

// Router configura todas as rotas da API
type Router struct {
	taskHandler *handlers.TaskHandler
}

// NewRouter cria uma nova instância do Router
func NewRouter(taskRepo database.TaskRepository) *Router {
	return &Router{
		taskHandler: handlers.NewTaskHandler(taskRepo),
	}
}

// Setup configura o roteador HTTP
func (r *Router) Setup() http.Handler {
	// Criar o multiplexador
	mux := http.NewServeMux()

	// Aplicar middleware de logging a todas as requisições
	handler := middleware.Logger(mux)

	// Configurar rotas para a API de tarefas
	mux.HandleFunc("/api/tasks", r.handleTasksRoutes)
	mux.HandleFunc("/api/tasks/", r.handleTaskRoutes)

	return handler
}

// handleTasksRoutes gerencia as requisições para /api/tasks
func (r *Router) handleTasksRoutes(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.taskHandler.GetTasks(w, req)
	case http.MethodPost:
		r.taskHandler.CreateTask(w, req)
	default:
		w.Header().Set("Allow", "GET, POST")
		handlers.RespondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
}

// handleTaskRoutes gerencia as requisições para /api/tasks/{id}
func (r *Router) handleTaskRoutes(w http.ResponseWriter, req *http.Request) {
	// Verificar se a rota inclui um ID (ex: /api/tasks/{id})
	if !strings.HasPrefix(req.URL.Path, "/api/tasks/") {
		http.NotFound(w, req)
		return
	}

	switch req.Method {
	case http.MethodGet:
		r.taskHandler.GetTask(w, req)
	case http.MethodPut:
		r.taskHandler.UpdateTask(w, req)
	case http.MethodDelete:
		r.taskHandler.DeleteTask(w, req)
	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		handlers.RespondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
} 