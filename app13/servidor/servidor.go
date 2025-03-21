package servidor

import (
	"encoding/json"
	"net/http"
	"strconv"

	"app13/calculadora"
)

// Resposta é a estrutura padrão de resposta da API
type Resposta struct {
	Resultado float64 `json:"resultado,omitempty"`
	Erro      string  `json:"erro,omitempty"`
}

// ConfigurarRotas configura as rotas da API
func ConfigurarRotas() http.Handler {
	mux := http.NewServeMux()

	// Rota para somar dois números
	mux.HandleFunc("/soma", somaHandler)

	// Rota para subtrair dois números
	mux.HandleFunc("/subtracao", subtracaoHandler)

	// Rota para multiplicar dois números
	mux.HandleFunc("/multiplicacao", multiplicacaoHandler)

	// Rota para dividir dois números
	mux.HandleFunc("/divisao", divisaoHandler)

	// Rota para calcular raiz quadrada
	mux.HandleFunc("/raiz", raizHandler)

	// Rota para calcular potência
	mux.HandleFunc("/potencia", potenciaHandler)

	return mux
}

// somaHandler processa requisições para somar dois números
func somaHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := obterParametros(r)
	if err != nil {
		responderComErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado := calculadora.Soma(a, b)
	responderComSucesso(w, resultado)
}

// subtracaoHandler processa requisições para subtrair dois números
func subtracaoHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := obterParametros(r)
	if err != nil {
		responderComErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado := calculadora.Subtracao(a, b)
	responderComSucesso(w, resultado)
}

// multiplicacaoHandler processa requisições para multiplicar dois números
func multiplicacaoHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := obterParametros(r)
	if err != nil {
		responderComErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado := calculadora.Multiplicacao(a, b)
	responderComSucesso(w, resultado)
}

// divisaoHandler processa requisições para dividir dois números
func divisaoHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := obterParametros(r)
	if err != nil {
		responderComErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado, err := calculadora.Divisao(a, b)
	if err != nil {
		responderComErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	responderComSucesso(w, resultado)
}

// raizHandler processa requisições para calcular raiz quadrada
func raizHandler(w http.ResponseWriter, r *http.Request) {
	// Obter parâmetro n da query string
	n, err := obterParametroUnico(r, "n")
	if err != nil {
		responderComErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado, err := calculadora.RaizQuadrada(n)
	if err != nil {
		responderComErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	responderComSucesso(w, resultado)
}

// potenciaHandler processa requisições para calcular potência
func potenciaHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := obterParametros(r)
	if err != nil {
		responderComErro(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado := calculadora.Potencia(a, b)
	responderComSucesso(w, resultado)
}

// obterParametros extrai os parâmetros a e b da query string
func obterParametros(r *http.Request) (float64, float64, error) {
	// Obter parâmetros da query string
	query := r.URL.Query()
	aStr := query.Get("a")
	bStr := query.Get("b")

	// Converter para float64
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

// obterParametroUnico extrai um único parâmetro da query string
func obterParametroUnico(r *http.Request, nome string) (float64, error) {
	// Obter parâmetro da query string
	query := r.URL.Query()
	valorStr := query.Get(nome)

	// Converter para float64
	valor, err := strconv.ParseFloat(valorStr, 64)
	if err != nil {
		return 0, err
	}

	return valor, nil
}

// responderComSucesso envia uma resposta JSON de sucesso
func responderComSucesso(w http.ResponseWriter, resultado float64) {
	resposta := Resposta{
		Resultado: resultado,
	}

	// Configurar cabeçalhos e enviar resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resposta)
}

// responderComErro envia uma resposta JSON de erro
func responderComErro(w http.ResponseWriter, mensagem string, statusCode int) {
	resposta := Resposta{
		Erro: mensagem,
	}

	// Configurar cabeçalhos e enviar resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resposta)
} 