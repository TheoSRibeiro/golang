package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usr := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	fmt.Println(usr)

	usr2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus I",
		},
	}
	fmt.Println(usr2)

	delete(usr2, "nome")
	fmt.Println(usr2)

	usr2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usr2)

}
