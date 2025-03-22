# App2 - Servidor com Diferentes Métodos HTTP

## Descrição
Este projeto implementa um servidor HTTP que responde a diferentes métodos HTTP (GET e POST) em várias rotas.

## Como executar
```bash
go run main.go
```

Depois teste as diferentes rotas:
- GET http://localhost:8080/ - Página inicial
- GET http://localhost:8080/usuarios - Lista de usuários
- POST http://localhost:8080/usuarios - Criar usuário
- GET http://localhost:8080/produtos - Lista de produtos
- POST http://localhost:8080/produtos - Criar produto

Para testar os métodos POST, você pode usar ferramentas como:
- curl: `curl -X POST http://localhost:8080/usuarios`
- Postman
- Qualquer cliente HTTP que permita escolher o método

## Conceitos abordados
- Diferentes métodos HTTP (GET, POST)
- Roteamento condicional baseado no método HTTP
- Verificação e validação de métodos HTTP
- Retorno de códigos de erro HTTP apropriados
- Uso de switch para controle de fluxo baseado no método 