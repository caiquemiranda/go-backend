package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Produto é o modelo para os produtos no sistema
type Produto struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Nome        string    `json:"nome" gorm:"size:100;not null"`
	Descricao   string    `json:"descricao" gorm:"type:text"`
	Preco       float64   `json:"preco" gorm:"not null"`
	Estoque     int       `json:"estoque" gorm:"default:0"`
	Categoria   string    `json:"categoria" gorm:"size:50;index"`
	SKU         string    `json:"sku" gorm:"size:20;uniqueIndex"`
	Disponivel  bool      `json:"disponivel" gorm:"default:true"`
	DataCriacao time.Time `json:"dataCriacao" gorm:"autoCreateTime"`
	DataAtualizacao time.Time `json:"dataAtualizacao" gorm:"autoUpdateTime"`
}

// ProdutoRepository é responsável pelas operações de banco de dados com produtos
type ProdutoRepository struct {
	db *gorm.DB
}

// NovoProdutoRepository cria um novo repositório de produtos
func NovoProdutoRepository(db *gorm.DB) *ProdutoRepository {
	return &ProdutoRepository{db: db}
}

// InicializarBancoDeDados configura a conexão com o banco de dados e as migrações
func InicializarBancoDeDados() (*gorm.DB, error) {
	// Abre a conexão com o banco de dados SQLite
	db, err := gorm.Open(sqlite.Open("produtos.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Realiza a migração automática (cria/atualiza tabelas)
	err = db.AutoMigrate(&Produto{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Criar adiciona um novo produto ao banco de dados
func (r *ProdutoRepository) Criar(produto *Produto) error {
	return r.db.Create(produto).Error
}

// ObterTodos retorna todos os produtos
func (r *ProdutoRepository) ObterTodos(filtros map[string]interface{}) ([]Produto, error) {
	var produtos []Produto
	query := r.db

	// Filtra por categoria
	if categoria, ok := filtros["categoria"]; ok && categoria != "" {
		query = query.Where("categoria = ?", categoria)
	}

	// Filtra por disponibilidade
	if disponivel, ok := filtros["disponivel"]; ok {
		query = query.Where("disponivel = ?", disponivel)
	}

	// Filtra por preço mínimo
	if precoMin, ok := filtros["precoMin"]; ok {
		query = query.Where("preco >= ?", precoMin)
	}

	// Filtra por preço máximo
	if precoMax, ok := filtros["precoMax"]; ok {
		query = query.Where("preco <= ?", precoMax)
	}

	// Define a ordenação
	ordem := "data_criacao DESC" // Padrão: ordenado por data de criação (mais recente primeiro)
	if campo, ok := filtros["ordenarPor"]; ok && campo != "" {
		direcao := "ASC"
		if dir, ok := filtros["direcao"]; ok && dir == "desc" {
			direcao = "DESC"
		}
		ordem = fmt.Sprintf("%s %s", campo, direcao)
	}

	// Executa a consulta com os filtros e ordenação
	err := query.Order(ordem).Find(&produtos).Error
	return produtos, err
}

// ObterPorID busca um produto pelo ID
func (r *ProdutoRepository) ObterPorID(id uint) (*Produto, error) {
	var produto Produto
	result := r.db.First(&produto, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produto, nil
}

// ObterPorSKU busca um produto pelo SKU
func (r *ProdutoRepository) ObterPorSKU(sku string) (*Produto, error) {
	var produto Produto
	result := r.db.Where("sku = ?", sku).First(&produto)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produto, nil
}

// Atualizar modifica um produto existente
func (r *ProdutoRepository) Atualizar(id uint, produto *Produto) error {
	return r.db.Model(&Produto{ID: id}).Updates(produto).Error
}

// AtualizarEstoque modifica apenas o estoque de um produto
func (r *ProdutoRepository) AtualizarEstoque(id uint, quantidade int) error {
	return r.db.Model(&Produto{ID: id}).Update("estoque", quantidade).Error
}

// AtualizarDisponibilidade modifica apenas a disponibilidade de um produto
func (r *ProdutoRepository) AtualizarDisponibilidade(id uint, disponivel bool) error {
	return r.db.Model(&Produto{ID: id}).Update("disponivel", disponivel).Error
}

// Remover exclui um produto pelo ID
func (r *ProdutoRepository) Remover(id uint) error {
	return r.db.Delete(&Produto{}, id).Error
}

// RespostaErro representa um erro da API
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
	respondJSON(w, status, RespostaErro{Status: status, Mensagem: mensagem})
}

// handleProdutos gerencia as requisições para a rota /produtos
func handleProdutos(w http.ResponseWriter, r *http.Request, repo *ProdutoRepository) {
	switch r.Method {
	case http.MethodGet:
		// Extrai os parâmetros de filtro
		filtros := make(map[string]interface{})
		
		if categoria := r.URL.Query().Get("categoria"); categoria != "" {
			filtros["categoria"] = categoria
		}
		
		if disponivel := r.URL.Query().Get("disponivel"); disponivel != "" {
			filtros["disponivel"] = disponivel == "true"
		}
		
		if precoMin := r.URL.Query().Get("precoMin"); precoMin != "" {
			if valor, err := strconv.ParseFloat(precoMin, 64); err == nil {
				filtros["precoMin"] = valor
			}
		}
		
		if precoMax := r.URL.Query().Get("precoMax"); precoMax != "" {
			if valor, err := strconv.ParseFloat(precoMax, 64); err == nil {
				filtros["precoMax"] = valor
			}
		}
		
		if ordenarPor := r.URL.Query().Get("ordenarPor"); ordenarPor != "" {
			filtros["ordenarPor"] = ordenarPor
		}
		
		if direcao := r.URL.Query().Get("direcao"); direcao != "" {
			filtros["direcao"] = direcao
		}

		// Obtém os produtos com os filtros aplicados
		produtos, err := repo.ObterTodos(filtros)
		if err != nil {
			respondErro(w, http.StatusInternalServerError, "Erro ao obter produtos: "+err.Error())
			return
		}
		
		respondJSON(w, http.StatusOK, produtos)

	case http.MethodPost:
		var produto Produto
		
		// Decodifica o corpo da requisição
		err := json.NewDecoder(r.Body).Decode(&produto)
		if err != nil {
			respondErro(w, http.StatusBadRequest, "Erro ao decodificar produto: "+err.Error())
			return
		}
		
		// Valida os campos obrigatórios
		if produto.Nome == "" {
			respondErro(w, http.StatusBadRequest, "O nome do produto é obrigatório")
			return
		}
		
		if produto.Preco <= 0 {
			respondErro(w, http.StatusBadRequest, "O preço deve ser maior que zero")
			return
		}
		
		// Cria o produto
		err = repo.Criar(&produto)
		if err != nil {
			respondErro(w, http.StatusInternalServerError, "Erro ao criar produto: "+err.Error())
			return
		}
		
		respondJSON(w, http.StatusCreated, produto)

	default:
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
}

// handleProduto gerencia as requisições para a rota /produtos/{id}
func handleProduto(w http.ResponseWriter, r *http.Request, repo *ProdutoRepository) {
	// Extrai o ID da URL
	path := r.URL.Path[len("/produtos/"):]
	idUint, err := strconv.ParseUint(path, 10, 64)
	if err != nil {
		respondErro(w, http.StatusBadRequest, "ID inválido")
		return
	}
	
	id := uint(idUint)

	switch r.Method {
	case http.MethodGet:
		produto, err := repo.ObterPorID(id)
		if err != nil {
			respondErro(w, http.StatusNotFound, "Produto não encontrado")
			return
		}
		
		respondJSON(w, http.StatusOK, produto)

	case http.MethodPut:
		var produto Produto
		
		// Decodifica o corpo da requisição
		err := json.NewDecoder(r.Body).Decode(&produto)
		if err != nil {
			respondErro(w, http.StatusBadRequest, "Erro ao decodificar produto: "+err.Error())
			return
		}
		
		// Define o ID do produto sendo atualizado
		produto.ID = id
		
		// Atualiza o produto
		err = repo.Atualizar(id, &produto)
		if err != nil {
			respondErro(w, http.StatusInternalServerError, "Erro ao atualizar produto: "+err.Error())
			return
		}
		
		// Obtém o produto atualizado
		produtoAtualizado, err := repo.ObterPorID(id)
		if err != nil {
			respondErro(w, http.StatusInternalServerError, "Erro ao obter produto atualizado: "+err.Error())
			return
		}
		
		respondJSON(w, http.StatusOK, produtoAtualizado)

	case http.MethodDelete:
		// Remove o produto
		err := repo.Remover(id)
		if err != nil {
			respondErro(w, http.StatusInternalServerError, "Erro ao remover produto: "+err.Error())
			return
		}
		
		w.WriteHeader(http.StatusNoContent)

	case http.MethodPatch:
		// Verifica quais campos serão atualizados
		
		// Atualiza apenas o estoque
		if quantidade := r.URL.Query().Get("estoque"); quantidade != "" {
			estoque, err := strconv.Atoi(quantidade)
			if err != nil {
				respondErro(w, http.StatusBadRequest, "Valor de estoque inválido")
				return
			}
			
			err = repo.AtualizarEstoque(id, estoque)
			if err != nil {
				respondErro(w, http.StatusInternalServerError, "Erro ao atualizar estoque: "+err.Error())
				return
			}
			
			produto, err := repo.ObterPorID(id)
			if err != nil {
				respondErro(w, http.StatusInternalServerError, "Erro ao obter produto atualizado: "+err.Error())
				return
			}
			
			respondJSON(w, http.StatusOK, produto)
			return
		}
		
		// Atualiza apenas a disponibilidade
		if disponivel := r.URL.Query().Get("disponivel"); disponivel != "" {
			disponibilidade := disponivel == "true"
			
			err := repo.AtualizarDisponibilidade(id, disponibilidade)
			if err != nil {
				respondErro(w, http.StatusInternalServerError, "Erro ao atualizar disponibilidade: "+err.Error())
				return
			}
			
			produto, err := repo.ObterPorID(id)
			if err != nil {
				respondErro(w, http.StatusInternalServerError, "Erro ao obter produto atualizado: "+err.Error())
				return
			}
			
			respondJSON(w, http.StatusOK, produto)
			return
		}
		
		respondErro(w, http.StatusBadRequest, "É necessário especificar quais campos atualizar")

	default:
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
}

// handleProdutoPorSKU gerencia as requisições para a rota /produtos/sku/{sku}
func handleProdutoPorSKU(w http.ResponseWriter, r *http.Request, repo *ProdutoRepository) {
	// Extrai o SKU da URL
	path := r.URL.Path[len("/produtos/sku/"):]
	
	if path == "" {
		respondErro(w, http.StatusBadRequest, "SKU inválido")
		return
	}

	if r.Method != http.MethodGet {
		respondErro(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	produto, err := repo.ObterPorSKU(path)
	if err != nil {
		respondErro(w, http.StatusNotFound, "Produto não encontrado")
		return
	}
	
	respondJSON(w, http.StatusOK, produto)
}

// PopularDadosIniciais insere alguns produtos de exemplo no banco
func PopularDadosIniciais(repo *ProdutoRepository) {
	produtos := []Produto{
		{
			Nome:       "Notebook Dell Inspiron 15",
			Descricao:  "Notebook com processador Intel Core i5, 8GB RAM, 256GB SSD",
			Preco:      3500.00,
			Estoque:    10,
			Categoria:  "Notebooks",
			SKU:        "NTBK-DELL-001",
			Disponivel: true,
		},
		{
			Nome:       "Smartphone Samsung Galaxy S21",
			Descricao:  "Smartphone com tela AMOLED 6.2\", 128GB armazenamento, 8GB RAM",
			Preco:      2800.00,
			Estoque:    15,
			Categoria:  "Smartphones",
			SKU:        "SMRTPH-SAM-001",
			Disponivel: true,
		},
		{
			Nome:       "Monitor LG UltraWide 29\"",
			Descricao:  "Monitor ultrawide com resolução 2560x1080, 29 polegadas",
			Preco:      1200.00,
			Estoque:    5,
			Categoria:  "Monitores",
			SKU:        "MONI-LG-001",
			Disponivel: true,
		},
		{
			Nome:       "Teclado Mecânico Logitech G Pro",
			Descricao:  "Teclado mecânico para jogos com switches GX Blue",
			Preco:      450.00,
			Estoque:    20,
			Categoria:  "Periféricos",
			SKU:        "TECL-LOG-001",
			Disponivel: true,
		},
		{
			Nome:       "Mouse Gamer Razer DeathAdder",
			Descricao:  "Mouse com 20.000 DPI, 5 botões programáveis, iluminação RGB",
			Preco:      300.00,
			Estoque:    30,
			Categoria:  "Periféricos",
			SKU:        "MOUS-RAZ-001",
			Disponivel: true,
		},
	}

	// Verifica se cada produto já existe e adiciona se não existir
	for _, produto := range produtos {
		existente, err := repo.ObterPorSKU(produto.SKU)
		if err != nil || existente == nil {
			repo.Criar(&produto)
		}
	}
}

func main() {
	// Inicializa o banco de dados
	db, err := InicializarBancoDeDados()
	if err != nil {
		log.Fatalf("Erro ao inicializar banco de dados: %v", err)
	}

	// Cria o repositório
	repo := NovoProdutoRepository(db)

	// Popula dados iniciais
	PopularDadosIniciais(repo)

	// Configura os handlers
	http.HandleFunc("/produtos", func(w http.ResponseWriter, r *http.Request) {
		handleProdutos(w, r, repo)
	})

	http.HandleFunc("/produtos/", func(w http.ResponseWriter, r *http.Request) {
		// Verifica se é uma rota para buscar por SKU
		if len(r.URL.Path) >= 14 && r.URL.Path[:14] == "/produtos/sku/" {
			handleProdutoPorSKU(w, r, repo)
			return
		}

		// Se não for a rota raiz, trata como acesso a um produto específico
		if r.URL.Path != "/produtos/" {
			handleProduto(w, r, repo)
			return
		}

		http.NotFound(w, r)
	})

	// Inicia o servidor
	fmt.Println("Servidor de produtos (GORM) iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
} 