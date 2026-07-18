package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	// Trabajador codicioso: expande su posesión del bloqueo innecesariamente
	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()               // Adquiere el bloqueo
			time.Sleep(3 * time.Nanosecond) // Simula trabajo extenso dentro del bloqueo
			sharedLock.Unlock()             // Lo libera
			count++
		}
		fmt.Printf("Trabajador codicioso ejecutó %v ciclos de trabajo\n", count)
	}

	// Trabajador educado: intenta liberar el bloqueo frecuentemente para ser justo
	politeWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			// Divide su trabajo en secciones críticas más pequeñas
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		fmt.Printf("Trabajador educado ejecutó %v ciclos de trabajo\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
