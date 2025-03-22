package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define o handler para a rota raiz
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Informa que o servidor está iniciando
	fmt.Println("Servidor iniciado na porta 8080...")
	
	// Inicia o servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
} 