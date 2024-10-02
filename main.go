package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operation string

	fmt.Println("Calculadora Simples")
	fmt.Println("--------------------")
	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&num1)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)

	fmt.Println("Escolha a operação (+, -, *, /):")
	fmt.Scanln(&operation)

	fmt.Println("voce escolheu: ", operation)

	switch operation {
	case "+":
		result := num1 + num2
		fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", num1, num2, result)
	case "-":
		result := num1 - num2
		fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", num1, num2, result)
	case "*":
		result := num1 * num2
		fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", num1, num2, result)
	case "/":
		if num2 != 0 {
			result := num1 / num2
			fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", num1, num2, result)
		} else {
			fmt.Println("Erro: Divisão por zero não é permitida.")
		}
	default:
		fmt.Println("Operação inválida. Por favor, escolha uma operação válida.")
	}
}
