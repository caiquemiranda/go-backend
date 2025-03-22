# App8 - CRUD com SQLite

## Descrição
Este projeto implementa um sistema de gerenciamento de notas com persistência em banco de dados SQLite. Ele demonstra como utilizar o pacote `database/sql` para realizar operações CRUD em um banco de dados relacional.

## Pré-requisitos
- Go instalado
- Driver SQLite para Go
  ```bash
  go get github.com/mattn/go-sqlite3
  ```

## Como executar
```bash
cd app8
go mod init app8
go get github.com/mattn/go-sqlite3
go run main.go
```

O servidor será iniciado na porta 8080 e criará um arquivo `notas.db` no diretório atual para armazenar os dados.

## Endpoints disponíveis

### Operações com notas
- **GET /notas** - Lista todas as notas
- **GET /notas?categoria=trabalho** - Lista notas filtradas por categoria
- **GET /notas/1** - Obtém a nota com ID 1
- **POST /notas** - Cria uma nova nota
- **PUT /notas/1** - Atualiza a nota com ID 1
- **DELETE /notas/1** - Remove a nota com ID 1
- **PATCH /notas/1?arquivada=true** - Arquiva a nota com ID 1

### JSON para criar ou atualizar notas
```json
{
  "titulo": "Reunião de equipe",
  "conteudo": "Discutir os novos projetos para o próximo trimestre",
  "categoria": "trabalho"
}
```

## Conceitos abordados
- Conexão com banco de dados SQLite
- Uso do pacote `database/sql`
- Preparação e execução de consultas SQL
- Mapeamento de resultados para estruturas Go
- Tratamento de transações e erros de banco de dados
- Repository pattern para operações de banco de dados
- Migrations simples (criação de tabela)
- Uso de parâmetros na consulta SQL para evitar injeção
- Leitura e escrita de dados estruturados no banco 