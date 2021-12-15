package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "string",
		true:         float64(12),
	}
	fmt.Println(mapa)
}
