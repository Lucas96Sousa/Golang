package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1543.6,
			"Guga Pereira":   845.6,
		},
		"J": {
			"Jose Joao": 11325.63,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
