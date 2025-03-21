package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers"
	"./models"
	"./services"
)

func main() {
	// Cria as instâncias dos serviços
	livroService := services.NovoLivroService()

	// Adiciona alguns livros de exemplo
	adicionarLivrosExemplo(livroService)

	// Cria as instâncias dos handlers
	livroHandler := handlers.NovoLivroHandler(livroService)

	// Configura as rotas
	http.HandleFunc("/livros", livroHandler.LivrosHandler)
	http.HandleFunc("/livros/", livroHandler.LivroHandler)

	// Inicia o servidor
	fmt.Println("Servidor de biblioteca iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// adicionarLivrosExemplo adiciona alguns livros para teste
func adicionarLivrosExemplo(service *services.LivroService) {
	livros := []*models.Livro{
		models.NovoLivro("O Senhor dos Anéis", "J.R.R. Tolkien", "HarperCollins", "9788595084759", 2001),
		models.NovoLivro("Harry Potter e a Pedra Filosofal", "J.K. Rowling", "Rocco", "9788532511010", 2000),
		models.NovoLivro("O Pequeno Príncipe", "Antoine de Saint-Exupéry", "Agir", "9788525406989", 2009),
		models.NovoLivro("Dom Quixote", "Miguel de Cervantes", "Penguin", "9788582850350", 2016),
		models.NovoLivro("1984", "George Orwell", "Companhia das Letras", "9788535914849", 2009),
	}

	for _, livro := range livros {
		service.Criar(livro)
	}
} 