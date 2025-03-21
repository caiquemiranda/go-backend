# App10 - Autenticação e Middlewares

## Descrição
Este projeto implementa uma API REST com sistema completo de autenticação JWT e controle de acesso baseado em perfis de usuário. A aplicação demonstra diversos conceitos avançados de middlewares, autenticação e autorização em Go.

## Características
- Sistema completo de autenticação com JSON Web Tokens (JWT)
- Controle de acesso baseado em perfis (RBAC - Role-Based Access Control)
- Middlewares para autenticação, autorização, logging e CORS
- Tokens de acesso e refresh tokens para renovação de sessão
- Diferentes níveis de acesso a recursos (público, usuário, editor, admin)

## Pré-requisitos
- Go 1.16 ou superior
- Dependências:
  - github.com/golang-jwt/jwt/v5
  - golang.org/x/crypto

## Como executar
1. Navegue até o diretório do projeto:
```
cd app10
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

## Estrutura do projeto
```
app10/
├── models/             # Modelos de dados
│   ├── usuario.go      # Modelo de usuário
│   ├── recurso.go      # Modelo de recurso protegido
│   └── repositories.go # Implementações de repositórios
├── handlers/           # Manipuladores HTTP
│   ├── auth_handler.go    # Manipulador de autenticação
│   └── recurso_handler.go # Manipulador de recursos
├── middlewares/        # Middlewares HTTP
│   └── auth_middleware.go # Middlewares de autenticação e autorização
├── utils/              # Utilitários
│   └── auth.go         # Funções de autenticação com JWT
├── main.go             # Arquivo principal
└── go.mod              # Definição de dependências
```

## Usuários de teste
A aplicação é inicializada com três usuários de teste:
- Admin: admin@exemplo.com / senha123
- Editor: editor@exemplo.com / senha123
- Usuário comum: usuario@exemplo.com / senha123

## Rotas da API

### Autenticação
- **POST /auth/registro** - Registro de novos usuários
  ```json
  {
    "nome": "Novo Usuário",
    "email": "novo@exemplo.com",
    "senha": "senha123"
  }
  ```

- **POST /auth/login** - Autenticação de usuários
  ```json
  {
    "email": "admin@exemplo.com",
    "senha": "senha123"
  }
  ```

- **POST /auth/refresh** - Renovação de token
  ```json
  {
    "refreshToken": "seu-refresh-token"
  }
  ```

### Recursos
- **GET /recursos-publicos** - Lista recursos públicos (não requer autenticação)
- **GET /recursos** - Lista recursos acessíveis ao usuário autenticado
- **POST /recursos** - Cria um novo recurso (requer autenticação)
  ```json
  {
    "titulo": "Novo Recurso",
    "descricao": "Descrição do recurso",
    "conteudo": "Conteúdo detalhado...",
    "categoria": "Tecnologia",
    "acessoLevel": 1
  }
  ```
- **GET /recursos/{id}** - Obtém um recurso específico (verificação de permissão)
- **PUT /recursos/{id}** - Atualiza completamente um recurso
- **PATCH /recursos/{id}** - Atualiza parcialmente um recurso
- **DELETE /recursos/{id}** - Remove um recurso

### Áreas restritas
- **GET /admin** - Área administrativa (requer perfil "admin")
- **GET /editor** - Área de editores (requer perfil "editor" ou "admin")

## Conceitos abordados

### Autenticação
- **JWT (JSON Web Tokens)**: Implementação completa de autenticação baseada em tokens
- **Refresh Tokens**: Mecanismo para renovação de sessão sem reautenticação
- **Hashing de senhas**: Armazenamento seguro de senhas com bcrypt
- **Authorization Header**: Padrão de envio de tokens via cabeçalho "Bearer"

### Middlewares
- **Cadeia de middlewares**: Composição e encadeamento de middlewares
- **Middleware de Autenticação**: Verificação de tokens JWT
- **Middleware de Autorização**: Verificação de papéis/permissões
- **Middleware de Logger**: Registro de informações de requisições
- **Middleware CORS**: Configuração para acesso cross-origin

### Controle de Acesso
- **RBAC (Role-Based Access Control)**: Controle baseado em papéis
- **Níveis de acesso hierárquicos**: Usuário > Editor > Admin
- **Controle granular**: Recursos com diferentes níveis de acesso
- **Filtragem de conteúdo**: Exibição de dados conforme nível de acesso

### Resposta de API
- **Códigos HTTP adequados**: Uso correto dos status codes HTTP
- **Mensagens de erro padronizadas**: Formato consistente para erros
- **Sanitização de dados**: Filtragem de informações sensíveis em respostas 