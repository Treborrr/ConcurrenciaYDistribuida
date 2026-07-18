package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	CAPACIDAD_MAX = 3  // máximo de técnicos simultáneos en la sala
	NUM_TECNICOS  = 10 // total de técnicos que intentan entrar
)

func main() {
	// Canal con buffer = semáforo multiplex (capacidad máxima N)
	sala := make(chan struct{}, CAPACIDAD_MAX)

	var wg sync.WaitGroup
	var enSala int32 // contador atómico para mostrar ocupación actual

	fmt.Printf("=== Sala de Servidores: capacidad máxima %d técnicos ===\n\n", CAPACIDAD_MAX)

	for i := 1; i <= NUM_TECNICOS; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("Técnico %2d: esperando permiso para entrar...\n", id)
			sala <- struct{}{} // wait: toma un cupo (bloquea si lleno)

			ocupados := atomic.AddInt32(&enSala, 1)
			fmt.Printf("Técnico %2d: ENTRÓ  | técnicos en sala: %d/%d\n", id, ocupados, CAPACIDAD_MAX)

			// Trabajo en la sala de servidores
			time.Sleep(time.Duration(rand.Intn(400)+200) * time.Millisecond)

			ocupados = atomic.AddInt32(&enSala, -1)
			fmt.Printf("Técnico %2d: SALIÓ  | técnicos en sala: %d/%d\n", id, ocupados, CAPACIDAD_MAX)
			<-sala // signal: libera el cupo
		}(i)
	}

	wg.Wait()
	fmt.Printf("\n[OK] Todos los técnicos completaron su trabajo. Máximo simultáneo garantizado: %d\n", CAPACIDAD_MAX)
}
