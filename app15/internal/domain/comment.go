package domain

import (
	"errors"
	"time"
)

// Comment representa a entidade de comentário no domínio
type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	PostID    string    `json:"post_id"`
	AuthorID  string    `json:"author_id"`
	Author    *User     `json:"author,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewComment cria uma nova instância de Comment
func NewComment(content, postID, authorID string) (*Comment, error) {
	if content == "" {
		return nil, errors.New("conteúdo não pode ser vazio")
	}
	
	if postID == "" {
		return nil, errors.New("ID do post não pode ser vazio")
	}
	
	if authorID == "" {
		return nil, errors.New("ID do autor não pode ser vazio")
	}
	
	now := time.Now()
	
	return &Comment{
		Content:   content,
		PostID:    postID,
		AuthorID:  authorID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// UpdateContent atualiza o conteúdo do comentário
func (c *Comment) UpdateContent(content string) error {
	if content == "" {
		return errors.New("conteúdo não pode ser vazio")
	}
	
	c.Content = content
	c.UpdatedAt = time.Now()
	return nil
}

// IsAuthor verifica se o usuário especificado é o autor do comentário
func (c *Comment) IsAuthor(userID string) bool {
	return c.AuthorID == userID
} 