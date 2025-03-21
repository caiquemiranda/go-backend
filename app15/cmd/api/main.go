package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpAdapter "app15/internal/adapters/http"
)

func main() {
	// Inicializar o router
	router := httpAdapter.SetupRouter()
	
	// Definir porta do servidor
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	// Configurar o servidor HTTP
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Iniciar o servidor em uma goroutine
	go func() {
		fmt.Printf("Servidor iniciado na porta %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v\n", err)
		}
	}()

	// Configurar canal para capturar sinais de término
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Desligando o servidor...")
	
	// Criar contexto com timeout para desligamento gracioso
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Tentar desligar o servidor de forma graciosa
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Erro ao desligar o servidor: %v\n", err)
	}
	
	fmt.Println("Servidor encerrado com sucesso")
} 