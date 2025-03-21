#!/bin/bash

# Script para facilitar a execução das aplicações Go

# Cores para mensagens no terminal
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Script de execução das aplicações Go ===${NC}"
echo ""

# Verificar se o Docker está instalado
if ! command -v docker &> /dev/null; then
    echo -e "${YELLOW}Docker não encontrado. Por favor instale o Docker para continuar.${NC}"
    exit 1
fi

# Verificar se o Docker Compose está instalado
if ! command -v docker-compose &> /dev/null; then
    echo -e "${YELLOW}Docker Compose não encontrado. Por favor instale o Docker Compose para continuar.${NC}"
    exit 1
fi

# Menu de opções
while true; do
    echo "Escolha uma opção:"
    echo "1) Executar App11 - Middleware para Autenticação e Logging"
    echo "2) Executar App12 - Upload de Arquivos"
    echo "3) Executar App13 - Testes Automatizados"
    echo "4) Executar App14 - API RESTful com Boas Práticas"
    echo "5) Executar App15 - Blog API com Arquitetura Hexagonal"
    echo "6) Executar Todas as Aplicações"
    echo "7) Parar Todas as Aplicações"
    echo "0) Sair"
    
    read -p "Opção: " choice
    
    case $choice in
        1)
            echo -e "${GREEN}Iniciando App11...${NC}"
            docker-compose -f docker-compose-app11.yml up -d
            echo -e "${GREEN}App11 disponível em: http://localhost:8080${NC}"
            ;;
        2)
            echo -e "${GREEN}Iniciando App12...${NC}"
            docker-compose -f docker-compose-app12.yml up -d
            echo -e "${GREEN}App12 disponível em: http://localhost:8080${NC}"
            ;;
        3)
            echo -e "${GREEN}Iniciando App13...${NC}"
            docker-compose -f docker-compose-app13.yml up -d
            echo -e "${GREEN}App13 disponível em: http://localhost:8080${NC}"
            ;;
        4)
            echo -e "${GREEN}Iniciando App14...${NC}"
            docker-compose -f docker-compose-app14.yml up -d
            echo -e "${GREEN}App14 disponível em: http://localhost:8080${NC}"
            ;;
        5)
            echo -e "${GREEN}Iniciando App15...${NC}"
            docker-compose -f docker-compose-app15.yml up -d
            echo -e "${GREEN}App15 disponível em: http://localhost:8080${NC}"
            ;;
        6)
            echo -e "${GREEN}Iniciando todas as aplicações...${NC}"
            docker-compose up -d
            echo -e "${GREEN}Apps disponíveis em:${NC}"
            echo -e "${GREEN}- App11: http://localhost:8011${NC}"
            echo -e "${GREEN}- App12: http://localhost:8012${NC}"
            echo -e "${GREEN}- App13: http://localhost:8013${NC}"
            echo -e "${GREEN}- App14: http://localhost:8014${NC}"
            echo -e "${GREEN}- App15: http://localhost:8015${NC}"
            ;;
        7)
            echo -e "${YELLOW}Parando todas as aplicações...${NC}"
            docker-compose down
            docker-compose -f docker-compose-app11.yml down
            docker-compose -f docker-compose-app12.yml down
            docker-compose -f docker-compose-app13.yml down
            docker-compose -f docker-compose-app14.yml down
            docker-compose -f docker-compose-app15.yml down
            echo -e "${GREEN}Todas as aplicações foram paradas.${NC}"
            ;;
        0)
            echo -e "${GREEN}Saindo...${NC}"
            exit 0
            ;;
        *)
            echo -e "${YELLOW}Opção inválida!${NC}"
            ;;
    esac
    
    echo ""
done 