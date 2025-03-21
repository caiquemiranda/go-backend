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
    echo "1) Executar App01 - Hello World"
    echo "2) Executar App02 - Variáveis e Tipos"
    echo "3) Executar App03 - Estruturas de Controle"
    echo "4) Executar App04 - Funções"
    echo "5) Executar App05 - Estruturas e Interfaces"
    echo "6) Executar App06 - Concorrência"
    echo "7) Executar App07 - APIs REST Básicas"
    echo "8) Executar App08 - Templates HTML"
    echo "9) Executar App09 - Banco de Dados"
    echo "10) Executar App10 - Padrões de Design"
    echo "11) Executar App11 - Middleware para Autenticação e Logging"
    echo "12) Executar App12 - Upload de Arquivos"
    echo "13) Executar App13 - Testes Automatizados"
    echo "14) Executar App14 - API RESTful com Boas Práticas"
    echo "15) Executar App15 - Blog API com Arquitetura Hexagonal"
    echo "16) Executar Todas as Aplicações"
    echo "17) Parar Todas as Aplicações"
    echo "0) Sair"
    
    read -p "Opção: " choice
    
    case $choice in
        1)
            echo -e "${GREEN}Iniciando App01...${NC}"
            docker-compose -f docker-compose-app01.yml up -d
            echo -e "${GREEN}App01 disponível em: http://localhost:8080${NC}"
            ;;
        2)
            echo -e "${GREEN}Iniciando App02...${NC}"
            docker-compose -f docker-compose-app02.yml up -d
            echo -e "${GREEN}App02 disponível em: http://localhost:8080${NC}"
            ;;
        3)
            echo -e "${GREEN}Iniciando App03...${NC}"
            docker-compose -f docker-compose-app03.yml up -d
            echo -e "${GREEN}App03 disponível em: http://localhost:8080${NC}"
            ;;
        4)
            echo -e "${GREEN}Iniciando App04...${NC}"
            docker-compose -f docker-compose-app04.yml up -d
            echo -e "${GREEN}App04 disponível em: http://localhost:8080${NC}"
            ;;
        5)
            echo -e "${GREEN}Iniciando App05...${NC}"
            docker-compose -f docker-compose-app05.yml up -d
            echo -e "${GREEN}App05 disponível em: http://localhost:8080${NC}"
            ;;
        6)
            echo -e "${GREEN}Iniciando App06...${NC}"
            docker-compose -f docker-compose-app06.yml up -d
            echo -e "${GREEN}App06 disponível em: http://localhost:8080${NC}"
            ;;
        7)
            echo -e "${GREEN}Iniciando App07...${NC}"
            docker-compose -f docker-compose-app07.yml up -d
            echo -e "${GREEN}App07 disponível em: http://localhost:8080${NC}"
            ;;
        8)
            echo -e "${GREEN}Iniciando App08...${NC}"
            docker-compose -f docker-compose-app08.yml up -d
            echo -e "${GREEN}App08 disponível em: http://localhost:8080${NC}"
            ;;
        9)
            echo -e "${GREEN}Iniciando App09...${NC}"
            docker-compose -f docker-compose-app09.yml up -d
            echo -e "${GREEN}App09 disponível em: http://localhost:8080${NC}"
            ;;
        10)
            echo -e "${GREEN}Iniciando App10...${NC}"
            docker-compose -f docker-compose-app10.yml up -d
            echo -e "${GREEN}App10 disponível em: http://localhost:8080${NC}"
            ;;
        11)
            echo -e "${GREEN}Iniciando App11...${NC}"
            docker-compose -f docker-compose-app11.yml up -d
            echo -e "${GREEN}App11 disponível em: http://localhost:8080${NC}"
            ;;
        12)
            echo -e "${GREEN}Iniciando App12...${NC}"
            docker-compose -f docker-compose-app12.yml up -d
            echo -e "${GREEN}App12 disponível em: http://localhost:8080${NC}"
            ;;
        13)
            echo -e "${GREEN}Iniciando App13...${NC}"
            docker-compose -f docker-compose-app13.yml up -d
            echo -e "${GREEN}App13 disponível em: http://localhost:8080${NC}"
            ;;
        14)
            echo -e "${GREEN}Iniciando App14...${NC}"
            docker-compose -f docker-compose-app14.yml up -d
            echo -e "${GREEN}App14 disponível em: http://localhost:8080${NC}"
            ;;
        15)
            echo -e "${GREEN}Iniciando App15...${NC}"
            docker-compose -f docker-compose-app15.yml up -d
            echo -e "${GREEN}App15 disponível em: http://localhost:8080${NC}"
            ;;
        16)
            echo -e "${GREEN}Iniciando todas as aplicações...${NC}"
            docker-compose up -d
            echo -e "${GREEN}Apps disponíveis em:${NC}"
            echo -e "${GREEN}- App01: http://localhost:8001${NC}"
            echo -e "${GREEN}- App02: http://localhost:8002${NC}"
            echo -e "${GREEN}- App03: http://localhost:8003${NC}"
            echo -e "${GREEN}- App04: http://localhost:8004${NC}"
            echo -e "${GREEN}- App05: http://localhost:8005${NC}"
            echo -e "${GREEN}- App06: http://localhost:8006${NC}"
            echo -e "${GREEN}- App07: http://localhost:8007${NC}"
            echo -e "${GREEN}- App08: http://localhost:8008${NC}"
            echo -e "${GREEN}- App09: http://localhost:8009${NC}"
            echo -e "${GREEN}- App10: http://localhost:8010${NC}"
            echo -e "${GREEN}- App11: http://localhost:8011${NC}"
            echo -e "${GREEN}- App12: http://localhost:8012${NC}"
            echo -e "${GREEN}- App13: http://localhost:8013${NC}"
            echo -e "${GREEN}- App14: http://localhost:8014${NC}"
            echo -e "${GREEN}- App15: http://localhost:8015${NC}"
            ;;
        17)
            echo -e "${YELLOW}Parando todas as aplicações...${NC}"
            docker-compose down
            docker-compose -f docker-compose-app01.yml down
            docker-compose -f docker-compose-app02.yml down
            docker-compose -f docker-compose-app03.yml down
            docker-compose -f docker-compose-app04.yml down
            docker-compose -f docker-compose-app05.yml down
            docker-compose -f docker-compose-app06.yml down
            docker-compose -f docker-compose-app07.yml down
            docker-compose -f docker-compose-app08.yml down
            docker-compose -f docker-compose-app09.yml down
            docker-compose -f docker-compose-app10.yml down
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