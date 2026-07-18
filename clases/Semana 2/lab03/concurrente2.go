package main

import (
	"fmt"
	"time"
)

var n int //recurso compartido

func p() {
	k1 := 1
	time.Sleep(2 * time.Nanosecond)
	n = k1 //instrucción atómica
}
func q() {
	k2 := 2
	n = k2 //instrucción atómica
}

// proceso principal
func main() {
	n = 0

	//manejar la concurrencia
	//goroutines
	go p() //goroutine = proc concurrente
	go q() //goroutine = proc concurrente
	//función anónima para proceso concurrente
	go func() {
		k3 := 3
		time.Sleep(2500 * time.Nanosecond)
		n = k3
	}()

	time.Sleep(200 * time.Millisecond) //retardo

	fmt.Printf("El valor final de n es %d\n", n)

}
