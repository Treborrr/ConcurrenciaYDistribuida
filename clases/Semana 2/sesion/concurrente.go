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

func main() {
	n = 0
	//crear los goroutines
	go p()
	go q()

	time.Sleep(time.Millisecond * 600)

	fmt.Printf("El valor final de n es %d\n", n)

}
