package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu sync.Mutex
	v  int
}

func main() {

	var wg sync.WaitGroup //agrupador de goroutines

	printSum := func(v1, v2 *value) {
		defer wg.Done()      //indica que esta goroutine ha terminado
		v1.mu.Lock()         //bloquea el mutex de v1 nuestro primer recurso
		defer v1.mu.Unlock() //desbloquea el mutex de v1 al finalizar la función

		//simulamos un retardo para aumentar la probabilidad de que ocurra el deadlock
		time.Sleep(2 * time.Second)

		v2.mu.Lock()         //bloquea el mutex de v2 nuestro segundo recurso
		defer v2.mu.Unlock() //desbloquea el mutex de v2 al finalizar la función

		fmt.Printf("La suma de %d y %d es %d\n", v1.v, v2.v, v1.v+v2.v)
	}

	var v1, v2 value
	wg.Add(2)             //indica que vamos a esperar a 2 goroutines
	go printSum(&v1, &v2) //primera goroutine que intenta bloquear v1 y luego v2
	go printSum(&v2, &v1) //segunda goroutine que intenta bloquear v2 y luego v1

	wg.Wait() //espera a que ambas goroutines terminen

	fmt.Println("Programa terminado")
}
