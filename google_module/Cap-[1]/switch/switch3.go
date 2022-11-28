package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(2))
	fmt.Println(tipo("opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))

}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "Type not detected"
	}
}
