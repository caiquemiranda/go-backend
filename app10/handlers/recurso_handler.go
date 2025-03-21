package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"../middlewares"
	"../models"
	"../utils"
)

// RepositorioRecurso define a interface para acessar dados de recursos
type RepositorioRecurso interface {
	ObterTodos(filtros map[string]interface{}) ([]*models.Recurso, error)
	ObterPorID(id uint) (*models.Recurso, error)
	Criar(recurso *models.Recurso) error
	Atualizar(id uint, recurso *models.Recurso) error
	Remover(id uint) error
}

// RecursoHandler gerencia as rotas relacionadas a recursos
type RecursoHandler struct {
	repo RepositorioRecurso
}

// NovoRecursoHandler cria uma nova instância do handler de recursos
func NovoRecursoHandler(repo RepositorioRecurso) *RecursoHandler {
	return &RecursoHandler{
		repo: repo,
	}
}

// ObterRecursos retorna todos os recursos acessíveis ao usuário
func (h *RecursoHandler) ObterRecursos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	// Obtém o nível de acesso do usuário atual
	var nivelAcesso int
	if userRole, ok := r.Context().Value(middlewares.UserRoleKey).(string); ok {
		nivelAcesso = utils.NivelAcessoParaRole(userRole)
	} else {
		// Se não houver usuário autenticado, considera nível público
		nivelAcesso = 0
	}

	// Configura filtros
	filtros := make(map[string]interface{})
	
	// Filtra por nível de acesso
	filtros["nivelAcesso"] = nivelAcesso
	
	// Filtra por categoria
	if categoria := r.URL.Query().Get("categoria"); categoria != "" {
		filtros["categoria"] = categoria
	}
	
	// Busca recursos
	recursos, err := h.repo.ObterTodos(filtros)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao buscar recursos: "+err.Error())
		return
	}

	// Para usuários não autenticados ou com nível de acesso baixo,
	// converte para a versão pública dos recursos
	if nivelAcesso < 2 {
		recursosPublicos := make([]models.RecursoVisaoPublica, 0, len(recursos))
		for _, r := range recursos {
			recursosPublicos = append(recursosPublicos, r.ParaPublico())
		}
		respondJSON(w, http.StatusOK, recursosPublicos)
		return
	}

	// Para usuários com nível de acesso alto, retorna os recursos completos
	respondJSON(w, http.StatusOK, recursos)
}

// ObterRecurso retorna um recurso específico
func (h *RecursoHandler) ObterRecurso(w http.ResponseWriter, r *http.Request, id uint) {
	if r.Method != http.MethodGet {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	// Busca o recurso
	recurso, err := h.repo.ObterPorID(id)
	if err != nil {
		respondErro(w, http.StatusNotFound, "Recurso não encontrado")
		return
	}

	// Obtém o nível de acesso do usuário atual
	var nivelAcesso int
	if userRole, ok := r.Context().Value(middlewares.UserRoleKey).(string); ok {
		nivelAcesso = utils.NivelAcessoParaRole(userRole)
	} else {
		nivelAcesso = 0
	}

	// Verifica se o usuário tem permissão para acessar este recurso
	if nivelAcesso < recurso.AcessoLevel {
		respondErro(w, http.StatusForbidden, "Acesso negado a este recurso")
		return
	}

	// Para usuários com nível de acesso baixo, retorna a versão pública
	if nivelAcesso < 2 {
		respondJSON(w, http.StatusOK, recurso.ParaPublico())
		return
	}

	// Para usuários com nível de acesso alto, retorna o recurso completo
	respondJSON(w, http.StatusOK, recurso)
}

// CriarRecurso cria um novo recurso
func (h *RecursoHandler) CriarRecurso(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	// Obtém o ID do usuário do contexto
	userIDValue := r.Context().Value(middlewares.UserIDKey)
	if userIDValue == nil {
		respondErro(w, http.StatusUnauthorized, "Usuário não autenticado")
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		respondErro(w, http.StatusInternalServerError, "Erro ao obter ID do usuário")
		return
	}

	// Estrutura para decodificar a solicitação
	type SolicitacaoRecurso struct {
		Titulo      string `json:"titulo"`
		Descricao   string `json:"descricao"`
		Conteudo    string `json:"conteudo"`
		Categoria   string `json:"categoria"`
		AcessoLevel int    `json:"acessoLevel"`
	}

	var solicita SolicitacaoRecurso
	err := json.NewDecoder(r.Body).Decode(&solicita)
	if err != nil {
		respondErro(w, http.StatusBadRequest, "Erro ao decodificar JSON: "+err.Error())
		return
	}

	// Valida campos obrigatórios
	if solicita.Titulo == "" {
		respondErro(w, http.StatusBadRequest, "Título é obrigatório")
		return
	}

	// Obtém o papel do usuário para verificar permissões de nível de acesso
	var nivelAcessoUsuario int
	if userRole, ok := r.Context().Value(middlewares.UserRoleKey).(string); ok {
		nivelAcessoUsuario = utils.NivelAcessoParaRole(userRole)
	} else {
		nivelAcessoUsuario = 0
	}

	// Não permite que um usuário crie um recurso com nível de acesso superior ao seu
	if solicita.AcessoLevel > nivelAcessoUsuario {
		solicita.AcessoLevel = nivelAcessoUsuario
	}

	// Cria o novo recurso
	recurso := models.NovoRecurso(
		solicita.Titulo,
		solicita.Descricao,
		solicita.Conteudo,
		solicita.Categoria,
		solicita.AcessoLevel,
		userID,
	)

	// Salva o recurso
	err = h.repo.Criar(recurso)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao salvar recurso: "+err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, recurso)
}

// AtualizarRecurso atualiza um recurso existente
func (h *RecursoHandler) AtualizarRecurso(w http.ResponseWriter, r *http.Request, id uint) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	// Busca o recurso existente
	recurso, err := h.repo.ObterPorID(id)
	if err != nil {
		respondErro(w, http.StatusNotFound, "Recurso não encontrado")
		return
	}

	// Obtém o ID e papel do usuário para verificações de permissão
	userIDValue := r.Context().Value(middlewares.UserIDKey)
	userRoleValue := r.Context().Value(middlewares.UserRoleKey)
	
	if userIDValue == nil || userRoleValue == nil {
		respondErro(w, http.StatusUnauthorized, "Usuário não autenticado")
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		respondErro(w, http.StatusInternalServerError, "Erro ao obter ID do usuário")
		return
	}

	userRole, ok := userRoleValue.(string)
	if !ok {
		respondErro(w, http.StatusInternalServerError, "Erro ao obter papel do usuário")
		return
	}

	// Verifica se o usuário é o proprietário do recurso ou um administrador/editor
	isAdmin := userRole == "admin"
	isEditor := userRole == "editor"
	isOwner := recurso.CriadoPor == userID

	if !isAdmin && !isEditor && !isOwner {
		respondErro(w, http.StatusForbidden, "Você não tem permissão para atualizar este recurso")
		return
	}

	// Estrutura para decodificar a solicitação
	type SolicitacaoAtualizacao struct {
		Titulo      string `json:"titulo"`
		Descricao   string `json:"descricao"`
		Conteudo    string `json:"conteudo"`
		Categoria   string `json:"categoria"`
		AcessoLevel int    `json:"acessoLevel"`
		Publicado   bool   `json:"publicado"`
	}

	var solicita SolicitacaoAtualizacao
	err = json.NewDecoder(r.Body).Decode(&solicita)
	if err != nil {
		respondErro(w, http.StatusBadRequest, "Erro ao decodificar JSON: "+err.Error())
		return
	}

	// Atualiza os campos se fornecidos
	recurso.AtualizarConteudo(
		solicita.Titulo,
		solicita.Descricao,
		solicita.Conteudo,
		solicita.Categoria,
	)

	// Apenas administradores e editores podem alterar o nível de acesso e status de publicação
	if isAdmin || isEditor {
		if r.Method == http.MethodPut || solicita.AcessoLevel != 0 {
			recurso.AlterarNivelAcesso(solicita.AcessoLevel)
		}

		// Atualiza o status de publicação se especificado
		if r.Method == http.MethodPut || r.URL.Query().Get("publicar") != "" {
			if solicita.Publicado {
				recurso.Publicar()
			} else {
				recurso.Despublicar()
			}
		}
	}

	// Salva as alterações
	err = h.repo.Atualizar(id, recurso)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao atualizar recurso: "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, recurso)
}

