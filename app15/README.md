# App15 - Blog API com Arquitetura Hexagonal

## Descrição
Este projeto implementa uma API completa para um sistema de blog, utilizando arquitetura hexagonal (também conhecida como arquitetura de portas e adaptadores) para garantir uma separação clara de responsabilidades e facilitar a manutenção e testabilidade do código.

## Características
- Sistema completo de autenticação com JWT
- Operações CRUD para posts e comentários
- Validação robusta de dados
- Tratamento de erros centralizado
- Logs estruturados
- Documentação de API com Swagger
- Arquitetura limpa e testável

## Arquitetura do Projeto

O projeto segue a arquitetura hexagonal (também conhecida como ports and adapters):

```
app15/
│
├── cmd/                     # Pontos de entrada da aplicação
│   └── api/                 # API HTTP
│       └── main.go          # Ponto de entrada principal
│
├── internal/                # Código interno não reutilizável
│   ├── domain/              # Entidades e regras de negócio
│   │   ├── user.go          # Entidade de usuário
│   │   ├── post.go          # Entidade de post
│   │   └── comment.go       # Entidade de comentário
│   │
│   ├── application/         # Casos de uso/serviços da aplicação
│   │   ├── auth_service.go  # Serviço de autenticação
│   │   ├── post_service.go  # Serviço de posts
│   │   └── comment_service.go # Serviço de comentários
│   │
│   ├── ports/               # Portas (interfaces) para adaptadores
│   │   ├── repositories/    # Interfaces para repositórios
│   │   └── services/        # Interfaces para serviços
│   │
│   └── adapters/            # Adaptadores para infraestrutura externa
│       ├── http/            # Adaptadores para API HTTP
│       │   ├── handlers/    # Handlers HTTP
│       │   ├── middleware/  # Middleware HTTP
│       │   └── router.go    # Configuração de rotas
│       └── db/              # Adaptadores para banco de dados
│           └── memory/      # Implementação em memória
│
├── pkg/                     # Código potencialmente reutilizável
│   ├── logger/              # Pacote de logging
│   ├── validator/           # Pacote de validação
│   └── jwt/                 # Pacote para JWT
│
├── docs/                    # Documentação
│   └── swagger/             # Documentação da API
│
└── tests/                   # Testes de integração e ponta a ponta
```

## Como Executar
1. Certifique-se de ter Go instalado (versão 1.20+)
2. Clone este repositório
3. Configure as variáveis de ambiente (ou crie um arquivo `.env` baseado no `.env.example`)
4. Execute o comando:
   ```
   go run cmd/api/main.go
   ```
5. A API estará disponível em http://localhost:8080

## Rotas da API

### Autenticação
- `POST /api/auth/register` - Registro de novo usuário
- `POST /api/auth/login` - Login de usuário

### Posts
- `GET /api/posts` - Listar todos os posts
- `GET /api/posts/{id}` - Obter detalhes de um post
- `POST /api/posts` - Criar um novo post
- `PUT /api/posts/{id}` - Atualizar um post existente
- `DELETE /api/posts/{id}` - Remover um post

### Comentários
- `GET /api/posts/{postId}/comments` - Listar comentários de um post
- `POST /api/posts/{postId}/comments` - Adicionar comentário a um post
- `PUT /api/comments/{id}` - Atualizar um comentário
- `DELETE /api/comments/{id}` - Remover um comentário

## Conceitos Abordados
- **Arquitetura Hexagonal**: Separação clara entre domínio, aplicação e infraestrutura
- **SOLID**: Aplicação dos princípios SOLID
- **Injeção de Dependência**: Desacoplamento através de interfaces
- **Desenvolvimento Orientado a Testes**: Projeto estruturado para facilitar testes
- **Clean Code**: Práticas de código limpo e manutenível
- **RESTful API**: Design de API seguindo princípios REST
- **Autenticação e Autorização**: Implementação de sistema de autenticação com JWT
- **Middleware**: Uso de middleware para aspectos transversais como logging e autenticação

## Observações
- Este projeto foi criado para fins didáticos e demonstra boas práticas de desenvolvimento
- Em um ambiente de produção, seria necessário adicionar mais recursos como:
  - Persistência em banco de dados relacional ou NoSQL
  - Cache para melhorar desempenho
  - Métricas e monitoramento
  - CI/CD para implantação automatizada 