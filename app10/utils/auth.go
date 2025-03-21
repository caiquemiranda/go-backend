package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Configuração do JWT
const (
	// Na aplicação real, essas chaves deveriam ser armazenadas de forma segura
	// e não diretamente no código-fonte
	jwtSecret           = "meu_segredo_super_secreto_123"
	jwtExpirationHours  = 24
	jwtRefreshSecret    = "meu_segredo_super_secreto_para_refresh_456"
	jwtRefreshExpMonths = 1
)

// UserClaims define as claims personalizadas para o token JWT
type UserClaims struct {
	UserID uint   `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GerarToken cria um novo token JWT para o usuário
func GerarToken(userID uint, email, role string) (string, error) {
	claims := UserClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * jwtExpirationHours)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "app10-api",
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GerarRefreshToken cria um token de atualização com validade mais longa
func GerarRefreshToken(userID uint, email string) (string, error) {
	claims := UserClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, jwtRefreshExpMonths, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "app10-api",
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtRefreshSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidarToken verifica se o token JWT é válido e retorna suas claims
func ValidarToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura é o esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}

// ValidarRefreshToken valida um token de atualização
func ValidarRefreshToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}

		return []byte(jwtRefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("refresh token inválido")
	}

	return claims, nil
}

// ExtrairTokenDaRequisicao obtém o token do cabeçalho Authorization
func ExtrairTokenDaRequisicao(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("cabeçalho Authorization ausente")
	}

	// O formato esperado é "Bearer {token}"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("formato do cabeçalho Authorization inválido")
	}

	return parts[1], nil
}

// NivelAcessoParaRole converte um papel de usuário em nível de acesso
func NivelAcessoParaRole(role string) int {
	switch role {
	case "admin":
		return 3
	case "editor":
		return 2
	case "usuario":
		return 1
	default:
		return 0
	}
}

// VerificarRole verifica se o usuário tem o papel mínimo necessário
func VerificarRole(userRole string, requiredRole string) bool {
	userLevel := NivelAcessoParaRole(userRole)
	requiredLevel := NivelAcessoParaRole(requiredRole)
	
	return userLevel >= requiredLevel
} 