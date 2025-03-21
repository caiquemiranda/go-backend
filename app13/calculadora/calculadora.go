package calculadora

import (
	"errors"
	"math"
)

// Soma retorna a soma de dois números
func Soma(a, b float64) float64 {
	return a + b
}

// Subtracao retorna a diferença entre dois números
func Subtracao(a, b float64) float64 {
	return a - b
}

// Multiplicacao retorna o produto de dois números
func Multiplicacao(a, b float64) float64 {
	return a * b
}

// Divisao retorna o quociente da divisão de a por b
// Retorna um erro se b for zero
func Divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não é permitida")
	}
	return a / b, nil
}

// RaizQuadrada retorna a raiz quadrada de n
// Retorna um erro se n for negativo
func RaizQuadrada(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("não é possível calcular raiz quadrada de número negativo")
	}
	return math.Sqrt(n), nil
}

// Potencia retorna a elevado a b
func Potencia(a, b float64) float64 {
	return math.Pow(a, b)
} 