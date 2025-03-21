# Coleção de Aplicações em Go

Esta é uma coleção de aplicações desenvolvidas em Go, cada uma demonstrando diferentes conceitos, arquiteturas e padrões de design.

## Aplicações Disponíveis

### App01 - Hello World
Introdução básica à linguagem Go, imprimindo "Hello World" na tela.
- **Conceitos**: Sintaxe básica de Go, pacotes, função main
- **Docker**: `docker-compose -f docker-compose-app01.yml up`

### App02 - Variáveis e Tipos
Demonstração dos tipos de dados e variáveis em Go.
- **Conceitos**: Tipos de dados, declaração de variáveis, conversões
- **Docker**: `docker-compose -f docker-compose-app02.yml up`

### App03 - Estruturas de Controle
Implementação de estruturas de controle em Go.
- **Conceitos**: If/else, loops, switch, defer
- **Docker**: `docker-compose -f docker-compose-app03.yml up`

### App04 - Funções
Aplicação demonstrando uso de funções em Go.
- **Conceitos**: Declaração de funções, retornos múltiplos, funções anônimas
- **Docker**: `docker-compose -f docker-compose-app04.yml up`

### App05 - Estruturas e Interfaces
Demonstração de estruturas e interfaces em Go.
- **Conceitos**: Structs, métodos, interfaces, embedding
- **Docker**: `docker-compose -f docker-compose-app05.yml up`

### App06 - Concorrência
Aplicação demonstrando concorrência em Go.
- **Conceitos**: Goroutines, channels, select, WaitGroups
- **Docker**: `docker-compose -f docker-compose-app06.yml up`

### App07 - APIs REST Básicas
Implementação de uma API REST simples.
- **Conceitos**: HTTP handlers, rotas, JSON
- **Docker**: `docker-compose -f docker-compose-app07.yml up`

### App08 - Templates HTML
Aplicação web com templates HTML.
- **Conceitos**: Templates, parsing, renderização
- **Docker**: `docker-compose -f docker-compose-app08.yml up`

### App09 - Banco de Dados
Aplicação com integração a banco de dados.
- **Conceitos**: SQL, CRUD, migrations
- **Docker**: `docker-compose -f docker-compose-app09.yml up`

### App10 - Padrões de Design
Implementação de padrões de design em Go.
- **Conceitos**: Singleton, Factory, Observer, etc.
- **Docker**: `docker-compose -f docker-compose-app10.yml up`

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
cd app01

# Executar a aplicação
go run main.go
```

### Usando Docker Compose

```bash
# Executar uma aplicação específica
docker-compose -f docker-compose-app01.yml up

# Parar a aplicação
docker-compose -f docker-compose-app01.yml down
```

### Usando o Script de Execução

Um script Bash foi criado para facilitar a execução das aplicações:

```bash
# Dar permissão de execução ao script (Linux/Mac)
chmod +x run-all-apps.sh

# Executar o script
./run-all-apps.sh
```

O script oferece um menu interativo para escolher qual aplicação executar, incluindo a opção de executar todas as aplicações simultaneamente.

## Requisitos

- Go 1.16+ (para execução local)
- Docker e Docker Compose (para execução via contêineres)

## Estrutura do Repositório

Cada aplicação está em seu próprio diretório e é independente das demais:

```
go-backend/
├── app01/  # Hello World
├── app02/  # Variáveis e Tipos
├── app03/  # Estruturas de Controle
├── app04/  # Funções
├── app05/  # Estruturas e Interfaces
├── app06/  # Concorrência
├── app07/  # APIs REST Básicas
├── app08/  # Templates HTML
├── app09/  # Banco de Dados
├── app10/  # Padrões de Design
├── app11/  # Middleware para Autenticação e Logging
├── app12/  # Upload de Arquivos
├── app13/  # Testes Automatizados
├── app14/  # API RESTful com Boas Práticas
└── app15/  # Blog API com Arquitetura Hexagonal
```

## Contribuição

Sinta-se à vontade para contribuir com este repositório, seja melhorando as aplicações existentes ou adicionando novas.
