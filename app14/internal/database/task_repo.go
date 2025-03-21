package database

import (
	"errors"
	"sync"
	"time"

	"app14/internal/models"
)

var (
	ErrTaskNotFound = errors.New("tarefa não encontrada")
)

// TaskRepository define a interface para operações de repositório de tarefas
type TaskRepository interface {
	GetAll() ([]*models.Task, error)
	GetByID(id int) (*models.Task, error)
	Create(task *models.Task) error
	Update(id int, task *models.Task) error
	Delete(id int) error
}

// InMemoryTaskRepository implementa TaskRepository usando armazenamento em memória
type InMemoryTaskRepository struct {
	tasks  map[int]*models.Task
	nextID int
	mutex  sync.RWMutex
}

// NewInMemoryTaskRepository cria uma nova instância de InMemoryTaskRepository
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[int]*models.Task),
		nextID: 1,
	}
}

// GetAll retorna todas as tarefas
func (r *InMemoryTaskRepository) GetAll() ([]*models.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	tasks := make([]*models.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetByID retorna uma tarefa pelo ID
func (r *InMemoryTaskRepository) GetByID(id int) (*models.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

// Create cria uma nova tarefa
func (r *InMemoryTaskRepository) Create(task *models.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	task.ID = r.nextID
	r.nextID++

	r.tasks[task.ID] = task
	return nil
}

// Update atualiza uma tarefa existente
func (r *InMemoryTaskRepository) Update(id int, task *models.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	// Atualizar o timestamp
	task.UpdatedAt = time.Now()

	// Atualizar completedAt se o status for completado
	if task.Status == models.StatusCompleted && r.tasks[id].Status != models.StatusCompleted {
		now := time.Now()
		task.CompletedAt = &now
	} else if task.Status != models.StatusCompleted {
		task.CompletedAt = nil
	}

	r.tasks[id] = task
	return nil
}

// Delete remove uma tarefa
func (r *InMemoryTaskRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
} 