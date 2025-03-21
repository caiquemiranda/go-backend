# Coleção de Aplicações em Go

Esta é uma coleção de aplicações desenvolvidas em Go, cada uma demonstrando diferentes conceitos, arquiteturas e padrões de design.

## Aplicações Disponíveis

### App11 - Middleware para Autenticação e Logging
Demonstra a implementação e uso de middlewares em Go para autenticação básica HTTP e logging de requisições.
- **Conceitos**: Middlewares, Autenticação Básica, Logging, Http Handlers
- **Docker**: `docker-compose -f docker-compose-app11.yml up`

### App12 - Upload de Arquivos
Sistema completo de gerenciamento de uploads de arquivos com interface web.
- **Conceitos**: Manipulação de arquivos, formulários multipart, templates HTML
- **Docker**: `docker-compose -f docker-compose-app12.yml up`

### App13 - Testes Automatizados
Calculadora com API web que demonstra testes unitários e de integração.
- **Conceitos**: Testes unitários, testes de integração, table-driven tests, mocking
- **Docker**: `docker-compose -f docker-compose-app13.yml up`

### App14 - API RESTful com Boas Práticas
API RESTful para gerenciamento de tarefas usando boas práticas de desenvolvimento.
- **Conceitos**: Arquitetura limpa, CRUD, middleware, graceful shutdown
- **Docker**: `docker-compose -f docker-compose-app14.yml up`

### App15 - Blog API com Arquitetura Hexagonal
API completa para um blog usando arquitetura hexagonal (ports and adapters).
- **Conceitos**: Arquitetura hexagonal, SOLID, JWT, domínio rico
- **Docker**: `docker-compose -f docker-compose-app15.yml up`

## Como Executar as Aplicações

Cada aplicação pode ser executada individualmente usando Go diretamente ou via Docker Compose.

### Usando Go

```bash
# Navegar até o diretório da aplicação
cd app11

# Executar a aplicação
go run main.go
```

### Usando Docker Compose

```bash
# Executar uma aplicação específica
docker-compose -f docker-compose-app11.yml up

# Parar a aplicação
docker-compose -f docker-compose-app11.yml down
```

## Requisitos

- Go 1.16+ (para execução local)
- Docker e Docker Compose (para execução via contêineres)

## Estrutura do Repositório

Cada aplicação está em seu próprio diretório e é independente das demais:

```
go-backend/
├── app11/  # Middleware para Autenticação e Logging
├── app12/  # Upload de Arquivos
├── app13/  # Testes Automatizados
├── app14/  # API RESTful com Boas Práticas
└── app15/  # Blog API com Arquitetura Hexagonal
```

## Contribuição

Sinta-se à vontade para contribuir com este repositório, seja melhorando as aplicações existentes ou adicionando novas.
