// Esquema de sincronización de procesos (goroutines)
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//declarar la estructura mutex
	var mu sync.Mutex
	var strIn string //variable para parar la ejecución

	//lanzar un conjunto de goroutines
	for i := 0; i < 10; i++ {

		//función anónima
		go func(x int) {
			//lógica de MU
			//parte que se bloquea, puede ser de recurso compartido
			mu.Lock()

			fmt.Printf("Proceso %d, inicia SC\n", x)
			//simular hacer algo
			time.Sleep(time.Millisecond * 60)
			fmt.Printf("Proceso %d, finaliza SC\n", x)

			mu.Unlock()

		}(i)
	}

	//pausa
	fmt.Scanln(&strIn) //esperando el input por la consola para finalizar

}
