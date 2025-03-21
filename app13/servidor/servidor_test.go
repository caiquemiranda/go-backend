package servidor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestSomaEndpoint testa o endpoint /soma
func TestSomaEndpoint(t *testing.T) {
	// Criar servidor de teste
	servidor := ConfigurarRotas()
	servidorTeste := httptest.NewServer(servidor)
	defer servidorTeste.Close()

	// Casos de teste
	testes := []struct {
		nome            string
		a, b            string
		statusEsperado  int
		resultadoOuErro string
		temErro         bool
	}{
		{"Valores válidos", "5", "3", http.StatusOK, "8", false},
		{"A inválido", "abc", "3", http.StatusBadRequest, "strconv.ParseFloat", true},
		{"B inválido", "5", "xyz", http.StatusBadRequest, "strconv.ParseFloat", true},
		{"Valores negativos", "-5", "-3", http.StatusOK, "-8", false},
	}

	// Executar os casos de teste
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			// Construir a URL
			url := servidorTeste.URL + "/soma?a=" + teste.a + "&b=" + teste.b

			// Fazer a requisição
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Erro ao fazer requisição: %v", err)
			}
			defer resp.Body.Close()

			// Verificar status code
			if resp.StatusCode != teste.statusEsperado {
				t.Errorf("Status code esperado %d, obtido %d", teste.statusEsperado, resp.StatusCode)
			}

			// Decodificar resposta
			var resposta Resposta
			if err := json.NewDecoder(resp.Body).Decode(&resposta); err != nil {
				t.Fatalf("Erro ao decodificar resposta: %v", err)
			}

			// Verificar resultado ou erro
			if teste.temErro {
				if resposta.Erro == "" {
					t.Errorf("Esperava erro contendo '%s', mas não obteve erro", teste.resultadoOuErro)
				}
			} else {
				resultado := float64(0)
				if resposta.Resultado != 0 {
					resultado = resposta.Resultado
				}
				resultadoEsperado := float64(0)
				json.Unmarshal([]byte(teste.resultadoOuErro), &resultadoEsperado)
				if resultado != resultadoEsperado {
					t.Errorf("Resultado esperado %v, obtido %v", resultadoEsperado, resultado)
				}
			}
		})
	}
}

// TestDivisaoEndpoint testa o endpoint /divisao
func TestDivisaoEndpoint(t *testing.T) {
	// Criar servidor de teste
	servidor := ConfigurarRotas()
	servidorTeste := httptest.NewServer(servidor)
	defer servidorTeste.Close()

	// Casos de teste
	testes := []struct {
		nome            string
		a, b            string
		statusEsperado  int
		resultadoOuErro string
		temErro         bool
	}{
		{"Valores válidos", "10", "2", http.StatusOK, "5", false},
		{"Divisão por zero", "10", "0", http.StatusBadRequest, "divisão por zero", true},
		{"A inválido", "abc", "2", http.StatusBadRequest, "strconv.ParseFloat", true},
		{"B inválido", "10", "xyz", http.StatusBadRequest, "strconv.ParseFloat", true},
	}

	// Executar os casos de teste
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			// Construir a URL
			url := servidorTeste.URL + "/divisao?a=" + teste.a + "&b=" + teste.b

			// Fazer a requisição
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Erro ao fazer requisição: %v", err)
			}
			defer resp.Body.Close()

			// Verificar status code
			if resp.StatusCode != teste.statusEsperado {
				t.Errorf("Status code esperado %d, obtido %d", teste.statusEsperado, resp.StatusCode)
			}

			// Decodificar resposta
			var resposta Resposta
			if err := json.NewDecoder(resp.Body).Decode(&resposta); err != nil {
				t.Fatalf("Erro ao decodificar resposta: %v", err)
			}

			// Verificar resultado ou erro
			if teste.temErro {
				if resposta.Erro == "" || !contains(resposta.Erro, teste.resultadoOuErro) {
					t.Errorf("Esperava erro contendo '%s', obteve '%s'", teste.resultadoOuErro, resposta.Erro)
				}
			} else {
				resultado := float64(0)
				if resposta.Resultado != 0 {
					resultado = resposta.Resultado
				}
				resultadoEsperado := float64(0)
				json.Unmarshal([]byte(teste.resultadoOuErro), &resultadoEsperado)
				if resultado != resultadoEsperado {
					t.Errorf("Resultado esperado %v, obtido %v", resultadoEsperado, resultado)
				}
			}
		})
	}
}

// TestRaizEndpoint testa o endpoint /raiz
func TestRaizEndpoint(t *testing.T) {
	// Criar servidor de teste
	servidor := ConfigurarRotas()
	servidorTeste := httptest.NewServer(servidor)
	defer servidorTeste.Close()

	// Casos de teste
	testes := []struct {
		nome            string
		n               string
		statusEsperado  int
		resultadoOuErro string
		temErro         bool
	}{
		{"Valor válido", "9", http.StatusOK, "3", false},
		{"Valor negativo", "-9", http.StatusBadRequest, "número negativo", true},
		{"Valor inválido", "abc", http.StatusBadRequest, "strconv.ParseFloat", true},
	}

	// Executar os casos de teste
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			// Construir a URL
			url := servidorTeste.URL + "/raiz?n=" + teste.n

			// Fazer a requisição
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Erro ao fazer requisição: %v", err)
			}
			defer resp.Body.Close()

			// Verificar status code
			if resp.StatusCode != teste.statusEsperado {
				t.Errorf("Status code esperado %d, obtido %d", teste.statusEsperado, resp.StatusCode)
			}

			// Decodificar resposta
			var resposta Resposta
			if err := json.NewDecoder(resp.Body).Decode(&resposta); err != nil {
				t.Fatalf("Erro ao decodificar resposta: %v", err)
			}

			// Verificar resultado ou erro
			if teste.temErro {
				if resposta.Erro == "" || !contains(resposta.Erro, teste.resultadoOuErro) {
					t.Errorf("Esperava erro contendo '%s', obteve '%s'", teste.resultadoOuErro, resposta.Erro)
				}
			} else {
				resultado := float64(0)
				if resposta.Resultado != 0 {
					resultado = resposta.Resultado
				}
				resultadoEsperado := float64(0)
				json.Unmarshal([]byte(teste.resultadoOuErro), &resultadoEsperado)
				if resultado != resultadoEsperado {
					t.Errorf("Resultado esperado %v, obtido %v", resultadoEsperado, resultado)
				}
			}
		})
	}
}

// Função auxiliar para verificar se uma string contém outra
func contains(s, substr string) bool {
	return s != "" && substr != "" && len(s) >= len(substr) && s[0:len(substr)] == substr
} 