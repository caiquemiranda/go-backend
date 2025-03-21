package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Produto representa um item na loja
type Produto struct {
	ID       int     `json:"id"`
	Nome     string  `json:"nome"`
	Preco    float64 `json:"preco"`
	Estoque  int     `json:"estoque"`
	Categoria string  `json:"categoria"`
}

// Armazenamento em memória para os produtos
var produtos = map[int]Produto{
	1: {ID: 1, Nome: "Notebook", Preco: 3500.00, Estoque: 10, Categoria: "Eletrônicos"},
	2: {ID: 2, Nome: "Mouse", Preco: 50.00, Estoque: 20, Categoria: "Periféricos"},
	3: {ID: 3, Nome: "Teclado", Preco: 100.00, Estoque: 15, Categoria: "Periféricos"},
}

// ContadorID para gerar IDs únicos
var contadorID = 4

func main() {
	// Rota para listar todos os produtos
	http.HandleFunc("/produtos", handleProdutos)
	
	// Rota para manipular um produto específico pelo ID
	http.HandleFunc("/produtos/", handleProduto)
	
	// Rota para listar produtos por categoria
	http.HandleFunc("/categorias/", handleCategorias)
	
	fmt.Println("Servidor de produtos iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleProdutos lida com a listagem e criação de produtos
func handleProdutos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Converte o map para slice para facilitar a serialização
		listaProdutos := make([]Produto, 0, len(produtos))
		for _, produto := range produtos {
			listaProdutos = append(listaProdutos, produto)
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(listaProdutos)
	
	case http.MethodPost:
		// Cria um novo produto
		var novoProduto Produto
		err := json.NewDecoder(r.Body).Decode(&novoProduto)
		if err != nil {
			http.Error(w, "Erro ao decodificar produto", http.StatusBadRequest)
			return
		}
		
		// Atribui um ID único e adiciona ao map
		novoProduto.ID = contadorID
		contadorID++
		produtos[novoProduto.ID] = novoProduto
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(novoProduto)
	
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// handleProduto lida com operações em um produto específico
func handleProduto(w http.ResponseWriter, r *http.Request) {
	// Extrair ID do caminho da URL
	// URL esperada: /produtos/{id}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		http.Error(w, "Caminho inválido", http.StatusBadRequest)
		return
	}
	
	idStr := pathParts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	// Verificar se o produto existe
	produto, existe := produtos[id]
	if !existe {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		// Retorna o produto especificado
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(produto)
	
	case http.MethodPut:
		// Atualiza o produto
		var produtoAtualizado Produto
		err := json.NewDecoder(r.Body).Decode(&produtoAtualizado)
		if err != nil {
			http.Error(w, "Erro ao decodificar produto", http.StatusBadRequest)
			return
		}
		
		// Mantém o ID original
		produtoAtualizado.ID = id
		produtos[id] = produtoAtualizado
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(produtoAtualizado)
	
	case http.MethodDelete:
		// Remove o produto
		delete(produtos, id)
		w.WriteHeader(http.StatusNoContent)
	
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// handleCategorias lida com a listagem de produtos por categoria
func handleCategorias(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	// Extrair categoria do caminho da URL
	// URL esperada: /categorias/{categoria}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		http.Error(w, "Caminho inválido", http.StatusBadRequest)
		return
	}
	
	categoria := pathParts[2]
	
	// Filtra produtos pela categoria
	produtosFiltrados := make([]Produto, 0)
	for _, produto := range produtos {
		if strings.EqualFold(produto.Categoria, categoria) {
			produtosFiltrados = append(produtosFiltrados, produto)
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtosFiltrados)
} 