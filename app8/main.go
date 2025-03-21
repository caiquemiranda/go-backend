package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Nota representa uma nota no sistema
type Nota struct {
	ID        int       `json:"id"`
	Titulo    string    `json:"titulo"`
	Conteudo  string    `json:"conteudo"`
	Categoria string    `json:"categoria"`
	DataCriacao time.Time `json:"dataCriacao"`
	Arquivada bool      `json:"arquivada"`
}

// NotaRepository gerencia operações de banco de dados para notas
type NotaRepository struct {
	db *sql.DB
}

// Inicializa o banco de dados e cria a tabela se não existir
func inicializarBancoDeDados() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./notas.db")
	if err != nil {
		return nil, err
	}

	// Cria a tabela se não existir
	sqlCriarTabela := `
	CREATE TABLE IF NOT EXISTS notas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titulo TEXT NOT NULL,
		conteudo TEXT,
		categoria TEXT,
		data_criacao DATETIME,
		arquivada BOOLEAN DEFAULT FALSE
	);
	`

	_, err = db.Exec(sqlCriarTabela)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NovoNotaRepository cria um novo repositório de notas
func NovoNotaRepository(db *sql.DB) *NotaRepository {
	return &NotaRepository{db: db}
}

// Criar insere uma nova nota no banco de dados
func (r *NotaRepository) Criar(nota Nota) (Nota, error) {
	sql := `
	INSERT INTO notas (titulo, conteudo, categoria, data_criacao, arquivada)
	VALUES (?, ?, ?, ?, ?)
	`

	// Define a data de criação como agora
	nota.DataCriacao = time.Now()

	// Executa a inserção
	result, err := r.db.Exec(sql, nota.Titulo, nota.Conteudo, nota.Categoria, nota.DataCriacao, nota.Arquivada)
	if err != nil {
		return Nota{}, err
	}

	// Obtém o ID gerado
	ultimoID, err := result.LastInsertId()
	if err != nil {
		return Nota{}, err
	}

	// Atribui o ID à nota
	nota.ID = int(ultimoID)

	return nota, nil
}

// ObterTodas retorna todas as notas
func (r *NotaRepository) ObterTodas() ([]Nota, error) {
	sql := `
	SELECT id, titulo, conteudo, categoria, data_criacao, arquivada 
	FROM notas
	ORDER BY data_criacao DESC
	`

	// Executa a consulta
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Processa os resultados
	var notas []Nota
	for rows.Next() {
		var nota Nota
		err := rows.Scan(&nota.ID, &nota.Titulo, &nota.Conteudo, &nota.Categoria, &nota.DataCriacao, &nota.Arquivada)
		if err != nil {
			return nil, err
		}
		notas = append(notas, nota)
	}

	return notas, nil
}

// ObterPorCategoria retorna notas filtradas por categoria
func (r *NotaRepository) ObterPorCategoria(categoria string) ([]Nota, error) {
	sql := `
	SELECT id, titulo, conteudo, categoria, data_criacao, arquivada 
	FROM notas
	WHERE categoria = ?
	ORDER BY data_criacao DESC
	`

	// Executa a consulta
	rows, err := r.db.Query(sql, categoria)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Processa os resultados
	var notas []Nota
	for rows.Next() {
		var nota Nota
		err := rows.Scan(&nota.ID, &nota.Titulo, &nota.Conteudo, &nota.Categoria, &nota.DataCriacao, &nota.Arquivada)
		if err != nil {
			return nil, err
		}
		notas = append(notas, nota)
	}

	return notas, nil
}

// ObterPorID obtém uma nota pelo ID
func (r *NotaRepository) ObterPorID(id int) (Nota, error) {
	sql := `
	SELECT id, titulo, conteudo, categoria, data_criacao, arquivada 
	FROM notas 
	WHERE id = ?
	`

	// Executa a consulta
	var nota Nota
	err := r.db.QueryRow(sql, id).Scan(&nota.ID, &nota.Titulo, &nota.Conteudo, &nota.Categoria, &nota.DataCriacao, &nota.Arquivada)
	if err != nil {
		if err == sql.ErrNoRows {
			return Nota{}, fmt.Errorf("nota não encontrada: %w", err)
		}
		return Nota{}, err
	}

	return nota, nil
}

