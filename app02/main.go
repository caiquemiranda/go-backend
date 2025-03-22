package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handler para rota raiz (GET)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "Página inicial - Método GET")
	})

	// Handler para rota /usuarios (GET)
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "Lista de usuários - Método GET")
		} else if r.Method == http.MethodPost {
			fmt.Fprintf(w, "Usuário criado com sucesso - Método POST")
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Handler para rota /produtos (GET e POST)
	http.HandleFunc("/produtos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Lista de produtos - Método GET")
		case http.MethodPost:
			fmt.Fprintf(w, "Produto criado com sucesso - Método POST")
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	// Informa que o servidor está iniciando
	fmt.Println("Servidor iniciado na porta 8080...")
	
	// Inicia o servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
} 