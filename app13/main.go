package main

import (
	"fmt"
	"log"
	"net/http"

	"app13/servidor"
)

func main() {
	// Configurar rotas
	handler := servidor.ConfigurarRotas()

	// Iniciar servidor
	porta := ":8080"
	fmt.Printf("Servidor da Calculadora rodando em http://localhost%s\n", porta)
	fmt.Println("Rotas disponíveis:")
	fmt.Println(" - /soma?a=<numero>&b=<numero>")
	fmt.Println(" - /subtracao?a=<numero>&b=<numero>")
	fmt.Println(" - /multiplicacao?a=<numero>&b=<numero>")
	fmt.Println(" - /divisao?a=<numero>&b=<numero>")
	fmt.Println(" - /raiz?n=<numero>")
	fmt.Println(" - /potencia?a=<base>&b=<expoente>")
	
	log.Fatal(http.ListenAndServe(porta, handler))
} 