// Atualizar modifica uma nota existente
func (r *NotaRepository) Atualizar(id int, nota Nota) (Nota, error) {
	sql := `
	UPDATE notas
	SET titulo = ?, conteudo = ?, categoria = ?, arquivada = ?
	WHERE id = ?
	`

	// Executa a atualização
	_, err := r.db.Exec(sql, nota.Titulo, nota.Conteudo, nota.Categoria, nota.Arquivada, id)
	if err != nil {
		return Nota{}, err
	}

	// Verifica se a nota existe
	result, err := r.ObterPorID(id)
	if err != nil {
		return Nota{}, err
	}

	return result, nil
}

// Remover exclui uma nota pelo ID
func (r *NotaRepository) Remover(id int) error {
	sql := `DELETE FROM notas WHERE id = ?`

	// Executa a remoção
	_, err := r.db.Exec(sql, id)
	return err
}

// ArquivarNota altera o status de arquivamento de uma nota
func (r *NotaRepository) ArquivarNota(id int, arquivar bool) (Nota, error) {
	sql := `
	UPDATE notas
	SET arquivada = ?
	WHERE id = ?
	`

	// Executa a atualização
	_, err := r.db.Exec(sql, arquivar, id)
	if err != nil {
		return Nota{}, err
	}

	// Obtém a nota atualizada
	return r.ObterPorID(id)
}

// Handler para gerenciar as requisições para /notas
func handleNotas(w http.ResponseWriter, r *http.Request, repo *NotaRepository) {
	switch r.Method {
	case http.MethodGet:
		// Verifica se há filtro por categoria
		categoria := r.URL.Query().Get("categoria")
		if categoria != "" {
			notas, err := repo.ObterPorCategoria(categoria)
			if err != nil {
				http.Error(w, "Erro ao obter notas: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(notas)
			return
		}

		// Caso não tenha filtro, retorna todas as notas
		notas, err := repo.ObterTodas()
		if err != nil {
			http.Error(w, "Erro ao obter notas: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notas)

	case http.MethodPost:
		// Decodifica o corpo da requisição
		var nota Nota
		err := json.NewDecoder(r.Body).Decode(&nota)
		if err != nil {
			http.Error(w, "Erro ao decodificar nota: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Cria a nota
		notaCriada, err := repo.Criar(nota)
		if err != nil {
			http.Error(w, "Erro ao criar nota: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Retorna a nota criada
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(notaCriada)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// Handler para gerenciar as requisições para /notas/{id}
func handleNota(w http.ResponseWriter, r *http.Request, repo *NotaRepository) {
	// Extrai o ID da URL
	path := strings.TrimPrefix(r.URL.Path, "/notas/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Obtém a nota pelo ID
		nota, err := repo.ObterPorID(id)
		if err != nil {
			http.Error(w, "Erro ao obter nota: "+err.Error(), http.StatusNotFound)
			return
		}

		// Retorna a nota
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nota)

	case http.MethodPut:
		// Decodifica o corpo da requisição
		var nota Nota
		err := json.NewDecoder(r.Body).Decode(&nota)
		if err != nil {
			http.Error(w, "Erro ao decodificar nota: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Atualiza a nota
		notaAtualizada, err := repo.Atualizar(id, nota)
		if err != nil {
			http.Error(w, "Erro ao atualizar nota: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Retorna a nota atualizada
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notaAtualizada)

	case http.MethodDelete:
		// Remove a nota
		err := repo.Remover(id)
		if err != nil {
			http.Error(w, "Erro ao remover nota: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Retorna sem conteúdo
		w.WriteHeader(http.StatusNoContent)

	case http.MethodPatch:
		// Verifica se o parâmetro arquivada está presente
		arquivarStr := r.URL.Query().Get("arquivada")
		if arquivarStr == "" {
			http.Error(w, "Parâmetro 'arquivada' obrigatório", http.StatusBadRequest)
			return
		}

		// Converte para booleano
		arquivar := arquivarStr == "true"

		// Arquiva/desarquiva a nota
		notaAtualizada, err := repo.ArquivarNota(id, arquivar)
		if err != nil {
			http.Error(w, "Erro ao arquivar nota: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Retorna a nota atualizada
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notaAtualizada)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Inicializa o banco de dados
	db, err := inicializarBancoDeDados()
	if err != nil {
		log.Fatalf("Erro ao inicializar banco de dados: %v", err)
	}
	defer db.Close()

	// Cria o repositório
	repo := NovoNotaRepository(db)

	// Configura os handlers
	http.HandleFunc("/notas", func(w http.ResponseWriter, r *http.Request) {
		handleNotas(w, r, repo)
	})

	http.HandleFunc("/notas/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/notas/" {
			http.NotFound(w, r)
			return
		}
		handleNota(w, r, repo)
	})

	// Inicia o servidor
	fmt.Println("Servidor de notas (SQLite) iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
} 