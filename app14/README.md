# App14 - API RESTful com Boas Práticas

## Descrição
Este projeto demonstra a implementação de uma API RESTful completa em Go usando boas práticas de desenvolvimento, arquitetura limpa e organização modular. A API gerencia tarefas (tasks) e implementa todos os endpoints CRUD com tratamento adequado de erros e validações.

## Como Executar
1. Certifique-se de ter Go instalado (versão 1.16+)
2. Navegue até o diretório do projeto
3. Execute o comando:
   ```
   go run cmd/server/main.go
   ```
4. O servidor será iniciado na porta 8080 (ou a porta definida na variável de ambiente SERVER_PORT)

## Endpoints da API
- `GET /api/tasks` - Lista todas as tarefas
- `GET /api/tasks/{id}` - Obtém uma tarefa específica
- `POST /api/tasks` - Cria uma nova tarefa
- `PUT /api/tasks/{id}` - Atualiza uma tarefa existente
- `DELETE /api/tasks/{id}` - Remove uma tarefa

### Exemplo de Payload para Criar/Atualizar Tarefa
```json
{
  "title": "Minha Tarefa",
  "description": "Descrição da minha tarefa",
  "status": "pending"
}
```

### Status Possíveis
- `pending` - Pendente
- `in_progress` - Em Progresso
- `completed` - Concluída
- `cancelled` - Cancelada

## Estrutura do Projeto
```
app14/
│
├── cmd/
│   └── server/
│       └── main.go             # Ponto de entrada da aplicação
│
├── internal/                   # Código interno da aplicação
│   ├── api/
│   │   └── router.go           # Configuração de rotas
│   │
│   ├── config/
│   │   └── config.go           # Configurações da aplicação
│   │
│   ├── database/
│   │   └── task_repo.go        # Implementação do repositório
│   │
│   ├── handlers/
│   │   └── task.go             # Handlers HTTP
│   │
│   ├── middleware/
│   │   └── logger.go           # Middleware de logging
│   │
│   └── models/
│       └── task.go             # Definição de modelos
│
└── go.mod                      # Dependências do módulo
```

## Conceitos Abordados
- **Arquitetura Limpa**: Separação clara de responsabilidades
- **Design RESTful**: Nomenclatura adequada de endpoints e métodos HTTP
- **Validação de Dados**: Verificação da integridade dos dados de entrada
- **Tratamento de Erros**: Respostas HTTP com códigos e mensagens apropriadas
- **Middleware**: Camada para processamento transversal de requisições
- **Graceful Shutdown**: Encerramento adequado do servidor com tempo para finalizar conexões
- **Repositórios**: Padrão para abstração de acesso a dados
- **Configuração Externalizada**: Uso de variáveis de ambiente para configuração
- **Injeção de Dependências**: Passagem explícita de dependências
- **Testes**: Estrutura preparada para testes unitários e de integração

## Observações
- Este projeto utiliza armazenamento em memória para simplificar a demonstração. Em um ambiente de produção, seria utilizado um banco de dados persistente.
- Não há autenticação implementada, o que seria essencial em uma API real. 