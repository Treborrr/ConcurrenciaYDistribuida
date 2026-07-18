package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var resA, resB sync.Mutex
	var wg sync.WaitGroup

	//función que simula un trabajador intentando adquirir ambos recursos
	worker := func(name string, mu1, mu2 *sync.Mutex) {
		defer wg.Done()
		for {
			mu1.Lock() // intenta adquirir el primer recurso
			fmt.Printf("%s: Bloqueó el recurso 1. Intentando adquirir el recurso 2...\n", name)

			//retardo para aumentar la probabilidad de que ambos trabajadores se bloqueen mutuamente
			//y entren en un estado de livelock
			time.Sleep(10 * time.Millisecond)

			if mu2.TryLock() { // intenta adquirir el segundo recurso sin bloquearse
				fmt.Printf("%s: Adquirió ambos recursos. Trabajando...\n", name)
				mu2.Unlock() //liberar el segundo recurso después de usarlo
				mu1.Unlock() //liberar el primer recurso después de usarlo
				return
			}

			fmt.Printf("%s: No pudo adquirir el recurso 2. Liberando el recurso 1 y reintentando...\n", name)

			mu1.Unlock()                      //liberar el primer recurso para permitir que el otro trabajador lo adquiera
			time.Sleep(10 * time.Millisecond) //retardo antes de reintentar
		}
	}
	wg.Add(2)
	go worker("Trabajador 1", &resA, &resB)
	go worker("Trabajador 2", &resB, &resA)
	wg.Wait()
}
