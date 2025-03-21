# App12 - Upload de Arquivos com Go

## Descrição
Este projeto demonstra como implementar um sistema completo de gerenciamento de uploads de arquivos em Go, incluindo:
- Interface web amigável para upload
- Visualização de arquivos enviados
- Download de arquivos
- Exclusão de arquivos
- Limitação de tamanho de arquivos (10MB)
- Tratamento de erros

## Como Executar
1. Certifique-se de ter Go instalado (versão 1.16+)
2. Navegue até o diretório do projeto
3. Execute o comando:
   ```
   go run main.go
   ```
4. Acesse o aplicativo em seu navegador: http://localhost:8080

## Conceitos Abordados
- **Manipulação de Formulários Multipart**: Processamento de formulários com envio de arquivos
- **Manipulação de Arquivos em Go**: Operações de leitura, escrita e exclusão de arquivos
- **Cabeçalhos HTTP**: Configuração correta para download e visualização de arquivos
- **Templates HTML**: Renderização de páginas dinâmicas
- **Tratamento de Erros**: Validação de entradas e tratamento adequado de erros
- **Segurança Web Básica**: Validação de caminhos e prevenção de acesso a arquivos não autorizados
- **Redirecionamentos HTTP**: Implementação de fluxo de usuário adequado após operações
- **Persistência de Dados em Disco**: Armazenamento de arquivos em sistema de arquivos

## Funcionalidades
1. **Upload de Arquivos**: Interface para selecionar e enviar arquivos
2. **Limitação de Tamanho**: Restrição de 10MB por arquivo
3. **Listagem de Arquivos**: Exibição dos arquivos enviados com detalhes (nome, tamanho, data)
4. **Visualização**: Capacidade de visualizar arquivos diretamente no navegador
5. **Download**: Opção para baixar os arquivos
6. **Exclusão**: Remoção de arquivos com confirmação

## Estrutura do Projeto
- `main.go`: Código principal da aplicação
- `uploads/`: Diretório onde os arquivos enviados são armazenados

## Observações
- Os arquivos são armazenados localmente na pasta `uploads/`
- Para uma aplicação de produção, considere implementar:
  - Validação mais robusta de tipos de arquivo
  - Autenticação de usuários
  - Rastreamento de quem fez upload
  - Armazenamento em serviços de nuvem para maior escala 