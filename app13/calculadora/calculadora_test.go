package calculadora

import (
	"math"
	"testing"
)

// TestSoma testa a função Soma
func TestSoma(t *testing.T) {
	// Casos de teste
	testes := []struct {
		nome     string
		a, b     float64
		esperado float64
	}{
		{"Positivos", 2, 3, 5},
		{"Negativo e Positivo", -2, 3, 1},
		{"Negativos", -2, -3, -5},
		{"Zero", 0, 0, 0},
		{"Decimal", 1.5, 2.5, 4.0},
	}

	// Executar os casos de teste
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			resultado := Soma(teste.a, teste.b)
			if resultado != teste.esperado {
				t.Errorf("Soma(%v, %v) = %v; esperado %v", 
					teste.a, teste.b, resultado, teste.esperado)
			}
		})
	}
}

// TestSubtracao testa a função Subtracao
func TestSubtracao(t *testing.T) {
	// Casos de teste
	testes := []struct {
		nome     string
		a, b     float64
		esperado float64
	}{
		{"Positivos", 5, 3, 2},
		{"Resultado Negativo", 2, 3, -1},
		{"Negativo e Positivo", -2, 3, -5},
		{"Negativos", -2, -3, 1},
		{"Zero", 0, 0, 0},
	}

	// Executar os casos de teste
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			resultado := Subtracao(teste.a, teste.b)
			if resultado != teste.esperado {
				t.Errorf("Subtracao(%v, %v) = %v; esperado %v", 
					teste.a, teste.b, resultado, teste.esperado)
			}
		})
	}
}

// TestMultiplicacao testa a função Multiplicacao
func TestMultiplicacao(t *testing.T) {
	// Casos de teste
	testes := []struct {
		nome     string
		a, b     float64
		esperado float64
	}{
		{"Positivos", 2, 3, 6},
		{"Negativo e Positivo", -2, 3, -6},
		{"Negativos", -2, -3, 6},
		{"Zero", 5, 0, 0},
		{"Um", 5, 1, 5},
	}

	// Executar os casos de teste
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			resultado := Multiplicacao(teste.a, teste.b)
			if resultado != teste.esperado {
				t.Errorf("Multiplicacao(%v, %v) = %v; esperado %v", 
					teste.a, teste.b, resultado, teste.esperado)
			}
		})
	}
}

// TestDivisao testa a função Divisao
func TestDivisao(t *testing.T) {
	// Casos de teste bem-sucedidos
	testes := []struct {
		nome     string
		a, b     float64
		esperado float64
	}{
		{"Positivos", 6, 3, 2},
		{"Negativo e Positivo", -6, 3, -2},
		{"Negativos", -6, -3, 2},
		{"Decimal", 5.5, 2, 2.75},
		{"Um", 5, 1, 5},
	}

	// Executar os casos de teste bem-sucedidos
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			resultado, err := Divisao(teste.a, teste.b)
			if err != nil {
				t.Errorf("Divisao(%v, %v) retornou erro inesperado: %v", 
					teste.a, teste.b, err)
			}
			if resultado != teste.esperado {
				t.Errorf("Divisao(%v, %v) = %v; esperado %v", 
					teste.a, teste.b, resultado, teste.esperado)
			}
		})
	}

	// Teste para divisão por zero
	t.Run("Divisão por Zero", func(t *testing.T) {
		_, err := Divisao(5, 0)
		if err == nil {
			t.Error("Divisao(5, 0) não retornou erro; esperado erro de divisão por zero")
		}
	})
}

// TestRaizQuadrada testa a função RaizQuadrada
func TestRaizQuadrada(t *testing.T) {
	// Casos de teste bem-sucedidos
	testes := []struct {
		nome     string
		n        float64
		esperado float64
	}{
		{"Zero", 0, 0},
		{"Um", 1, 1},
		{"Quatro", 4, 2},
		{"Nove", 9, 3},
		{"Número não-perfeito", 2, math.Sqrt(2)},
	}

	// Executar os casos de teste bem-sucedidos
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			resultado, err := RaizQuadrada(teste.n)
			if err != nil {
				t.Errorf("RaizQuadrada(%v) retornou erro inesperado: %v", 
					teste.n, err)
			}
			if resultado != teste.esperado {
				t.Errorf("RaizQuadrada(%v) = %v; esperado %v", 
					teste.n, resultado, teste.esperado)
			}
		})
	}

	// Teste para raiz quadrada de número negativo
	t.Run("Número Negativo", func(t *testing.T) {
		_, err := RaizQuadrada(-1)
		if err == nil {
			t.Error("RaizQuadrada(-1) não retornou erro; esperado erro de número negativo")
		}
	})
}

// TestPotencia testa a função Potencia
func TestPotencia(t *testing.T) {
	// Casos de teste
	testes := []struct {
		nome     string
		a, b     float64
		esperado float64
	}{
		{"Base Positiva, Expoente Positivo", 2, 3, 8},
		{"Base Positiva, Expoente Zero", 2, 0, 1},
		{"Base Positiva, Expoente Negativo", 2, -1, 0.5},
		{"Base Negativa, Expoente Par", -2, 2, 4},
		{"Base Negativa, Expoente Ímpar", -2, 3, -8},
		{"Base Zero, Expoente Positivo", 0, 3, 0},
	}

	// Executar os casos de teste
	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			resultado := Potencia(teste.a, teste.b)
			if resultado != teste.esperado {
				t.Errorf("Potencia(%v, %v) = %v; esperado %v", 
					teste.a, teste.b, resultado, teste.esperado)
			}
		})
	}
} 