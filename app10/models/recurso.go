package models

import (
	"time"
)

// Recurso representa um recurso protegido no sistema
type Recurso struct {
	ID          uint      `json:"id"`
	Titulo      string    `json:"titulo"`
	Descricao   string    `json:"descricao"`
	Conteudo    string    `json:"conteudo"`
	Categoria   string    `json:"categoria"`
	AcessoLevel int       `json:"acessoLevel"` // 0=público, 1=usuário, 2=editor, 3=admin
	CriadoPor   uint      `json:"criadoPor"`
	DataCriacao time.Time `json:"dataCriacao"`
	Publicado   bool      `json:"publicado"`
}

// NovoRecurso cria uma nova instância de recurso
func NovoRecurso(titulo, descricao, conteudo, categoria string, acessoLevel int, criadoPor uint) *Recurso {
	return &Recurso{
		Titulo:      titulo,
		Descricao:   descricao,
		Conteudo:    conteudo,
		Categoria:   categoria,
		AcessoLevel: acessoLevel,
		CriadoPor:   criadoPor,
		DataCriacao: time.Now(),
		Publicado:   false,
	}
}

// Publicar marca o recurso como publicado
func (r *Recurso) Publicar() {
	r.Publicado = true
}

// Despublicar marca o recurso como não publicado
func (r *Recurso) Despublicar() {
	r.Publicado = false
}

// AtualizarConteudo atualiza o conteúdo do recurso
func (r *Recurso) AtualizarConteudo(titulo, descricao, conteudo, categoria string) {
	if titulo != "" {
		r.Titulo = titulo
	}
	
	if descricao != "" {
		r.Descricao = descricao
	}
	
	if conteudo != "" {
		r.Conteudo = conteudo
	}
	
	if categoria != "" {
		r.Categoria = categoria
	}
}

// AlterarNivelAcesso modifica o nível de acesso necessário para o recurso
func (r *Recurso) AlterarNivelAcesso(nivel int) {
	// Garante que o nível está entre 0 e 3
	if nivel < 0 {
		nivel = 0
	} else if nivel > 3 {
		nivel = 3
	}
	
	r.AcessoLevel = nivel
}

// RecursoVisaoPublica representa a visão pública de um recurso
type RecursoVisaoPublica struct {
	ID          uint      `json:"id"`
	Titulo      string    `json:"titulo"`
	Descricao   string    `json:"descricao"`
	Categoria   string    `json:"categoria"`
	DataCriacao time.Time `json:"dataCriacao"`
}

// ParaPublico converte um recurso para sua visão pública (sem conteúdo completo)
func (r *Recurso) ParaPublico() RecursoVisaoPublica {
	return RecursoVisaoPublica{
		ID:          r.ID,
		Titulo:      r.Titulo,
		Descricao:   r.Descricao,
		Categoria:   r.Categoria,
		DataCriacao: r.DataCriacao,
	}
} 