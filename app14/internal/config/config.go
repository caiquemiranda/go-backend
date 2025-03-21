package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config contém a configuração do servidor
type Config struct {
	ServerPort int
	APIVersion string
}

// Valores padrão
const (
	defaultServerPort = 8080
	defaultAPIVersion = "v1"
)

// LoadConfig carrega a configuração do ambiente ou usa valores padrão
func LoadConfig() *Config {
	return &Config{
		ServerPort: getEnvAsInt("SERVER_PORT", defaultServerPort),
		APIVersion: getEnv("API_VERSION", defaultAPIVersion),
	}
}

// String retorna uma representação string da configuração
func (c *Config) String() string {
	return fmt.Sprintf("ServerPort: %d, APIVersion: %s", c.ServerPort, c.APIVersion)
}

// getEnv retorna o valor da variável de ambiente ou o valor padrão
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt retorna o valor inteiro da variável de ambiente ou o valor padrão
func getEnvAsInt(key string, defaultValue int) int {
	if strValue := getEnv(key, ""); strValue != "" {
		if value, err := strconv.Atoi(strValue); err == nil {
			return value
		}
	}
	return defaultValue
} 