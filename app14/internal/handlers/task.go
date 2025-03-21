package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"app14/internal/database"
	"app14/internal/models"
)

// TaskHandler contém os handlers para a API de tarefas
type TaskHandler struct {
	repo database.TaskRepository
}

// NewTaskHandler cria uma nova instância de TaskHandler
func NewTaskHandler(repo database.TaskRepository) *TaskHandler {
	return &TaskHandler{
		repo: repo,
	}
}

// GetTasks retorna todas as tarefas
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.GetAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, tasks)
}

// GetTask retorna uma tarefa específica pelo ID
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskIDFromURL(r.URL.Path)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	task, err := h.repo.GetByID(id)
	if err != nil {
		if err == database.ErrTaskNotFound {
			RespondWithError(w, http.StatusNotFound, "Tarefa não encontrada")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, task)
}

// CreateTask cria uma nova tarefa
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input models.TaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Dados inválidos")
		return
	}

	if err := input.Validate(); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	task := models.NewTask(input)
	if err := h.repo.Create(task); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, task)
}

// UpdateTask atualiza uma tarefa existente
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskIDFromURL(r.URL.Path)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	// Verificar se a tarefa existe
	existingTask, err := h.repo.GetByID(id)
	if err != nil {
		if err == database.ErrTaskNotFound {
			RespondWithError(w, http.StatusNotFound, "Tarefa não encontrada")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Decodificar os dados da requisição
	var input models.TaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Dados inválidos")
		return
	}

	if err := input.Validate(); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Atualizar os campos da tarefa
	existingTask.Title = input.Title
	existingTask.Description = input.Description
	if input.Status != "" {
		existingTask.Status = input.Status
	}

	// Salvar as alterações
	if err := h.repo.Update(id, existingTask); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, existingTask)
}

// DeleteTask deleta uma tarefa
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskIDFromURL(r.URL.Path)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.repo.Delete(id); err != nil {
		if err == database.ErrTaskNotFound {
			RespondWithError(w, http.StatusNotFound, "Tarefa não encontrada")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getTaskIDFromURL extrai o ID da tarefa da URL
func getTaskIDFromURL(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, ErrInvalidID
	}

	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil || id < 1 {
		return 0, ErrInvalidID
	}

	return id, nil
}

// Erros comuns
var (
	ErrInvalidID = NewAPIError("ID inválido", http.StatusBadRequest)
)

// APIError representa um erro da API
type APIError struct {
	Message string
	Code    int
}

func (e APIError) Error() string {
	return e.Message
}

// NewAPIError cria um novo APIError
func NewAPIError(message string, code int) APIError {
	return APIError{
		Message: message,
		Code:    code,
	}
}

// RespondWithError envia uma resposta JSON com um erro
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON envia uma resposta JSON
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Erro ao processar resposta JSON"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
} 