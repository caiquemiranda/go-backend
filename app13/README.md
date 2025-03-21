# App13 - Testes Automatizados em Go

## Descrição
Este projeto demonstra a implementação de testes automatizados em Go, abordando dois tipos principais de testes:

1. **Testes Unitários**: Testam componentes individuais isoladamente
2. **Testes de Integração**: Testam como os componentes interagem entre si

O projeto simula uma calculadora com API web, onde implementamos testes tanto para as funções matemáticas quanto para os endpoints HTTP.

## Como Executar a Aplicação
1. Certifique-se de ter Go instalado (versão 1.16+)
2. Navegue até o diretório do projeto
3. Execute o comando:
   ```
   go run main.go
   ```
4. Acesse a API através de um navegador ou ferramenta como cURL/Postman:
   - Exemplo: http://localhost:8080/soma?a=5&b=3

## Como Executar os Testes
Para executar todos os testes do projeto:
```
go test ./... -v
```

Para executar testes específicos:
```
go test ./calculadora -v  # Apenas testes unitários da calculadora
go test ./servidor -v     # Apenas testes de integração do servidor
```

Para executar um teste específico (usando regex):
```
go test ./calculadora -run TestSoma -v
```

Para verificar a cobertura de testes:
```
go test ./... -cover
```

## Estrutura do Projeto
- `calculadora/`: Pacote com funções matemáticas básicas
  - `calculadora.go`: Implementação das funções
  - `calculadora_test.go`: Testes unitários
- `servidor/`: Pacote com API HTTP para a calculadora
  - `servidor.go`: Implementação do servidor e handlers
  - `servidor_test.go`: Testes de integração
- `main.go`: Inicialização do servidor

## Conceitos Abordados
- **Estruturação e Organização de Testes**: Separação entre testes unitários e de integração
- **Table-Driven Tests**: Uso de tabelas para testar múltiplos casos
- **Testes Unitários**: Validação isolada de componentes individuais
- **Testes de Integração**: Validação da interação entre componentes
- **Test Runners**: Uso do pacote `testing` do Go
- **HTTP Testing**: Testes de API com `httptest`
- **Asserções em Testes**: Verificação de resultados esperados
- **Mocking**: Simulação de componentes em testes
- **Tratamento de Erros em Testes**: Verificação de fluxos de erro

## Boas Práticas Demonstradas
- Nomenclatura clara para funções de teste (`TestNomeDaFuncao`)
- Uso de subtestes para organizar casos de teste
- Mensagens de erro descritivas
- Isolamento adequado entre testes
- Verificação tanto de fluxos de sucesso quanto de erro
- Limpeza de recursos após testes (usando `defer`) 