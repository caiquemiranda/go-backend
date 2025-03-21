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

	"app14/internal/api"
	"app14/internal/config"
	"app14/internal/database"
)

func main() {
	// Carregar configuração
	cfg := config.LoadConfig()
	
	// Iniciar repositório
	taskRepo := database.NewInMemoryTaskRepository()
	
	// Configurar rotas
	router := api.NewRouter(taskRepo)
	handler := router.Setup()
	
	// Configurar servidor
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: handler,
	}
	
	// Canal para notificação de interrupção
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	// Iniciar servidor em uma goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	
	log.Printf("Server started on port %d", cfg.ServerPort)
	log.Printf("API endpoints:")
	log.Printf("- GET    /api/tasks")
	log.Printf("- POST   /api/tasks")
	log.Printf("- GET    /api/tasks/{id}")
	log.Printf("- PUT    /api/tasks/{id}")
	log.Printf("- DELETE /api/tasks/{id}")
	
	// Esperar sinal de interrupção
	<-done
	log.Print("Server stopping...")
	
	// Criar contexto com timeout para shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Tentar shutdown gracioso
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	
	log.Print("Server stopped")
} 