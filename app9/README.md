# App9 - CRUD com GORM (ORM)

## Descrição
Este projeto implementa um sistema de gerenciamento de produtos utilizando GORM, um ORM (Object-Relational Mapping) popular para Go. O aplicativo oferece uma API RESTful completa para gerenciar produtos, com recursos avançados de filtragem, ordenação e pesquisa.

## Pré-requisitos
- Go 1.16 ou superior
- Dependências:
  - gorm.io/gorm
  - gorm.io/driver/sqlite

## Como executar
1. Navegue até o diretório do projeto:
```
cd app9
```

2. Baixe as dependências:
```
go mod tidy
```

3. Execute a aplicação:
```
go run main.go
```

4. O servidor será iniciado na porta 8080

## Rotas e Endpoints

### Produtos
- **GET /produtos** - Lista todos os produtos (aceita filtros)
  - Parâmetros de consulta:
    - categoria: Filtra por categoria
    - disponivel: Filtra por disponibilidade (true/false)
    - precoMin: Filtra por preço mínimo
    - precoMax: Filtra por preço máximo
    - ordenarPor: Campo para ordenação
    - direcao: Direção da ordenação (asc/desc)

- **POST /produtos** - Cria um novo produto
- **GET /produtos/{id}** - Obtém um produto específico pelo ID
- **PUT /produtos/{id}** - Atualiza um produto existente
- **DELETE /produtos/{id}** - Remove um produto
- **PATCH /produtos/{id}?estoque=10** - Atualiza apenas o estoque do produto
- **PATCH /produtos/{id}?disponivel=true** - Atualiza apenas a disponibilidade do produto
- **GET /produtos/sku/{sku}** - Busca um produto pelo SKU

## Exemplo de JSON para Produtos

### Criar/Atualizar Produto
```json
{
  "nome": "Notebook Dell XPS 13",
  "descricao": "Notebook ultrafino com processador Intel Core i7, 16GB RAM, 512GB SSD",
  "preco": 6500.00,
  "estoque": 5,
  "categoria": "Notebooks",
  "sku": "NTBK-DELL-002",
  "disponivel": true
}
```

## Estrutura do Banco de Dados
O aplicativo utiliza SQLite como banco de dados, armazenado no arquivo `produtos.db` que é criado automaticamente na primeira execução. A estrutura da tabela é definida através das tags GORM no modelo `Produto`.

## Conceitos Abordados

### ORM (Object-Relational Mapping)
- **Mapeamento objeto-relacional**: Transformação entre objetos Go e tabelas do banco de dados
- **Tags GORM**: Configuração de campos, restrições e índices diretamente no modelo
- **Migrações automáticas**: Criação e atualização automática do esquema do banco de dados

### Consultas Avançadas com GORM
- **Where**: Filtragem com condições
- **Model e Updates**: Atualização parcial de registros
- **Order**: Ordenação de resultados
- **Transactions**: Operações atômicas garantindo consistência dos dados

### Design de API
- **Filtragem**: Implementação de filtros via query strings
- **Ordenação**: Controle da ordem dos resultados
- **Pesquisa por campos específicos**: Busca por ID ou SKU
- **PATCH para atualizações parciais**: Modificação de apenas um campo

### Padrões de Projeto
- **Repository Pattern**: Encapsulamento das operações de banco de dados
- **Modelos com tags e validações**: Garantia de integridade dos dados 