// RemoverRecurso remove um recurso
func (h *RecursoHandler) RemoverRecurso(w http.ResponseWriter, r *http.Request, id uint) {
	if r.Method != http.MethodDelete {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	// Busca o recurso existente
	recurso, err := h.repo.ObterPorID(id)
	if err != nil {
		respondErro(w, http.StatusNotFound, "Recurso não encontrado")
		return
	}

	// Obtém o ID e papel do usuário para verificações de permissão
	userIDValue := r.Context().Value(middlewares.UserIDKey)
	userRoleValue := r.Context().Value(middlewares.UserRoleKey)
	
	if userIDValue == nil || userRoleValue == nil {
		respondErro(w, http.StatusUnauthorized, "Usuário não autenticado")
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		respondErro(w, http.StatusInternalServerError, "Erro ao obter ID do usuário")
		return
	}

	userRole, ok := userRoleValue.(string)
	if !ok {
		respondErro(w, http.StatusInternalServerError, "Erro ao obter papel do usuário")
		return
	}

	// Apenas o proprietário e admin podem remover o recurso
	isAdmin := userRole == "admin"
	isOwner := recurso.CriadoPor == userID

	if !isAdmin && !isOwner {
		respondErro(w, http.StatusForbidden, "Você não tem permissão para remover este recurso")
		return
	}

	// Remove o recurso
	err = h.repo.Remover(id)
	if err != nil {
		respondErro(w, http.StatusInternalServerError, "Erro ao remover recurso: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ServeHTTP implementa a interface http.Handler
func (h *RecursoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Rota para listar ou criar recursos
	if path == "/recursos" {
		switch r.Method {
		case http.MethodGet:
			h.ObterRecursos(w, r)
		case http.MethodPost:
			h.CriarRecurso(w, r)
		default:
			respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		}
		return
	}

	// Rota para operações em um recurso específico
	if strings.HasPrefix(path, "/recursos/") {
		// Extrai o ID do recurso
		idStr := strings.TrimPrefix(path, "/recursos/")
		id64, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondErro(w, http.StatusBadRequest, "ID de recurso inválido")
			return
		}
		id := uint(id64)

		switch r.Method {
		case http.MethodGet:
			h.ObterRecurso(w, r, id)
		case http.MethodPut, http.MethodPatch:
			h.AtualizarRecurso(w, r, id)
		case http.MethodDelete:
			h.RemoverRecurso(w, r, id)
		default:
			respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		}
		return
	}

	http.NotFound(w, r)
} 