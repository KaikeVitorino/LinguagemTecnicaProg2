package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// CalculadoraInterface é uma estrutura para a calculadora.
type CalculadoraInterface struct {
	operations map[string]func(float64, float64) (float64, error)
}

// ServeHTTP é o método para lidar com solicitações HTTP.
func (c *CalculadoraInterface) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Verifica se o método da solicitação é POST
	if req.Method != "POST" {
		// Se não for, exibe o formulário HTML
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Calculadora? FOI kkkkk</title>
		</head>
		<body>
			<h2>Calculadora? FOI kkkkk</h2>
			<form method="post">
				<input type="text" name="num1" placeholder="Primeiro numero"><br>
				<input type="text" name="num2" placeholder="Segundo numero"><br>
				<select name="operacao">
					<option value="add">Adicao (+)</option>
					<option value="sub">Subtracao (-)</option>
					<option value="mul">Multiplicacao (*)</option>
					<option value="div">Divisao (/)</option>
				</select><br>
				<input type="submit" value="Run">
			</form>
		</body>
		</html>
		`
		// Escreve a página HTML de formulário como resposta
		fmt.Fprintln(res, html)
		return
	}

	// Parse do formulário
	err := req.ParseForm()
	if err != nil {
		http.Error(res, "Deu ruim em algum lugar", http.StatusBadRequest)
		return
	}

	// Obtém os valores dos campos do formulário
	num1, err := strconv.ParseFloat(req.Form.Get("num1"), 64)
	if err != nil {
		http.Error(res, "Primeiro numero invalido", http.StatusBadRequest)
		return
	}

	num2, err := strconv.ParseFloat(req.Form.Get("num2"), 64)
	if err != nil {
		http.Error(res, "Segundo numero invalido", http.StatusBadRequest)
		return
	}

	operation := req.Form.Get("operacao")

	// Verifica se a operação eh válida
	operationFunc, ok := c.operations[operation]
	if !ok {
		http.Error(res, "Essa opcao ai ta bugada", http.StatusBadRequest)
		return
	}

	// Realiza a operacao selecionada
	result, err := operationFunc(num1, num2)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Escreve o resultado como resposta
	fmt.Fprintf(res, "Result: %f", result)
}

func main() {
	CalculadoraInterface := &CalculadoraInterface{
		operations: map[string]func(float64, float64) (float64, error){
			"add": func(a, b float64) (float64, error) { return a + b, nil }, // Adição
			"sub": func(a, b float64) (float64, error) { return a - b, nil }, // Subtração
			"mul": func(a, b float64) (float64, error) { return a * b, nil }, // Multiplicação
			"div": func(a, b float64) (float64, error) { // Divisão
				if b == 0 || a == 0 {
					return 0, fmt.Errorf("Nao pode dividir por 0 o seu chupaengole!")
				}
				return a / b, nil
			},
		},
	}

	// Configuração do servidor HTTP
	http.Handle("/", CalculadoraInterface)

	// Inicia o servidor e verifica se há erros
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
