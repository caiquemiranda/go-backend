# App11 - Middleware para Autenticação e Logging

## Descrição
Este projeto demonstra a implementação e uso de middlewares em uma aplicação web Go, focando em duas funcionalidades essenciais:
1. **Middleware de Logging**: Registra informações sobre cada requisição HTTP, incluindo método, URI, endereço remoto e tempo de processamento.
2. **Middleware de Autenticação**: Implementa autenticação básica HTTP para proteger rotas privadas.

## Como Executar
1. Certifique-se de ter Go instalado (versão 1.16+)
2. Navegue até o diretório do projeto
3. Execute o comando:
   ```
   go run main.go
   ```
4. Acesse as rotas através de um navegador ou ferramenta como cURL/Postman:
   - Rota pública: http://localhost:8080/public
   - Rota privada: http://localhost:8080/private (requer autenticação)

### Credenciais para Teste
- Username: admin, Password: senha123
- Username: user, Password: 123456

## Conceitos Abordados
- **Middleware**: Funções que processam requisições HTTP antes de chegarem ao handler final
- **Encadeamento de Middlewares**: Como combinar múltiplos middlewares em sequência
- **Autenticação Básica HTTP**: Implementação e verificação de credenciais
- **Logging**: Rastreamento e registro de requisições com métricas de performance
- **http.Handler Interface**: Uso do padrão de design do Go para handlers HTTP
- **Proteção de Rotas**: Controle de acesso a recursos sensíveis

## Exemplos de Uso com cURL
```bash
# Acessar rota pública
curl http://localhost:8080/public

# Acessar rota privada sem autenticação (deve falhar)
curl http://localhost:8080/private

# Acessar rota privada com autenticação
curl -u admin:senha123 http://localhost:8080/private
``` 