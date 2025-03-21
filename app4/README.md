# App4 - Manipulação de Dados em Memória

## Descrição
Este projeto implementa uma API de gerenciamento de produtos que utiliza estruturas de dados em memória (maps e slices) para armazenar e manipular informações.

## Como executar
```bash
go run main.go
```

Depois teste as diferentes rotas:
- GET http://localhost:8080/produtos - Lista todos os produtos
- POST http://localhost:8080/produtos - Cria um novo produto
- GET http://localhost:8080/produtos/1 - Obtém o produto com ID 1
- PUT http://localhost:8080/produtos/1 - Atualiza o produto com ID 1
- DELETE http://localhost:8080/produtos/1 - Remove o produto com ID 1
- GET http://localhost:8080/categorias/periféricos - Lista todos os produtos da categoria "periféricos"

Para testar as rotas POST e PUT, envie um JSON com esta estrutura:
```json
{
  "nome": "Monitor",
  "preco": 800.50,
  "estoque": 5,
  "categoria": "Periféricos"
}
```

## Conceitos abordados
- Uso de maps para armazenamento de dados em memória
- Uso de slices para operações com coleções de dados
- Manipulação de parâmetros na URL
- Filtragem e busca de dados
- Conversão entre tipos de dados (map para slice)
- Implementação de CRUD completo (Create, Read, Update, Delete)
- Uso de funções de manipulação de strings
- Geração de IDs sequenciais 