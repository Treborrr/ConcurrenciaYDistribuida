// sincronización usando grupos de espera de PC (goroutines)
package main

import (
	"fmt"
	"sync"
)

func main() {
	//declaración
	//crear wg para la espera de todos los goroutines
	var wg sync.WaitGroup
	var mu sync.Mutex //exclusión mutua a SC

	var contador int //recurso compartido en la SC

	ngoroutines := 5

	//Inicio:
	wg.Add(ngoroutines)

	//lanzar los goroutines
	for i := 0; i < ngoroutines; i++ {
		//función anónima
		go func(x int) {

			mu.Lock()

			contador++

			mu.Unlock()
			fmt.Printf("el goroutine %d incrmenta el contador a %d\n", x, contador)
			wg.Done() //para confirmar q el goroutine finalizó su tarea
		}(i)
	}

	//esperar hasta q todos los goroutines finalicen su trabajo
	wg.Wait()
	//Fin
	//impresión final
	fmt.Printf("El valor final del contador es %d\n", contador)
}
