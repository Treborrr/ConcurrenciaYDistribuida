// sincronizar
package main

import (
	"fmt"
	"time"
)

func trabajador(done chan bool) {
	fmt.Println("Trabajando.....")
	time.Sleep(time.Millisecond * 500) //simular el tiempo de trabajo
	fmt.Println("Done")

	done <- true
}

func main() {
	done := make(chan bool)

	go trabajador(done)

	<-done
	fmt.Println("Fin proceso main ")
	//fin
}
