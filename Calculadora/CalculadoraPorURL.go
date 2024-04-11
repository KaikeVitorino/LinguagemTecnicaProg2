package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv" // Importar strconv para conversão de string para float64
	"time"
)

// Calculator é uma estrutura para a calculadora.
type Calculator struct{}

// ServeHTTP é o método para lidar com solicitações HTTP.
func (c *Calculator) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Verifica se o método da solicitação é GET
	if req.Method != "GET" {
		// Se não for, retorna um erro de método não permitido
		http.Error(res, "Metodo esquisito ai! Nao pode!", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os parâmetros da consulta da URL
	params := req.URL.Query()
	action := params.Get("action") // Ação da operação (add, sub, mul, div)
	num1 := params.Get("num1")     // Primeiro número
	num2 := params.Get("num2")     // Segundo número

	// Verifica se os parâmetros obrigatórios estão presentes
	if action == "" || num1 == "" || num2 == "" {
		// Se algum estiver faltando, retorna um erro de solicitação inválida
		http.Error(res, "Ta faltando coisa ai!", http.StatusBadRequest)
		return
	}

	// Mapa de operações matemáticas
	operations := map[string]func(float64, float64) float64{
		"add": func(a, b float64) float64 { return a + b }, // Adição
		"sub": func(a, b float64) float64 { return a - b }, // Subtração
		"mul": func(a, b float64) float64 { return a * b }, // Multiplicação
		"div": func(a, b float64) float64 { // Divisão
			if b == 0 {
				http.Error(res, "0 nao pode ser dividido!", http.StatusBadRequest)
				return 0
			}
			return a / b
		},
	}

	// Obtém a função de operação correspondente
	operationFunc, ok := operations[action]
	if !ok {
		// Se a ação não estiver no mapa, retorna um erro de ação inválida
		http.Error(res, "Nao conheco essa matematica ai nao", http.StatusBadRequest)
		return
	}

	// Converte os números de string para float64
	num1Float, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		// Se houver um erro ao converter o primeiro número, retorna um erro de número inválido
		http.Error(res, "Numero 1 ta errado!", http.StatusBadRequest)
		return
	}

	num2Float, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		// Se houver um erro ao converter o segundo número, retorna um erro de número inválido
		http.Error(res, "Numero 2 ta errado!", http.StatusBadRequest)
		return
	}

	// Executa a operação com os números fornecidos
	result := operationFunc(num1Float, num2Float)

	// Monta a resposta em formato JSON
	response := struct {
		Result float64 `json:"resultado"`
	}{
		Result: result,
	}

	// Define o cabeçalho Content-Type como aplicação/json
	res.Header().Set("Content-Type", "application/json")
	// Codifica a resposta JSON e a envia para o cliente
	json.NewEncoder(res).Encode(response)
}

func main() {
	// Cria uma instância do struct Calculator
	calculator := &Calculator{}

	// Configuração do servidor HTTP
	s := &http.Server{
		Addr:         "localhost:8080", // Endereço do servidor
		Handler:      calculator,       // Manipulador de solicitações HTTP (Calculator)
		ReadTimeout:  10 * time.Second, // Tempo máximo para ler a solicitação
		WriteTimeout: 10 * time.Second, // Tempo máximo para escrever a resposta
	}

	// Inicia o servidor e verifica se há erros
	log.Fatal(s.ListenAndServe())
}
