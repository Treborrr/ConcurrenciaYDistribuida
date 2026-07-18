// variable global
package main

import (
	"fmt"
	"time"
)

var n int

func proceso() {
	var temp int
	temp = n
	time.Sleep(time.Millisecond * 10) //retardo
	n = temp + 1
}
func main() {

	n = 0

	go proceso() //p
	go proceso() //q

	time.Sleep(time.Millisecond * 500)

	fmt.Printf("El valor final de n es %d\n", n)

}
