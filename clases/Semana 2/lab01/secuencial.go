package main

import "fmt"

func main() {
	n := 0
	k1 := 1
	k2 := 2
	n = k1
	n = k2

	fmt.Println("El valor final de n es ", n) //resultado deterministico
}
