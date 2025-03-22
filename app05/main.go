package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Tarefa representa uma tarefa no sistema de gerenciamento
type Tarefa struct {
	ID          int       `json:"id"`
	Titulo      string    `json:"titulo"`
	Descricao   string    `json:"descricao"`
	Concluida   bool      `json:"concluida"`
	DataCriacao time.Time `json:"dataCriacao"`
	Prioridade  int       `json:"prioridade"` // 1-baixa, 2-média, 3-alta
}

// Gerenciador é responsável pelo armazenamento e operações nas tarefas
type Gerenciador struct {
	tarefas   map[int]Tarefa
	proximoID int
	mutex     sync.RWMutex
}

// NovoGerenciador cria uma nova instância do gerenciador de tarefas
func NovoGerenciador() *Gerenciador {
	return &Gerenciador{
		tarefas:   make(map[int]Tarefa),
		proximoID: 1,
		mutex:     sync.RWMutex{},
	}
}

// Adicionar insere uma nova tarefa e retorna a tarefa criada
func (g *Gerenciador) Adicionar(t Tarefa) Tarefa {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	t.ID = g.proximoID
	t.DataCriacao = time.Now()
	if t.Prioridade < 1 || t.Prioridade > 3 {
		t.Prioridade = 1 // Define prioridade baixa como padrão
	}
	
	g.tarefas[t.ID] = t
	g.proximoID++
	
	return t
}

// Obter retorna uma tarefa pelo ID
func (g *Gerenciador) Obter(id int) (Tarefa, bool) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	tarefa, existe := g.tarefas[id]
	return tarefa, existe
}

// ListarTodas retorna todas as tarefas
func (g *Gerenciador) ListarTodas() []Tarefa {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	tarefas := make([]Tarefa, 0, len(g.tarefas))
	for _, tarefa := range g.tarefas {
		tarefas = append(tarefas, tarefa)
	}
	
	return tarefas
}

// Atualizar modifica uma tarefa existente
func (g *Gerenciador) Atualizar(id int, t Tarefa) (Tarefa, bool) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	if _, existe := g.tarefas[id]; !existe {
		return Tarefa{}, false
	}
	
	// Preserva o ID e a data de criação
	t.ID = id
	t.DataCriacao = g.tarefas[id].DataCriacao
	
	// Atualiza a tarefa
	g.tarefas[id] = t
	
	return t, true
}

// MarcarComoConcluida atualiza o status de uma tarefa para concluída
func (g *Gerenciador) MarcarComoConcluida(id int) (Tarefa, bool) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	if tarefa, existe := g.tarefas[id]; existe {
		tarefa.Concluida = true
		g.tarefas[id] = tarefa
		return tarefa, true
	}
	
	return Tarefa{}, false
}

// Remover exclui uma tarefa pelo ID
func (g *Gerenciador) Remover(id int) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	
	if _, existe := g.tarefas[id]; !existe {
		return false
	}
	
	delete(g.tarefas, id)
	return true
}

// ListarPorPrioridade retorna tarefas com a prioridade especificada
func (g *Gerenciador) ListarPorPrioridade(prioridade int) []Tarefa {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	tarefas := make([]Tarefa, 0)
	for _, tarefa := range g.tarefas {
		if tarefa.Prioridade == prioridade {
			tarefas = append(tarefas, tarefa)
		}
	}
	
	return tarefas
}

// ListarPorStatus retorna tarefas com base no status de conclusão
func (g *Gerenciador) ListarPorStatus(concluida bool) []Tarefa {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	
	tarefas := make([]Tarefa, 0)
	for _, tarefa := range g.tarefas {
		if tarefa.Concluida == concluida {
			tarefas = append(tarefas, tarefa)
		}
	}
	
	return tarefas
}

func main() {
	gerenciador := NovoGerenciador()
	
	// Adiciona algumas tarefas de exemplo
	gerenciador.Adicionar(Tarefa{
		Titulo:     "Estudar Go",
		Descricao:  "Aprender sobre servidores HTTP e manipulação de JSON",
		Prioridade: 3,
	})
	
	gerenciador.Adicionar(Tarefa{
		Titulo:     "Fazer compras",
		Descricao:  "Comprar frutas e legumes",
		Prioridade: 2,
	})
	
	gerenciador.Adicionar(Tarefa{
		Titulo:     "Ler livro",
		Descricao:  "Terminar de ler o livro sobre programação",
		Prioridade: 1,
	})
	
	// Configura as rotas
	http.HandleFunc("/tarefas", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Verifica se há filtros
			query := r.URL.Query()
			
			var tarefas []Tarefa
			
			if prioridade := query.Get("prioridade"); prioridade != "" {
				p, err := strconv.Atoi(prioridade)
				if err != nil || p < 1 || p > 3 {
					http.Error(w, "Prioridade inválida", http.StatusBadRequest)
					return
				}
				tarefas = gerenciador.ListarPorPrioridade(p)
			} else if status := query.Get("concluida"); status != "" {
				concluida := status == "true"
				tarefas = gerenciador.ListarPorStatus(concluida)
			} else {
				tarefas = gerenciador.ListarTodas()
			}
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tarefas)
			
		case http.MethodPost:
			var tarefa Tarefa
			if err := json.NewDecoder(r.Body).Decode(&tarefa); err != nil {
				http.Error(w, "Dados inválidos", http.StatusBadRequest)
				return
			}
			
			tarefaCriada := gerenciador.Adicionar(tarefa)
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(tarefaCriada)
			
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
	
	http.HandleFunc("/tarefas/", func(w http.ResponseWriter, r *http.Request) {
		// Extrai o ID da URL
		partes := strings.Split(r.URL.Path, "/")
		if len(partes) != 3 {
			http.Error(w, "URL inválida", http.StatusBadRequest)
			return
		}
		
		id, err := strconv.Atoi(partes[2])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		
		switch r.Method {
		case http.MethodGet:
			tarefa, existe := gerenciador.Obter(id)
			if !existe {
				http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tarefa)
			
		case http.MethodPut:
			var tarefa Tarefa
			if err := json.NewDecoder(r.Body).Decode(&tarefa); err != nil {
				http.Error(w, "Dados inválidos", http.StatusBadRequest)
				return
			}
			
			tarefaAtualizada, existe := gerenciador.Atualizar(id, tarefa)
			if !existe {
				http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tarefaAtualizada)
			
		case http.MethodPatch:
			// Usado apenas para marcar como concluída
			if r.URL.Query().Get("concluida") == "true" {
				tarefa, existe := gerenciador.MarcarComoConcluida(id)
				if !existe {
					http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
					return
				}
				
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(tarefa)
			} else {
				http.Error(w, "Operação não suportada", http.StatusBadRequest)
			}
			
		case http.MethodDelete:
			sucesso := gerenciador.Remover(id)
			if !sucesso {
				http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
				return
			}
			
			w.WriteHeader(http.StatusNoContent)
			
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
	
	// Inicia o servidor
	fmt.Println("Servidor de gerenciamento de tarefas iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
} 