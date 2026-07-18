// probando los procesos intercalan en bucle de 10
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n int

func pausa() {
	t := rand.Intn(50) + 50
	time.Sleep(time.Nanosecond * time.Duration(t))
}

func proceso() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		pausa()
		n = temp + 1
		pausa()
	}
}
func main() {
	n = 0

	go proceso()
	go proceso()

	//retardo
	time.Sleep(time.Millisecond * 500)

	fmt.Printf("El valor de n es %d\n", n)

}
