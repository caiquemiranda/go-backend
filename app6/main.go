package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Usuário representa os dados de um usuário no sistema
type Usuario struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Senha    string `json:"senha,omitempty"`
	Telefone string `json:"telefone"`
	CPF      string `json:"cpf"`
	Idade    int    `json:"idade"`
}

// ErroValidacao representa um erro de validação
type ErroValidacao struct {
	Campo   string `json:"campo"`
	Mensagem string `json:"mensagem"`
}

// RespostaErro representa uma resposta de erro da API
type RespostaErro struct {
	Status  int             `json:"status"`
	Mensagem string         `json:"mensagem"`
	Erros   []ErroValidacao `json:"erros,omitempty"`
}

// ValidadorUsuario contém as funções de validação para os campos do usuário
type ValidadorUsuario struct{}

// Valida verifica todos os campos do usuário
func (v *ValidadorUsuario) Valida(u Usuario) []ErroValidacao {
	var erros []ErroValidacao
	
	// Valida nome
	if erro := v.validaNome(u.Nome); erro != nil {
		erros = append(erros, *erro)
	}
	
	// Valida e-mail
	if erro := v.validaEmail(u.Email); erro != nil {
		erros = append(erros, *erro)
	}
	
	// Valida senha
	if erro := v.validaSenha(u.Senha); erro != nil {
		erros = append(erros, *erro)
	}
	
	// Valida telefone
	if erro := v.validaTelefone(u.Telefone); erro != nil {
		erros = append(erros, *erro)
	}
	
	// Valida CPF
	if erro := v.validaCPF(u.CPF); erro != nil {
		erros = append(erros, *erro)
	}
	
	// Valida idade
	if erro := v.validaIdade(u.Idade); erro != nil {
		erros = append(erros, *erro)
	}
	
	return erros
}

// validaNome verifica se o nome é válido
func (v *ValidadorUsuario) validaNome(nome string) *ErroValidacao {
	if nome == "" {
		return &ErroValidacao{
			Campo:   "nome",
			Mensagem: "O nome é obrigatório",
		}
	}
	
	if len(nome) < 3 {
		return &ErroValidacao{
			Campo:   "nome",
			Mensagem: "O nome deve ter pelo menos 3 caracteres",
		}
	}
	
	return nil
}

// validaEmail verifica se o e-mail é válido
func (v *ValidadorUsuario) validaEmail(email string) *ErroValidacao {
	if email == "" {
		return &ErroValidacao{
			Campo:   "email",
			Mensagem: "O e-mail é obrigatório",
		}
	}
	
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return &ErroValidacao{
			Campo:   "email",
			Mensagem: "O e-mail fornecido não é válido",
		}
	}
	
	return nil
}

// validaSenha verifica se a senha é válida
func (v *ValidadorUsuario) validaSenha(senha string) *ErroValidacao {
	if senha == "" {
		return &ErroValidacao{
			Campo:   "senha",
			Mensagem: "A senha é obrigatória",
		}
	}
	
	if len(senha) < 8 {
		return &ErroValidacao{
			Campo:   "senha",
			Mensagem: "A senha deve ter pelo menos 8 caracteres",
		}
	}
	
	// Verifica se a senha tem pelo menos um número
	numeros := regexp.MustCompile(`[0-9]`)
	if !numeros.MatchString(senha) {
		return &ErroValidacao{
			Campo:   "senha",
			Mensagem: "A senha deve conter pelo menos um número",
		}
	}
	
	// Verifica se a senha tem pelo menos uma letra maiúscula
	maiusculas := regexp.MustCompile(`[A-Z]`)
	if !maiusculas.MatchString(senha) {
		return &ErroValidacao{
			Campo:   "senha",
			Mensagem: "A senha deve conter pelo menos uma letra maiúscula",
		}
	}
	
	// Verifica se a senha tem pelo menos um caractere especial
	especiais := regexp.MustCompile(`[@#$%^&+=]`)
	if !especiais.MatchString(senha) {
		return &ErroValidacao{
			Campo:   "senha",
			Mensagem: "A senha deve conter pelo menos um caractere especial (@, #, $, %, ^, &, +, =)",
		}
	}
	
	return nil
}

// validaTelefone verifica se o telefone é válido
func (v *ValidadorUsuario) validaTelefone(telefone string) *ErroValidacao {
	if telefone == "" {
		return &ErroValidacao{
			Campo:   "telefone",
			Mensagem: "O telefone é obrigatório",
		}
	}
	
	// Remove caracteres não numéricos
	limpo := regexp.MustCompile(`\D`).ReplaceAllString(telefone, "")
	
	// Verifica se o telefone tem entre 10 e 11 dígitos (com ou sem DDD)
	if len(limpo) < 10 || len(limpo) > 11 {
		return &ErroValidacao{
			Campo:   "telefone",
			Mensagem: "O telefone deve ter entre 10 e 11 dígitos",
		}
	}
	
	return nil
}

// validaCPF verifica se o CPF é válido
func (v *ValidadorUsuario) validaCPF(cpf string) *ErroValidacao {
	if cpf == "" {
		return &ErroValidacao{
			Campo:   "cpf",
			Mensagem: "O CPF é obrigatório",
		}
	}
	
	// Remove caracteres não numéricos
	limpo := regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")
	
	// Verifica se o CPF tem 11 dígitos
	if len(limpo) != 11 {
		return &ErroValidacao{
			Campo:   "cpf",
			Mensagem: "O CPF deve ter 11 dígitos",
		}
	}
	
	// Verifica se todos os dígitos são iguais
	todosIguais := true
	for i := 1; i < len(limpo); i++ {
		if limpo[i] != limpo[0] {
			todosIguais = false
			break
		}
	}
	
	if todosIguais {
		return &ErroValidacao{
			Campo:   "cpf",
			Mensagem: "CPF inválido",
		}
	}
	
	// Em um sistema real, faríamos a validação completa do algoritmo do CPF
	// Aqui simplificamos por questões didáticas
	
	return nil
}

