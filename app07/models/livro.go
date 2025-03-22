package models

import (
	"time"
)

// Livro representa um livro na biblioteca
type Livro struct {
	ID            int       `json:"id"`
	Titulo        string    `json:"titulo"`
	Autor         string    `json:"autor"`
	Editora       string    `json:"editora"`
	AnoPublicacao int       `json:"anoPublicacao"`
	ISBN          string    `json:"isbn"`
	Disponivel    bool      `json:"disponivel"`
	DataCadastro  time.Time `json:"dataCadastro"`
}

// NovoLivro cria uma nova instância de livro
func NovoLivro(titulo, autor, editora, isbn string, anoPublicacao int) *Livro {
	return &Livro{
		Titulo:        titulo,
		Autor:         autor,
		Editora:       editora,
		AnoPublicacao: anoPublicacao,
		ISBN:          isbn,
		Disponivel:    true,
		DataCadastro:  time.Now(),
	}
} 