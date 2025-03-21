package models

import (
	"errors"
	"sync"
)

// Erros comuns
var (
	ErrUsuarioNaoEncontrado = errors.New("usuário não encontrado")
	ErrRecursoNaoEncontrado = errors.New("recurso não encontrado")
	ErrEmailJaExiste        = errors.New("email já existe")
)

// RepositorioUsuarioMemoria implementa o RepositorioUsuario com armazenamento em memória
type RepositorioUsuarioMemoria struct {
	usuarios map[uint]*Usuario
	emailIdx map[string]uint // Índice para busca por email
	nextID   uint
	mu       sync.RWMutex
}

// NovoRepositorioUsuarioMemoria cria um novo repositório de usuários em memória
func NovoRepositorioUsuarioMemoria() *RepositorioUsuarioMemoria {
	return &RepositorioUsuarioMemoria{
		usuarios: make(map[uint]*Usuario),
		emailIdx: make(map[string]uint),
		nextID:   1,
	}
}

// Criar adiciona um novo usuário
func (r *RepositorioUsuarioMemoria) Criar(usuario *Usuario) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verifica se o email já existe
	if _, ok := r.emailIdx[usuario.Email]; ok {
		return ErrEmailJaExiste
	}

	// Atribui um ID ao usuário
	usuario.ID = r.nextID
	r.nextID++

	// Armazena o usuário
	r.usuarios[usuario.ID] = usuario
	r.emailIdx[usuario.Email] = usuario.ID

	return nil
}

// ObterPorID busca um usuário pelo ID
func (r *RepositorioUsuarioMemoria) ObterPorID(id uint) (*Usuario, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	usuario, ok := r.usuarios[id]
	if !ok {
		return nil, ErrUsuarioNaoEncontrado
	}

	return usuario, nil
}

// ObterPorEmail busca um usuário pelo email
func (r *RepositorioUsuarioMemoria) ObterPorEmail(email string) (*Usuario, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id, ok := r.emailIdx[email]
	if !ok {
		return nil, ErrUsuarioNaoEncontrado
	}

	return r.usuarios[id], nil
}

// Atualizar modifica um usuário existente
func (r *RepositorioUsuarioMemoria) Atualizar(id uint, usuario *Usuario) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verifica se o usuário existe
	_, ok := r.usuarios[id]
	if !ok {
		return ErrUsuarioNaoEncontrado
	}

	// Se o email foi alterado, atualiza o índice
	existente := r.usuarios[id]
	if existente.Email != usuario.Email {
		delete(r.emailIdx, existente.Email)
		r.emailIdx[usuario.Email] = id
	}

	// Atualiza o usuário
	r.usuarios[id] = usuario

	return nil
}

// Remover exclui um usuário
func (r *RepositorioUsuarioMemoria) Remover(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verifica se o usuário existe
	usuario, ok := r.usuarios[id]
	if !ok {
		return ErrUsuarioNaoEncontrado
	}

	// Remove o usuário
	delete(r.usuarios, id)
	delete(r.emailIdx, usuario.Email)

	return nil
}

// RepositorioRecursoMemoria implementa o RepositorioRecurso com armazenamento em memória
type RepositorioRecursoMemoria struct {
	recursos map[uint]*Recurso
	nextID   uint
	mu       sync.RWMutex
}

// NovoRepositorioRecursoMemoria cria um novo repositório de recursos em memória
func NovoRepositorioRecursoMemoria() *RepositorioRecursoMemoria {
	return &RepositorioRecursoMemoria{
		recursos: make(map[uint]*Recurso),
		nextID:   1,
	}
}

// Criar adiciona um novo recurso
func (r *RepositorioRecursoMemoria) Criar(recurso *Recurso) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Atribui um ID ao recurso
	recurso.ID = r.nextID
	r.nextID++

	// Armazena o recurso
	r.recursos[recurso.ID] = recurso

	return nil
}

// ObterPorID busca um recurso pelo ID
func (r *RepositorioRecursoMemoria) ObterPorID(id uint) (*Recurso, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	recurso, ok := r.recursos[id]
	if !ok {
		return nil, ErrRecursoNaoEncontrado
	}

	return recurso, nil
}

// ObterTodos retorna todos os recursos com base nos filtros
func (r *RepositorioRecursoMemoria) ObterTodos(filtros map[string]interface{}) ([]*Recurso, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Prepara o resultado
	resultado := make([]*Recurso, 0)

	// Obtém o nível de acesso do usuário para filtrar recursos
	var nivelAcesso int
	if nivel, ok := filtros["nivelAcesso"]; ok {
		nivelAcesso, _ = nivel.(int)
	}

	// Obtém a categoria para filtrar recursos
	var categoria string
	if cat, ok := filtros["categoria"]; ok {
		categoria, _ = cat.(string)
	}

	// Filtra os recursos
	for _, recurso := range r.recursos {
		// Verifica se o recurso está publicado
		if !recurso.Publicado {
			continue
		}

		// Verifica o nível de acesso
		if recurso.AcessoLevel > nivelAcesso {
			continue
		}

		// Filtra por categoria se especificada
		if categoria != "" && recurso.Categoria != categoria {
			continue
		}

		resultado = append(resultado, recurso)
	}

	return resultado, nil
}

// Atualizar modifica um recurso existente
func (r *RepositorioRecursoMemoria) Atualizar(id uint, recurso *Recurso) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verifica se o recurso existe
	_, ok := r.recursos[id]
	if !ok {
		return ErrRecursoNaoEncontrado
	}

	// Atualiza o recurso
	r.recursos[id] = recurso

	return nil
}

// Remover exclui um recurso
func (r *RepositorioRecursoMemoria) Remover(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verifica se o recurso existe
	_, ok := r.recursos[id]
	if !ok {
		return ErrRecursoNaoEncontrado
	}

	// Remove o recurso
	delete(r.recursos, id)

	return nil
} 