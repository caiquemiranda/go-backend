# App6 - Validação de Dados da Requisição

## Descrição
Este projeto implementa um sistema de cadastro e autenticação de usuários com validação avançada de dados, demonstrando as melhores práticas para validar entradas em APIs REST.

## Como executar
```bash
go run main.go
```

Depois teste as rotas:

### Cadastro de usuário
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "nome": "João Silva",
  "email": "joao@exemplo.com",
  "senha": "Senha@123",
  "telefone": "(11)98765-4321",
  "cpf": "123.456.789-00",
  "idade": 30
}' http://localhost:8080/cadastro
```

### Login
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "email": "joao@exemplo.com",
  "senha": "Senha@123"
}' http://localhost:8080/login
```

### Listar usuários
```bash
curl -X GET http://localhost:8080/usuarios
```

## Conceitos abordados
- Validação detalhada de dados de entrada
- Expressões regulares para validação de padrões
- Formatação de respostas de erro padronizadas
- Validação de tipos específicos de dados (email, CPF, telefone)
- Validação de regras de segurança para senhas
- Modelagem de respostas HTTP com códigos de status adequados
- Uso de ponteiros para otimizar o retorno de erros de validação
- Separação clara entre validação e manipulação de requisições
- Ocultação de dados sensíveis (senha) nas respostas 