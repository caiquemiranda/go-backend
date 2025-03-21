package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	uploadPath = "./uploads"
	maxSize    = 10 << 20 // 10 MB
)

// Estrutura para informações do arquivo
type FileInfo struct {
	Name         string
	Size         string
	LastModified string
	Path         string
}

// Template para página inicial
var indexTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Upload de Arquivos com Go</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        h1, h2 {
            color: #333;
        }
        form {
            background-color: #f5f5f5;
            padding: 20px;
            border-radius: 5px;
            margin-bottom: 20px;
        }
        .file-list {
            background-color: #f9f9f9;
            padding: 15px;
            border-radius: 5px;
        }
        .file-item {
            border-bottom: 1px solid #ddd;
            padding: 10px 0;
        }
        .alert {
            padding: 10px;
            background-color: #f44336;
            color: white;
            border-radius: 5px;
            margin-bottom: 15px;
        }
        .success {
            padding: 10px;
            background-color: #4CAF50;
            color: white;
            border-radius: 5px;
            margin-bottom: 15px;
        }
    </style>
</head>
<body>
    <h1>Upload de Arquivos com Go</h1>
    
    {{if .Error}}
    <div class="alert">{{.Error}}</div>
    {{end}}
    
    {{if .Success}}
    <div class="success">{{.Success}}</div>
    {{end}}
    
    <form action="/upload" method="post" enctype="multipart/form-data">
        <h2>Upload de Arquivo</h2>
        <p>Selecione um arquivo para enviar (máximo 10MB):</p>
        <input type="file" name="arquivo" required>
        <p>
            <button type="submit">Enviar Arquivo</button>
        </p>
    </form>
    
    <div class="file-list">
        <h2>Arquivos Enviados</h2>
        {{if .Files}}
            {{range .Files}}
            <div class="file-item">
                <p><strong>Nome:</strong> {{.Name}}</p>
                <p><strong>Tamanho:</strong> {{.Size}}</p>
                <p><strong>Última modificação:</strong> {{.LastModified}}</p>
                <p><a href="/files/{{.Name}}" target="_blank">Visualizar</a> | 
                   <a href="/download/{{.Name}}">Download</a> | 
                   <a href="/delete/{{.Name}}" onclick="return confirm('Tem certeza que deseja excluir este arquivo?');">Excluir</a></p>
            </div>
            {{end}}
        {{else}}
            <p>Nenhum arquivo enviado ainda.</p>
        {{end}}
    </div>
