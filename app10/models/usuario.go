package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Usuario representa um usuário no sistema
type Usuario struct {
	ID             uint      `json:"id"`
	Nome           string    `json:"nome"`
	Email          string    `json:"email"`
	Senha          string    `json:"-"` // A senha não é retornada no JSON
	Perfil         string    `json:"perfil"` // admin, editor, usuario
	DataCriacao    time.Time `json:"dataCriacao"`
	UltimoAcesso   time.Time `json:"ultimoAcesso"`
	Ativo          bool      `json:"ativo"`
	TokenAtivacao  string    `json:"-"`
	TokenResetSenha string   `json:"-"`
}

// CredenciaisLogin representa os dados necessários para login
type CredenciaisLogin struct {
	Email    string `json:"email"`
	Senha    string `json:"senha"`
}

// NovoUsuario cria uma instância de usuário com valores padrão
func NovoUsuario(nome, email, senha, perfil string) (*Usuario, error) {
	if nome == "" || email == "" || senha == "" {
		return nil, errors.New("campos obrigatórios faltando")
	}

	if len(senha) < 6 {
		return nil, errors.New("a senha deve ter pelo menos 6 caracteres")
	}

	// Perfil padrão é "usuario" se não for especificado
	if perfil == "" {
		perfil = "usuario"
	}

	// Hash da senha
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Usuario{
		Nome:           nome,
		Email:          email,
		Senha:          string(senhaHash),
		Perfil:         perfil,
		DataCriacao:    time.Now(),
		UltimoAcesso:   time.Now(),
		Ativo:          false, // Usuário não está ativo até confirmar email
		TokenAtivacao:  gerarTokenAleatorio(),
		TokenResetSenha: "",
	}, nil
}

// VerificarSenha compara a senha fornecida com a senha armazenada
func (u *Usuario) VerificarSenha(senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha))
	return err == nil
}

// AlterarSenha muda a senha do usuário
func (u *Usuario) AlterarSenha(novaSenha string) error {
	if len(novaSenha) < 6 {
		return errors.New("a senha deve ter pelo menos 6 caracteres")
	}

	senhaHash, err := bcrypt.GenerateFromPassword([]byte(novaSenha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Senha = string(senhaHash)
	return nil
}

// AtualizarUltimoAcesso atualiza o timestamp de último acesso
func (u *Usuario) AtualizarUltimoAcesso() {
	u.UltimoAcesso = time.Now()
}

// Ativar ativa a conta do usuário
func (u *Usuario) Ativar() {
	u.Ativo = true
	u.TokenAtivacao = ""
}

// GerarTokenResetSenha cria um token para resetar a senha
func (u *Usuario) GerarTokenResetSenha() string {
	u.TokenResetSenha = gerarTokenAleatorio()
	return u.TokenResetSenha
}

// ResetarSenha redefine a senha do usuário
func (u *Usuario) ResetarSenha(token, novaSenha string) error {
	if u.TokenResetSenha == "" || u.TokenResetSenha != token {
		return errors.New("token inválido")
	}

	if err := u.AlterarSenha(novaSenha); err != nil {
		return err
	}

	u.TokenResetSenha = ""
	return nil
}

// gerarTokenAleatorio cria um token aleatório para ativação de conta ou reset de senha
func gerarTokenAleatorio() string {
	// Na implementação real, usaria crypto/rand para gerar um token seguro
	// Para simplicidade, retornamos um timestamp em formato string
	return time.Now().Format("20060102150405.000000000")
}

// DadosUsuarioPublicos retorna apenas os dados públicos do usuário
type DadosUsuarioPublicos struct {
	ID           uint      `json:"id"`
	Nome         string    `json:"nome"`
	Email        string    `json:"email"`
	Perfil       string    `json:"perfil"`
	DataCriacao  time.Time `json:"dataCriacao"`
	UltimoAcesso time.Time `json:"ultimoAcesso"`
	Ativo        bool      `json:"ativo"`
}

// ParaPublico converte um usuário para seus dados públicos
func (u *Usuario) ParaPublico() DadosUsuarioPublicos {
	return DadosUsuarioPublicos{
		ID:           u.ID,
		Nome:         u.Nome,
		Email:        u.Email,
		Perfil:       u.Perfil,
		DataCriacao:  u.DataCriacao,
		UltimoAcesso: u.UltimoAcesso,
		Ativo:        u.Ativo,
	}
} 