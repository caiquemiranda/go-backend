package models

import (
	"errors"
	"time"
)

// Status da tarefa como tipo enum
type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
	StatusCancelled  TaskStatus = "cancelled"
)

// Task representa uma tarefa no sistema
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// TaskInput representa os dados de entrada para criação/atualização de uma tarefa
type TaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status,omitempty"`
}

// Validate valida os dados da tarefa
func (t *TaskInput) Validate() error {
	if t.Title == "" {
		return errors.New("o título da tarefa é obrigatório")
	}

	if t.Status != "" && 
	   t.Status != StatusPending && 
	   t.Status != StatusInProgress && 
	   t.Status != StatusCompleted && 
	   t.Status != StatusCancelled {
		return errors.New("status inválido")
	}

	return nil
}

// NewTask cria uma nova instância de Task a partir de TaskInput
func NewTask(input TaskInput) *Task {
	now := time.Now()
	status := input.Status
	if status == "" {
		status = StatusPending
	}

	return &Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
} 