</body>
</html>
`

// Handler para página inicial
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		http.Error(w, "Erro ao processar template", http.StatusInternalServerError)
		log.Printf("Erro ao processar template: %v", err)
		return
	}

	// Listar arquivos do diretório de uploads
	files, err := listFiles()
	if err != nil {
		log.Printf("Erro ao listar arquivos: %v", err)
	}

	// Dados para o template
	data := struct {
		Files   []FileInfo
		Error   string
		Success string
	}{
		Files:   files,
		Error:   r.URL.Query().Get("error"),
		Success: r.URL.Query().Get("success"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erro ao renderizar página", http.StatusInternalServerError)
		log.Printf("Erro ao renderizar página: %v", err)
	}
}

// Handler para upload de arquivos
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/?error=Método não permitido", http.StatusSeeOther)
		return
	}

	// Parse do formulário multipart
	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		http.Redirect(w, r, "/?error=Erro ao processar o arquivo. Tamanho máximo: 10MB", http.StatusSeeOther)
		return
	}

	// Obter o arquivo do formulário
	file, handler, err := r.FormFile("arquivo")
	if err != nil {
		http.Redirect(w, r, "/?error=Erro ao obter o arquivo", http.StatusSeeOther)
		return
	}
	defer file.Close()

	// Validar tipo de arquivo (opcional)
	// Aqui você poderia adicionar validação para tipos de arquivo específicos

	// Garantir que o diretório de upload existe
	err = os.MkdirAll(uploadPath, 0755)
	if err != nil {
		http.Redirect(w, r, "/?error=Erro ao criar diretório de upload", http.StatusSeeOther)
		return
	}

	// Criar arquivo de destino
	filename := filepath.Join(uploadPath, handler.Filename)
	dst, err := os.Create(filename)
	if err != nil {
		http.Redirect(w, r, "/?error=Erro ao criar arquivo no servidor", http.StatusSeeOther)
		return
	}
	defer dst.Close()

	// Copiar conteúdo do arquivo enviado para o destino
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Redirect(w, r, "/?error=Erro ao salvar arquivo", http.StatusSeeOther)
		return
	}

	// Redirecionar com mensagem de sucesso
	mensagem := fmt.Sprintf("Arquivo %s enviado com sucesso!", handler.Filename)
	http.Redirect(w, r, "/?success="+mensagem, http.StatusSeeOther)
}

// Handler para visualização de arquivos
func fileHandler(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Base(r.URL.Path[7:]) // Remove o prefixo "/files/"
	filepath := filepath.Join(uploadPath, filename)

	// Verificar se o arquivo existe
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Servir o arquivo
	http.ServeFile(w, r, filepath)
}

// Handler para download de arquivos
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Base(r.URL.Path[10:]) // Remove o prefixo "/download/"
	filepath := filepath.Join(uploadPath, filename)

	// Verificar se o arquivo existe
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Configurar cabeçalhos para download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	// Servir o arquivo
	http.ServeFile(w, r, filepath)
}

// Handler para exclusão de arquivos
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Base(r.URL.Path[8:]) // Remove o prefixo "/delete/"
	filepath := filepath.Join(uploadPath, filename)

	// Verificar se o arquivo existe
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.Redirect(w, r, "/?error=Arquivo não encontrado", http.StatusSeeOther)
		return
	}

	// Excluir o arquivo
	err := os.Remove(filepath)
	if err != nil {
		http.Redirect(w, r, "/?error=Erro ao excluir arquivo", http.StatusSeeOther)
		return
	}

	// Redirecionar com mensagem de sucesso
	http.Redirect(w, r, "/?success=Arquivo excluído com sucesso", http.StatusSeeOther)
}

// Função para listar arquivos no diretório de uploads
func listFiles() ([]FileInfo, error) {
	var fileInfos []FileInfo

	// Verificar se o diretório existe
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadPath, 0755); err != nil {
			return nil, err
		}
		return fileInfos, nil // Retorna lista vazia se o diretório foi criado agora
	}

	// Ler diretório de uploads
	files, err := os.ReadDir(uploadPath)
	if err != nil {
		return nil, err
	}

	// Criar informações sobre cada arquivo
	for _, file := range files {
		if file.IsDir() {
			continue // Ignorar diretórios
		}

		info, err := file.Info()
		if err != nil {
			continue // Ignorar arquivos com erro
		}

		// Formatar tamanho do arquivo
		size := formatFileSize(info.Size())

		// Formatar data de modificação
		lastModified := info.ModTime().Format("02/01/2006 15:04:05")

		fileInfos = append(fileInfos, FileInfo{
			Name:         info.Name(),
			Size:         size,
			LastModified: lastModified,
			Path:         filepath.Join(uploadPath, info.Name()),
		})
	}

	return fileInfos, nil
}

// Função para formatar o tamanho do arquivo
func formatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func main() {
	// Garantir que o diretório de uploads existe
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		log.Fatalf("Erro ao criar diretório de uploads: %v", err)
	}

	// Configurar rotas
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/files/", fileHandler)
	http.HandleFunc("/download/", downloadHandler)
	http.HandleFunc("/delete/", deleteHandler)

	// Iniciar servidor
	porta := ":8080"
	fmt.Printf("Servidor rodando em http://localhost%s\n", porta)
	log.Fatal(http.ListenAndServe(porta, nil))
} 