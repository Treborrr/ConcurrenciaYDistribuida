package main

import (
	"fmt"
	"time"
)

//caso condición de carrera

func main() {
	var x int //recurso compartido

	go func() {
		x = 1 // escribir en el recurso compartido
	}()

	go func() {
		//retardo
		time.Sleep(100 * time.Nanosecond)
		fmt.Println(x) // leer el recurso compartido
	}()

	//esperar a que las goroutines terminen (no es una buena práctica, pero sirve para este ejemplo)
	time.Sleep(2 * time.Second)
}
