package main

import "fmt"

var n int

func init() {
	fmt.Println("INIT")
	n = 10
}

func main() {
	fmt.Println("MAIN")
	fmt.Println(n)
}
