package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Usuário simulado para autenticação
var usuarios = map[string]string{
	"admin": "senha123",
	"user":  "123456",
}

// Logger middleware
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		
		// Configurando o logger para incluir timestamp
		logger := log.New(os.Stdout, "HTTP: ", log.LstdFlags)
		logger.Printf("Requisição iniciada - %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		
		// Chamando o próximo handler na cadeia
		next.ServeHTTP(w, r)
		
		// Logando a finalização da requisição com tempo de duração
		logger.Printf("Requisição finalizada - %s %s - levou %v", r.Method, r.RequestURI, time.Since(inicio))
	})
}

// Middleware de autenticação básica
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtendo credenciais (username e password) da requisição
		username, password, ok := r.BasicAuth()
		
		// Verificando se as credenciais foram fornecidas e são válidas
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Acesso Restrito"`)
			http.Error(w, "Autenticação requerida", http.StatusUnauthorized)
			return
		}
		
		// Verificando se o usuário existe e a senha está correta
		expectedPassword, exists := usuarios[username]
		if !exists || expectedPassword != password {
			http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
			return
		}
		
		// Se a autenticação for bem-sucedida, prossiga para o próximo handler
		next.ServeHTTP(w, r)
	})
}

// Handler para rota pública
func publicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Esta é uma rota pública que qualquer pessoa pode acessar!")
}

// Handler para rota privada (protegida por autenticação)
func privateHandler(w http.ResponseWriter, r *http.Request) {
	username, _, _ := r.BasicAuth()
	fmt.Fprintf(w, "Bem-vindo à área restrita, %s!\n", username)
	fmt.Fprintln(w, "Aqui estão dados sensíveis que apenas usuários autenticados podem ver.")
}

func main() {
	// Criando um novo multiplexador
	mux := http.NewServeMux()
	
	// Rota pública - só usa middleware de logging
	mux.Handle("/public", loggerMiddleware(http.HandlerFunc(publicHandler)))
	
	// Rota privada - usa middleware de autenticação e logging
	// A ordem é importante: primeiro logger, depois auth
	mux.Handle("/private", loggerMiddleware(authMiddleware(http.HandlerFunc(privateHandler))))
	
	// Iniciando o servidor
	porta := ":8080"
	fmt.Printf("Servidor rodando em http://localhost%s\n", porta)
	fmt.Println("Rotas disponíveis:")
	fmt.Println(" - /public  (aberta)")
	fmt.Println(" - /private (requer autenticação)")
	fmt.Println("\nCredenciais para teste:")
	fmt.Println(" - Username: admin, Password: senha123")
	fmt.Println(" - Username: user,  Password: 123456")
	
	log.Fatal(http.ListenAndServe(porta, mux))
} 