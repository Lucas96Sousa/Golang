package main

import "fmt"

func main() {
	funcionarios := map[string]float64{
		"Jose Joao":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcionarios["Rafael Filho"] = 1350.0
	delete(funcionarios, "Inexistente")
	for nome, salario := range funcionarios {
		fmt.Printf("Nome: %s [----] Salário: %.2f\n", nome, salario)
	}
}
