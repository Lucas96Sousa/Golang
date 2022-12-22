package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(2)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"String":     "String",
		true:         float32(12),
	}

	fmt.Println(mapa)

}
