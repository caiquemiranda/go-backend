# App5 - CRUD Completo com Dados em Memória

## Descrição
Este projeto implementa um gerenciador de tarefas com operações CRUD (Create, Read, Update, Delete) completas, utilizando armazenamento em memória com proteção para concorrência.

## Como executar
```bash
go run main.go
```

Depois teste as diferentes rotas:

### Operações básicas:
- GET http://localhost:8080/tarefas - Lista todas as tarefas
- GET http://localhost:8080/tarefas/1 - Obtém a tarefa com ID 1
- POST http://localhost:8080/tarefas - Cria uma nova tarefa
- PUT http://localhost:8080/tarefas/1 - Atualiza a tarefa com ID 1
- DELETE http://localhost:8080/tarefas/1 - Remove a tarefa com ID 1

### Operações especiais:
- GET http://localhost:8080/tarefas?prioridade=3 - Lista tarefas com prioridade alta (3)
- GET http://localhost:8080/tarefas?concluida=true - Lista tarefas concluídas
- GET http://localhost:8080/tarefas?concluida=false - Lista tarefas pendentes
- PATCH http://localhost:8080/tarefas/1?concluida=true - Marca a tarefa como concluída

### JSON para criar ou atualizar tarefas:
```json
{
  "titulo": "Estudar Go",
  "descricao": "Aprender sobre CRUD e APIs REST",
  "prioridade": 2
}
```

## Conceitos abordados
- Implementação de CRUD completo
- Uso de diferentes métodos HTTP (GET, POST, PUT, PATCH, DELETE)
- Proteção de concorrência com mutex
- Uso de query parameters para filtragem de dados
- Encapsulamento da lógica de negócios
- Manipulação de tempo e datas
- Status HTTP apropriados para diferentes operações
- Uso avançado de estruturas de dados
- Uso de ponteiros para gerenciamento eficiente de recursos 