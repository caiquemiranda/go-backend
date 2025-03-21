package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Estrutura para representar um usuário
type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
}

func main() {
	// Handler para obter usuário em formato JSON
	http.HandleFunc("/usuario", handleUsuario)

	// Handler para criar um novo usuário via JSON
	http.HandleFunc("/criar-usuario", handleCriarUsuario)

	// Informa que o servidor está iniciando
	fmt.Println("Servidor JSON iniciado na porta 8080...")
	
	// Inicia o servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler para retornar um usuário em formato JSON
func handleUsuario(w http.ResponseWriter, r *http.Request) {
	// Apenas aceita método GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Cria um usuário de exemplo
	usuario := Usuario{
		ID:    1,
		Nome:  "João Silva",
		Email: "joao@exemplo.com",
		Idade: 30,
	}

	// Define o Content-Type da resposta
	w.Header().Set("Content-Type", "application/json")
	
	// Codifica o usuário para JSON e envia na resposta
	err := json.NewEncoder(w).Encode(usuario)
	if err != nil {
		http.Error(w, "Erro ao codificar JSON", http.StatusInternalServerError)
		return
	}
}

// Handler para criar um usuário a partir de JSON recebido
func handleCriarUsuario(w http.ResponseWriter, r *http.Request) {
	// Apenas aceita método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Lê o corpo da requisição
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Cria uma variável para armazenar os dados do usuário
	var novoUsuario Usuario

	// Decodifica o JSON para a estrutura de usuário
	err = json.Unmarshal(body, &novoUsuario)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Simula a criação de um ID (em um cenário real seria gerado pelo banco de dados)
	novoUsuario.ID = 999

	// Prepara a resposta JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	// Responde com o usuário criado, incluindo o ID gerado
	json.NewEncoder(w).Encode(novoUsuario)

	// Exibe no console para fins de log
	fmt.Printf("Usuário criado: %+v\n", novoUsuario)
} 