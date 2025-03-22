# App7 - Separação por Camadas (MVC)

## Descrição
Este projeto implementa um sistema de gerenciamento de biblioteca com separação clara entre as camadas de dados (models), lógica de negócio (services) e interface de usuário (handlers), seguindo o padrão MVC (Model-View-Controller) adaptado para APIs REST.

## Estrutura do Projeto
```
app7/
├── models/      # Definição das estruturas de dados
├── services/    # Regras de negócio e operações de dados
├── handlers/    # Manipulação de requisições HTTP
└── main.go      # Ponto de entrada da aplicação
```

## Como executar
```bash
cd app7
go run main.go
```

Depois teste as diferentes rotas:

### Operações com livros:
- GET http://localhost:8080/livros - Lista todos os livros
- GET http://localhost:8080/livros/1 - Obtém o livro com ID 1
- POST http://localhost:8080/livros - Cria um novo livro
- PUT http://localhost:8080/livros/1 - Atualiza o livro com ID 1
- DELETE http://localhost:8080/livros/1 - Remove o livro com ID 1
- PATCH http://localhost:8080/livros/1?disponivel=false - Marca o livro como indisponível

### JSON para criar ou atualizar livros:
```json
{
  "titulo": "Clean Code",
  "autor": "Robert C. Martin",
  "editora": "Prentice Hall",
  "anoPublicacao": 2008,
  "isbn": "9780132350884"
}
```

## Conceitos abordados
- Arquitetura em camadas (MVC para APIs)
- Separação de responsabilidades
- Injeção de dependências
- Encapsulamento de lógica de negócio
- Tratamento de erros em diferentes camadas
- Modelagem de domínio
- Organização de código em pacotes
- Uso de tipos de retorno compostos (tuplas)
- Controle de concorrência com mutex 