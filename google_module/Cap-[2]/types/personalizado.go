package main

import "fmt"

type nota float64

func (n nota) entre(ini, fim float64) bool {
	return float64(n) >= ini && float64(n) <= fim
}

func notaParaConceito(n nota) string {

	switch {
	case 9.0, 10.0:
		return "A"
	case 7.0, 8.99:
		return "B"
	case 6.0, 7.99:
		return "C"
	case 5.0, 6.99:
		return "D"
	case 3.0, 4.99:
		return "E"
	default:
		return "F"

	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