// validaIdade verifica se a idade é válida
func (v *ValidadorUsuario) validaIdade(idade int) *ErroValidacao {
	if idade < 0 {
		return &ErroValidacao{
			Campo:   "idade",
			Mensagem: "A idade não pode ser negativa",
		}
	}
	
	if idade < 18 {
		return &ErroValidacao{
			Campo:   "idade",
			Mensagem: "O usuário deve ter pelo menos 18 anos",
		}
	}
	
	if idade > 120 {
		return &ErroValidacao{
			Campo:   "idade",
			Mensagem: "Idade inválida",
		}
	}
	
	return nil
}

// Controlador gerencia as operações HTTP
type Controlador struct {
	validador *ValidadorUsuario
	usuarios  map[int]Usuario
	proximoID int
}

// NovoControlador cria um novo controlador
func NovoControlador() *Controlador {
	return &Controlador{
		validador: &ValidadorUsuario{},
		usuarios:  make(map[int]Usuario),
		proximoID: 1,
	}
}

// RespostaErroJSON retorna uma resposta de erro em formato JSON
func RespostaErroJSON(w http.ResponseWriter, status int, mensagem string, erros []ErroValidacao) {
	resp := RespostaErro{
		Status:  status,
		Mensagem: mensagem,
		Erros:   erros,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

// handleCadastro processa o cadastro de novos usuários
func (c *Controlador) handleCadastro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RespostaErroJSON(w, http.StatusMethodNotAllowed, "Método não permitido", nil)
		return
	}
	
	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		RespostaErroJSON(w, http.StatusBadRequest, "Erro ao processar os dados", []ErroValidacao{
			{Campo: "body", Mensagem: "JSON inválido"},
		})
		return
	}
	
	// Valida os dados do usuário
	erros := c.validador.Valida(usuario)
	if len(erros) > 0 {
		RespostaErroJSON(w, http.StatusUnprocessableEntity, "Dados inválidos", erros)
		return
	}
	
	// Em um sistema real, verificaríamos se o e-mail ou CPF já existem
	
	// Atribui um ID e salva o usuário
	usuario.ID = c.proximoID
	c.proximoID++
	
	// Armazena o usuário (em memória)
	c.usuarios[usuario.ID] = usuario
	
	// Prepara a resposta (sem enviar a senha)
	usuario.Senha = ""
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
	
	fmt.Printf("Usuário cadastrado com sucesso: %s (ID: %d)\n", usuario.Nome, usuario.ID)
}

// handleLogin simula o processo de login
func (c *Controlador) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RespostaErroJSON(w, http.StatusMethodNotAllowed, "Método não permitido", nil)
		return
	}
	
	// Estrutura para receber dados de login
	var credenciais struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}
	
	err := json.NewDecoder(r.Body).Decode(&credenciais)
	if err != nil {
		RespostaErroJSON(w, http.StatusBadRequest, "Erro ao processar os dados", []ErroValidacao{
			{Campo: "body", Mensagem: "JSON inválido"},
		})
		return
	}
	
	// Valida email e senha
	erros := []ErroValidacao{}
	
	if credenciais.Email == "" {
		erros = append(erros, ErroValidacao{Campo: "email", Mensagem: "O e-mail é obrigatório"})
	}
	
	if credenciais.Senha == "" {
		erros = append(erros, ErroValidacao{Campo: "senha", Mensagem: "A senha é obrigatória"})
	}
	
	if len(erros) > 0 {
		RespostaErroJSON(w, http.StatusUnprocessableEntity, "Credenciais inválidas", erros)
		return
	}
	
	// Em um sistema real, consultaríamos o banco de dados para buscar o usuário
	// e verificar a senha com um algoritmo de hash
	
	// Aqui vamos simular encontrando o usuário pelo email
	var usuarioAutenticado Usuario
	encontrado := false
	
	for _, u := range c.usuarios {
		if strings.EqualFold(u.Email, credenciais.Email) && u.Senha == credenciais.Senha {
			usuarioAutenticado = u
			encontrado = true
			break
		}
	}
	
	if !encontrado {
		RespostaErroJSON(w, http.StatusUnauthorized, "Credenciais inválidas", nil)
		return
	}
	
	// Em um sistema real, geraríamos um token JWT
	// Aqui, vamos simular isso retornando dados de usuário (sem a senha)
	usuarioAutenticado.Senha = ""
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"usuario": usuarioAutenticado,
		"token":   "simulacao-token-jwt", // Simulação simples
	})
	
	fmt.Printf("Usuário autenticado: %s (ID: %d)\n", usuarioAutenticado.Nome, usuarioAutenticado.ID)
}

// handleUsuarios retorna a lista de usuários cadastrados
func (c *Controlador) handleUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RespostaErroJSON(w, http.StatusMethodNotAllowed, "Método não permitido", nil)
		return
	}
	
	// Converte o map para slice
	usuarios := make([]Usuario, 0, len(c.usuarios))
	for _, u := range c.usuarios {
		// Não envia a senha
		u.Senha = ""
		usuarios = append(usuarios, u)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func main() {
	controlador := NovoControlador()
	
	// Configura as rotas
	http.HandleFunc("/cadastro", controlador.handleCadastro)
	http.HandleFunc("/login", controlador.handleLogin)
	http.HandleFunc("/usuarios", controlador.handleUsuarios)
	
	// Inicia o servidor
	fmt.Println("Servidor de validação iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
} 