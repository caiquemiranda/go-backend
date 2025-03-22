package services

import (
	"errors"
	"sync"

	"../models"
)

// Erros comuns
var (
	ErrLivroNaoEncontrado = errors.New("livro não encontrado")
	ErrISBNJaExiste       = errors.New("já existe um livro com este ISBN")
)

// LivroService gerencia operações relacionadas a livros
type LivroService struct {
	livros    map[int]*models.Livro
	proximoID int
	mutex     sync.RWMutex
}

// NovoLivroService cria uma nova instância do serviço de livros
func NovoLivroService() *LivroService {
	return &LivroService{
		livros:    make(map[int]*models.Livro),
		proximoID: 1,
		mutex:     sync.RWMutex{},
	}
}

// Criar adiciona um novo livro
func (s *LivroService) Criar(livro *models.Livro) (*models.Livro, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Verifica se o ISBN já existe
	for _, l := range s.livros {
		if l.ISBN == livro.ISBN {
			return nil, ErrISBNJaExiste
		}
	}

	// Atribui ID e adiciona ao mapa
	livro.ID = s.proximoID
	s.proximoID++
	s.livros[livro.ID] = livro

	return livro, nil
}

// ObterTodos retorna todos os livros
func (s *LivroService) ObterTodos() []*models.Livro {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	livros := make([]*models.Livro, 0, len(s.livros))
	for _, livro := range s.livros {
		livros = append(livros, livro)
	}

	return livros
}

// ObterPorID busca um livro pelo ID
func (s *LivroService) ObterPorID(id int) (*models.Livro, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	livro, existe := s.livros[id]
	if !existe {
		return nil, ErrLivroNaoEncontrado
	}

	return livro, nil
}

// Atualizar modifica os dados de um livro existente
func (s *LivroService) Atualizar(id int, livro *models.Livro) (*models.Livro, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Verifica se o livro existe
	_, existe := s.livros[id]
	if !existe {
		return nil, ErrLivroNaoEncontrado
	}

	// Verifica se o ISBN já existe (se foi alterado)
	livroAtual := s.livros[id]
	if livro.ISBN != livroAtual.ISBN {
		for _, l := range s.livros {
			if l.ID != id && l.ISBN == livro.ISBN {
				return nil, ErrISBNJaExiste
			}
		}
	}

	// Preserva alguns campos
	livro.ID = id
	livro.DataCadastro = livroAtual.DataCadastro

	// Atualiza o livro
	s.livros[id] = livro

	return livro, nil
}

// Remover exclui um livro pelo ID
func (s *LivroService) Remover(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, existe := s.livros[id]; !existe {
		return ErrLivroNaoEncontrado
	}

	delete(s.livros, id)
	return nil
}

// AlterarDisponibilidade muda o status de disponibilidade de um livro
func (s *LivroService) AlterarDisponibilidade(id int, disponivel bool) (*models.Livro, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	livro, existe := s.livros[id]
	if !existe {
		return nil, ErrLivroNaoEncontrado
	}

	livro.Disponivel = disponivel
	return livro, nil
} 