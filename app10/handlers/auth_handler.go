package handlers

import (
	"encoding/json"
	"net/http"

	"../models"
	"../utils"
)

// RepositorioUsuario define a interface para acessar dados de usuário
type RepositorioUsuario interface {
	ObterPorEmail(email string) (*models.Usuario, error)
	Criar(usuario *models.Usuario) error
	Atualizar(id uint, usuario *models.Usuario) error
}

// AuthHandler gerencia as rotas de autenticação
type AuthHandler struct {
	repo RepositorioUsuario
}

// NovoAuthHandler cria uma nova instância do handler de autenticação
func NovoAuthHandler(repo RepositorioUsuario) *AuthHandler {
	return &AuthHandler{
		repo: repo,
	}
}

// RespostaToken representa a resposta para login bem-sucedido
type RespostaToken struct {
	Token        string                     `json:"token"`
	RefreshToken string                     `json:"refreshToken"`
	Usuario      models.DadosUsuarioPublicos `json:"usuario"`
}

// RespostaErro representa uma resposta de erro
type RespostaErro struct {
	Status  int    `json:"status"`
	Mensagem string `json:"mensagem"`
}

// respondJSON envia uma resposta JSON para o cliente
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// respondErro envia uma resposta de erro para o cliente
func respondErro(w http.ResponseWriter, status int, mensagem string) {
	respondJSON(w, status, RespostaErro{
		Status:  status,
		Mensagem: mensagem,
	})
}

// Login gerencia a rota de login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	var credenciais models.CredenciaisLogin
	err := json.NewDecoder(r.Body).Decode(&credenciais)
	if err != nil {
		respondErro(w, http.StatusBadRequest, "Erro ao decodificar JSON: "+err.Error())
		return
	}

	// Valida os campos obrigatórios
	if credenciais.Email == "" || credenciais.Senha == "" {
		respondErro(w, http.StatusBadRequest, "Email e senha são obrigatórios")
		return
	}

	// Busca o usuário pelo email
	usuario, err := h.repo.ObterPorEmail(credenciais.Email)
	if err != nil {
		respondErro(w, http.StatusUnauthorized, "Credenciais inválidas")
		return
	}

	// Verifica se o usuário está ativo
	if !usuario.Ativo {
		respondErro(w, http.StatusUnauthorized, "Conta não ativada. Por favor, verifique seu email")
		return
	}

	// Verifica a senha
	if !usuario.VerificarSenha(credenciais.Senha) {
		respondErro(w, http.StatusUnauthorized, "Credenciais inválidas")
		return
	}

	// Atualiza o último acesso
	usuario.AtualizarUltimoAcesso()
	h.repo.Atualizar(usuario.ID, usuario)

	// Gera o token JWT
	token, err := utils.GerarToken(usuario.ID, usuario.Email, usuario.Perfil)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao gerar token: "+err.Error())
		return
	}

	// Gera o refresh token
	refreshToken, err := utils.GerarRefreshToken(usuario.ID, usuario.Email)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao gerar refresh token: "+err.Error())
		return
	}

	// Retorna o token e os dados do usuário
	respondJSON(w, http.StatusOK, RespostaToken{
		Token:        token,
		RefreshToken: refreshToken,
		Usuario:      usuario.ParaPublico(),
	})
}

// Registro gerencia a rota de registro de novos usuários
func (h *AuthHandler) Registro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	type DadosRegistro struct {
		Nome   string `json:"nome"`
		Email  string `json:"email"`
		Senha  string `json:"senha"`
		Perfil string `json:"perfil"`
	}

	var dados DadosRegistro
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		respondErro(w, http.StatusBadRequest, "Erro ao decodificar JSON: "+err.Error())
		return
	}

	// Verifica se o email já está em uso
	existente, _ := h.repo.ObterPorEmail(dados.Email)
	if existente != nil {
		respondErro(w, http.StatusConflict, "Email já está em uso")
		return
	}

	// Por segurança, o perfil é sempre 'usuario' para novos registros via API pública
	// Apenas administradores podem criar usuários com outros perfis
	dados.Perfil = "usuario"

	// Cria o novo usuário
	usuario, err := models.NovoUsuario(dados.Nome, dados.Email, dados.Senha, dados.Perfil)
	if err != nil {
		respondErro(w, http.StatusBadRequest, "Erro ao criar usuário: "+err.Error())
		return
	}

	// Na aplicação real, enviaria um email de confirmação aqui

	// Salva o usuário no repositório
	err = h.repo.Criar(usuario)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao salvar usuário: "+err.Error())
		return
	}

	// Ativa o usuário automaticamente (em uma aplicação real, exigiria confirmação por email)
	usuario.Ativo = true

	// Retorna os dados do usuário (sem senha)
	respondJSON(w, http.StatusCreated, usuario.ParaPublico())
}

// RefreshToken gerencia a renovação de tokens
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	type RefreshRequest struct {
		RefreshToken string `json:"refreshToken"`
	}

	var req RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondErro(w, http.StatusBadRequest, "Erro ao decodificar JSON: "+err.Error())
		return
	}

	// Valida o refresh token
	claims, err := utils.ValidarRefreshToken(req.RefreshToken)
	if err != nil {
		respondErro(w, http.StatusUnauthorized, "Refresh token inválido: "+err.Error())
		return
	}

	// Busca o usuário pelo email
	usuario, err := h.repo.ObterPorEmail(claims.Email)
	if err != nil {
		respondErro(w, http.StatusUnauthorized, "Usuário não encontrado")
		return
	}

	// Verifica se o usuário está ativo
	if !usuario.Ativo {
		respondErro(w, http.StatusUnauthorized, "Conta não ativada")
		return
	}

	// Gera um novo token JWT
	token, err := utils.GerarToken(usuario.ID, usuario.Email, usuario.Perfil)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao gerar token: "+err.Error())
		return
	}

	// Gera um novo refresh token
	refreshToken, err := utils.GerarRefreshToken(usuario.ID, usuario.Email)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao gerar refresh token: "+err.Error())
		return
	}

	// Retorna os novos tokens
	respondJSON(w, http.StatusOK, RespostaToken{
		Token:        token,
		RefreshToken: refreshToken,
		Usuario:      usuario.ParaPublico(),
	})
}

// ServeHTTP implementa a interface http.Handler
func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case path == "/auth/login":
		h.Login(w, r)
	case path == "/auth/registro":
		h.Registro(w, r)
	case path == "/auth/refresh":
		h.RefreshToken(w, r)
	default:
		http.NotFound(w, r)
	}
} 