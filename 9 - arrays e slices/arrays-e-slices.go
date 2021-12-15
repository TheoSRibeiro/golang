package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	//Arrays
	var array1 [5]string
	array1[0] = "Pos1"

	fmt.Println(array1)

	array2 := [5]string{"pos1", "pos2", "pos3", "pos4", "pos5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	//Slices
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "pos_alterada"
	fmt.Println(slice2)

	//Arrays Internos

	fmt.Println("Arrays Internos")

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	fmt.Println(slice3)

	slice3 = append(slice3, 6)
	fmt.Println(slice3)

	fmt.Println("tamanho slice 3 - ", len(slice3))    //length
	fmt.Println("capacidade slice 3 - ", cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println("tamanho slice 4 - ", len(slice4))    //length
	fmt.Println("capacidade slice 4 - ", cap(slice4)) // capacidade

	slice4 = append(slice4, 10)
	fmt.Println("tamanho slice 4 - ", len(slice4))    //length
	fmt.Println("capacidade slice 4 - ", cap(slice4)) // capacidade
}
