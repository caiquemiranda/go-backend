package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger middleware para registrar informações sobre requisições HTTP
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Criando um wrapper para o ResponseWriter para capturar o status code
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Status padrão
		}

		// Chamar o próximo handler na cadeia
		next.ServeHTTP(rw, r)

		// Calcular duração e registrar no log
		duration := time.Since(start)
		log.Printf(
			"[%s] %s %s %d %v",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			rw.statusCode,
			duration,
		)
	})
}

// responseWriter é um wrapper para http.ResponseWriter que captura o status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader sobrescreve o método original para capturar o status code
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
} 