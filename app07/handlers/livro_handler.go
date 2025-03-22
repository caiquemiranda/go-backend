package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"../models"
	"../services"
)

// LivroHandler gerencia as requisições HTTP relacionadas a livros
type LivroHandler struct {
	service *services.LivroService
}

// NovoLivroHandler cria uma nova instância do handler de livros
func NovoLivroHandler(service *services.LivroService) *LivroHandler {
	return &LivroHandler{
		service: service,
	}
}

// LivrosHandler gerencia requisições para a rota /livros
func (h *LivroHandler) LivrosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listarLivros(w, r)
	case http.MethodPost:
		h.criarLivro(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// LivroHandler gerencia requisições para a rota /livros/{id}
func (h *LivroHandler) LivroHandler(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID do caminho
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Caminho inválido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.obterLivro(w, r, id)
	case http.MethodPut:
		h.atualizarLivro(w, r, id)
	case http.MethodDelete:
		h.removerLivro(w, r, id)
	case http.MethodPatch:
		h.alterarDisponibilidade(w, r, id)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// listarLivros retorna todos os livros
func (h *LivroHandler) listarLivros(w http.ResponseWriter, r *http.Request) {
	livros := h.service.ObterTodos()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livros)
}

// criarLivro cria um novo livro
func (h *LivroHandler) criarLivro(w http.ResponseWriter, r *http.Request) {
	var livro models.Livro
	if err := json.NewDecoder(r.Body).Decode(&livro); err != nil {
		http.Error(w, "Erro ao decodificar dados do livro", http.StatusBadRequest)
		return
	}

	novoLivro, err := h.service.Criar(&livro)
	if err != nil {
		if err == services.ErrISBNJaExiste {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoLivro)
}

// obterLivro retorna um livro pelo ID
func (h *LivroHandler) obterLivro(w http.ResponseWriter, r *http.Request, id int) {
	livro, err := h.service.ObterPorID(id)
	if err != nil {
		if err == services.ErrLivroNaoEncontrado {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livro)
}

// atualizarLivro atualiza um livro existente
func (h *LivroHandler) atualizarLivro(w http.ResponseWriter, r *http.Request, id int) {
	var livro models.Livro
	if err := json.NewDecoder(r.Body).Decode(&livro); err != nil {
		http.Error(w, "Erro ao decodificar dados do livro", http.StatusBadRequest)
		return
	}

	livroAtualizado, err := h.service.Atualizar(id, &livro)
	if err != nil {
		switch err {
		case services.ErrLivroNaoEncontrado:
			http.Error(w, err.Error(), http.StatusNotFound)
		case services.ErrISBNJaExiste:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livroAtualizado)
}

// removerLivro exclui um livro
func (h *LivroHandler) removerLivro(w http.ResponseWriter, r *http.Request, id int) {
	err := h.service.Remover(id)
	if err != nil {
		if err == services.ErrLivroNaoEncontrado {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// alterarDisponibilidade muda o status de disponibilidade de um livro
func (h *LivroHandler) alterarDisponibilidade(w http.ResponseWriter, r *http.Request, id int) {
	// Verifica se o parâmetro foi enviado
	disponivel := r.URL.Query().Get("disponivel")
	if disponivel == "" {
		http.Error(w, "Parâmetro 'disponivel' é obrigatório", http.StatusBadRequest)
		return
	}

	// Converte para booleano
	disp := disponivel == "true"

	livro, err := h.service.AlterarDisponibilidade(id, disp)
	if err != nil {
		if err == services.ErrLivroNaoEncontrado {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livro)
} 