// algoritmo concurrente
package main

import (
	"fmt"
	"time"
)

var n int

func p() {
	k1 := 1
	n = k1
}

func q() {
	k2 := 2
	n = k2
}

func r() {
	k3 := 3
	n = k3
}

func main() {
	n = 0
	//crear los goroutines
	go p()
	go q()
	go r()

	time.Sleep(time.Millisecond * 600) //pausa del proceso main

	fmt.Printf("El valor final de n es %d\n", n)

}
