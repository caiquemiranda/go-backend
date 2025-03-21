package domain

import (
	"errors"
	"time"
)

// Post representa a entidade de post de blog no domínio
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"author_id"`
	Author    *User     `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPost cria uma nova instância de Post
func NewPost(title, content, authorID string) (*Post, error) {
	if title == "" {
		return nil, errors.New("título não pode ser vazio")
	}
	
	if content == "" {
		return nil, errors.New("conteúdo não pode ser vazio")
	}
	
	if authorID == "" {
		return nil, errors.New("ID do autor não pode ser vazio")
	}
	
	now := time.Now()
	
	return &Post{
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// UpdateTitle atualiza o título do post
func (p *Post) UpdateTitle(title string) error {
	if title == "" {
		return errors.New("título não pode ser vazio")
	}
	
	p.Title = title
	p.UpdatedAt = time.Now()
	return nil
}

// UpdateContent atualiza o conteúdo do post
func (p *Post) UpdateContent(content string) error {
	if content == "" {
		return errors.New("conteúdo não pode ser vazio")
	}
	
	p.Content = content
	p.UpdatedAt = time.Now()
	return nil
}

// IsAuthor verifica se o usuário especificado é o autor do post
func (p *Post) IsAuthor(userID string) bool {
	return p.AuthorID == userID
} 