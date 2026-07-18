package main

import (
	"fmt"
	"time"
)

var n int

func P() {
	k1 := -1
	n = k1 //p1: ist atómica
	time.Sleep(30 * time.Millisecond)
	n = k1 + 1 //p2: ist atómica
}
func Q() {
	k2 := 2
	n = k2 //q1: ist atómica
}

func main() {
	n = 0
	//creación de goroutines
	go P()
	go Q()

	//retardo
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("El valor final de n es %d\n", n)
}
