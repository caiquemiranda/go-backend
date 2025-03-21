# App3 - Envio e Recebimento de JSON

## Descrição
Este projeto demonstra como trabalhar com JSON em uma API REST usando Go, incluindo serialização (codificação) e desserialização (decodificação) de dados.

## Como executar
```bash
go run main.go
```

Depois teste as diferentes rotas:
- GET http://localhost:8080/usuario - Retorna um usuário em formato JSON
- POST http://localhost:8080/criar-usuario - Cria um novo usuário a partir de dados JSON

Para testar a rota POST, você pode usar:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"nome":"Maria Santos","email":"maria@exemplo.com","idade":28}' http://localhost:8080/criar-usuario
```

Ou usar ferramentas como Postman enviando um JSON com esta estrutura:
```json
{
  "nome": "Maria Santos",
  "email": "maria@exemplo.com",
  "idade": 28
}
```

## Conceitos abordados
- Serialização e desserialização de JSON
- Uso de estruturas (structs) com tags JSON
- Leitura do corpo da requisição
- Configuração de cabeçalhos HTTP
- Resposta com diferentes status HTTP
- Definição de tipos personalizados em Go 