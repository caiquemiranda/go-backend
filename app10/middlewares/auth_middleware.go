package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"../utils"
)

// Chaves para o contexto da requisição
type contextKey string

const (
	UserIDKey   contextKey = "userID"
	UserRoleKey contextKey = "userRole"
	UserEmailKey contextKey = "userEmail"
)

// RespostaErro representa uma resposta de erro da API
type RespostaErro struct {
	Status  int    `json:"status"`
	Mensagem string `json:"mensagem"`
}

// respondErro envia uma resposta de erro em formato JSON
func respondErro(w http.ResponseWriter, status int, mensagem string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(RespostaErro{
		Status:  status,
		Mensagem: mensagem,
	})
}

// RequererAutenticacao verifica se o usuário está autenticado
func RequererAutenticacao(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrai o token da requisição
		tokenString, err := utils.ExtrairTokenDaRequisicao(r)
		if err != nil {
			respondErro(w, http.StatusUnauthorized, "Não autorizado: "+err.Error())
			return
		}

		// Valida o token
		claims, err := utils.ValidarToken(tokenString)
		if err != nil {
			respondErro(w, http.StatusUnauthorized, "Token inválido: "+err.Error())
			return
		}

		// Adiciona as informações do usuário ao contexto
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)

		// Continua para o próximo handler com o contexto atualizado
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequererAutorizacao verifica se o usuário tem o papel necessário
func RequererAutorizacao(papel string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Obtém o papel do usuário do contexto
			roleValue := r.Context().Value(UserRoleKey)
			if roleValue == nil {
				respondErro(w, http.StatusUnauthorized, "Não autorizado: informações de usuário ausentes")
				return
			}

			role, ok := roleValue.(string)
			if !ok {
				respondErro(w, http.StatusInternalServerError, "Erro ao obter papel do usuário")
				return
			}

			// Verifica se o papel do usuário atende ao requisito mínimo
			if !utils.VerificarRole(role, papel) {
				respondErro(w, http.StatusForbidden, "Acesso negado: permissões insuficientes")
				return
			}

			// Continua para o próximo handler
			next.ServeHTTP(w, r)
		})
	})
}

// Logger é um middleware que registra informações sobre a requisição
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Aqui poderíamos usar um pacote de logging como o logrus ou zap
		// para registrar detalhes da requisição
		
		// Exemplo simples sem dependências externas:
		method := r.Method
		path := r.URL.Path
		
		// Registra a requisição (em uma aplicação real, usaríamos um logger adequado)
		// log.Printf("%s %s", method, path)
		
		// Define um header personalizado para demonstração
		w.Header().Set("X-Log-Info", method+" "+path)
		
		// Continua para o próximo handler
		next.ServeHTTP(w, r)
	})
}

// CORS é um middleware que configura os cabeçalhos CORS
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define cabeçalhos para permitir CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		
		// Para requisições OPTIONS (preflight), retorna imediatamente
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Continua para o próximo handler
		next.ServeHTTP(w, r)
	})
}

// RateLimiter limita o número de requisições por IP (implementação simplificada)
func RateLimiter(next http.Handler) http.Handler {
	// Em uma implementação real, usaríamos um sistema de cache como Redis
	// para controlar as requisições por IP em um ambiente distribuído
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtém o IP do cliente
		ip := r.RemoteAddr
		
		// Aqui verificaríamos se o IP já atingiu o limite
		// Em uma implementação real, usaríamos algo como:
		// count, _ := redis.Get("rate_limit:" + ip)
		// if count > MAX_REQUESTS {
		//     respondErro(w, http.StatusTooManyRequests, "Limite de requisições excedido")
		//     return
		// }
		// redis.Incr("rate_limit:" + ip)
		// redis.Expire("rate_limit:" + ip, 1 * time.Hour)
		
		// Para demonstração, apenas definimos um header
		w.Header().Set("X-Rate-Limit-IP", ip)
		
		// Continua para o próximo handler
		next.ServeHTTP(w, r)
	})
} 