package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"./handlers"
	"./middlewares"
	"./models"
	"./utils"
)

func main() {
	// Inicializa os repositórios
	repoUsuario := models.NovoRepositorioUsuarioMemoria()
	repoRecurso := models.NovoRepositorioRecursoMemoria()

	// Dados iniciais - Usuários
	criarUsuariosIniciais(repoUsuario)

	// Dados iniciais - Recursos
	criarRecursosIniciais(repoRecurso)

	// Inicializa os handlers
	authHandler := handlers.NovoAuthHandler(repoUsuario)
	recursoHandler := handlers.NovoRecursoHandler(repoRecurso)

	// Configura as rotas
	mux := http.NewServeMux()

	// Rotas públicas
	mux.Handle("/auth/", middlewares.CORS(middlewares.Logger(authHandler)))

	// Rotas protegidas que exigem autenticação
	recursoAutenticado := middlewares.RequererAutenticacao(recursoHandler)
	mux.Handle("/recursos", middlewares.CORS(middlewares.Logger(recursoAutenticado)))
	mux.Handle("/recursos/", middlewares.CORS(middlewares.Logger(recursoAutenticado)))

	// Rota pública para listar recursos públicos (visão limitada)
	mux.HandleFunc("/recursos-publicos", func(w http.ResponseWriter, r *http.Request) {
		// Define filtros para acessar apenas recursos públicos
		filtros := map[string]interface{}{
			"nivelAcesso": 0, // Nível público
		}

		// Busca recursos
		recursos, err := repoRecurso.ObterTodos(filtros)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Erro ao buscar recursos: %v", err)
			return
		}

		// Converte para a versão pública dos recursos
		recursosPublicos := make([]models.RecursoVisaoPublica, 0, len(recursos))
		for _, r := range recursos {
			recursosPublicos = append(recursosPublicos, r.ParaPublico())
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recursosPublicos)
	})

	// Rota para administração (exige autenticação e perfil admin)
	admin := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"mensagem": "Área Administrativa - Acesso restrito ao Admin"}`)
	})
	adminHandler := middlewares.RequererAutorizacao("admin")(admin)
	mux.Handle("/admin", middlewares.CORS(middlewares.Logger(middlewares.RequererAutenticacao(adminHandler))))

	// Rota para área de editores (exige autenticação e perfil editor ou admin)
	editor := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"mensagem": "Área de Editores - Acesso restrito"}`)
	})
	editorHandler := middlewares.RequererAutorizacao("editor")(editor)
	mux.Handle("/editor", middlewares.CORS(middlewares.Logger(middlewares.RequererAutenticacao(editorHandler))))

	// Rota raiz com informações sobre a API
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"api": "API de Demonstração com Autenticação e Middlewares",
			"versao": "1.0.0",
			"rotas": {
				"/auth/login": "Autenticação de usuários (POST)",
				"/auth/registro": "Registro de novos usuários (POST)",
				"/auth/refresh": "Renovação de token (POST)",
				"/recursos": "Gerenciamento de recursos (GET, POST)",
				"/recursos/{id}": "Operações em recursos específicos (GET, PUT, DELETE)",
				"/recursos-publicos": "Lista recursos públicos (GET)",
				"/admin": "Área administrativa (GET - requer perfil admin)",
				"/editor": "Área de editores (GET - requer perfil editor ou admin)"
			}
		}`)
	})

	// Inicia o servidor
	fmt.Println("=== API com Autenticação e Middlewares ===")
	fmt.Println("Servidor iniciado na porta 8080")
	fmt.Println("Usuários iniciais:")
	fmt.Println("  - Admin: admin@exemplo.com / senha123")
	fmt.Println("  - Editor: editor@exemplo.com / senha123")
	fmt.Println("  - Usuário: usuario@exemplo.com / senha123")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// criarUsuariosIniciais adiciona usuários de exemplo ao repositório
func criarUsuariosIniciais(repo *models.RepositorioUsuarioMemoria) {
	// Cria um usuário admin
	admin, _ := models.NovoUsuario(
		"Administrador",
		"admin@exemplo.com",
		"senha123",
		"admin",
	)
	admin.Ativo = true
	repo.Criar(admin)

	// Cria um usuário editor
	editor, _ := models.NovoUsuario(
		"Editor",
		"editor@exemplo.com",
		"senha123",
		"editor",
	)
	editor.Ativo = true
	repo.Criar(editor)

	// Cria um usuário comum
	usuario, _ := models.NovoUsuario(
		"Usuário",
		"usuario@exemplo.com",
		"senha123",
		"usuario",
	)
	usuario.Ativo = true
	repo.Criar(usuario)
}

// criarRecursosIniciais adiciona recursos de exemplo ao repositório
func criarRecursosIniciais(repo *models.RepositorioRecursoMemoria) {
	// Recursos públicos (visíveis para todos)
	recurso1 := models.NovoRecurso(
		"Introdução a API REST",
		"Conceitos básicos sobre APIs REST",
		"Conteúdo detalhado sobre o que são APIs REST e como funcionam...",
		"Programação",
		0, // Nível de acesso: público
		1, // Criado pelo admin (ID 1)
	)
	recurso1.Publicar()
	repo.Criar(recurso1)

	recurso2 := models.NovoRecurso(
		"Uso de Middlewares em Go",
		"Aprenda como implementar middlewares em Go",
		"Neste artigo, exploramos como criar e usar middlewares em Go para...",
		"Programação",
		0, // Nível de acesso: público
		2, // Criado pelo editor (ID 2)
	)
	recurso2.Publicar()
	repo.Criar(recurso2)

	// Recursos para usuários autenticados
	recurso3 := models.NovoRecurso(
		"Autenticação com JWT",
		"Implementação detalhada de JWT em Go",
		"Este guia completo mostra como implementar autenticação JWT em Go...",
		"Segurança",
		1, // Nível de acesso: usuário
		1, // Criado pelo admin (ID 1)
	)
	recurso3.Publicar()
	repo.Criar(recurso3)

	// Recursos para editores
	recurso4 := models.NovoRecurso(
		"Padrões Avançados de API",
		"Padrões de design para APIs escaláveis",
		"Conteúdo detalhado sobre como estruturar APIs de forma escalável...",
		"Arquitetura",
		2, // Nível de acesso: editor
		1, // Criado pelo admin (ID 1)
	)
	recurso4.Publicar()
	repo.Criar(recurso4)

	// Recursos para administradores
	recurso5 := models.NovoRecurso(
		"Configurações de Segurança",
		"Configurações avançadas de segurança",
		"Informações sensíveis sobre configurações de segurança da API...",
		"Segurança",
		3, // Nível de acesso: admin
		1, // Criado pelo admin (ID 1)
	)
	recurso5.Publicar()
	repo.Criar(recurso5)
} 