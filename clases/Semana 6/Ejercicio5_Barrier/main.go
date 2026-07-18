package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Barrera reutilizable de dos molinetes (two-turnstile reusable barrier).
// Permite sincronizar N goroutines en múltiples fases consecutivas sin deadlock.
type Barrera struct {
	n         int
	count     int
	mu        sync.Mutex
	molinete1 chan struct{} // inicia cerrado (0)
	molinete2 chan struct{} // inicia abierto (1)
}

func NuevaBarrera(n int) *Barrera {
	b := &Barrera{
		n:         n,
		molinete1: make(chan struct{}, 1),
		molinete2: make(chan struct{}, 1),
	}
	b.molinete2 <- struct{}{} // molinete2 empieza abierto
	return b
}

// Esperar bloquea hasta que los n trabajadores lleguen al punto de barrera.
func (b *Barrera) Esperar() {
	// --- Fase de llegada ---
	b.mu.Lock()
	b.count++
	if b.count == b.n {
		// Último en llegar: cierra molinete2 y abre molinete1
		<-b.molinete2
		b.molinete1 <- struct{}{}
	}
	b.mu.Unlock()

	// Molinete 1: pass-through (uno a uno hasta que todos lleguen)
	token := <-b.molinete1
	b.molinete1 <- token

	// --- Fase de salida ---
	b.mu.Lock()
	b.count--
	if b.count == 0 {
		// Último en salir: cierra molinete1 y abre molinete2
		<-b.molinete1
		b.molinete2 <- struct{}{}
	}
	b.mu.Unlock()

	// Molinete 2: pass-through (garantiza que todos completen antes de la próxima fase)
	token2 := <-b.molinete2
	b.molinete2 <- token2
}

const (
	N_TRABAJADORES = 5
	N_FASES        = 3
)

func main() {
	barrera := NuevaBarrera(N_TRABAJADORES)
	var wg sync.WaitGroup

	fmt.Printf("=== MapReduce Simplificado: %d trabajadores, %d fases ===\n\n", N_TRABAJADORES, N_FASES)

	for i := 1; i <= N_TRABAJADORES; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for fase := 1; fase <= N_FASES; fase++ {
				// Cálculo independiente (duración variable)
				dur := time.Duration(rand.Intn(400)+100) * time.Millisecond
				time.Sleep(dur)
				fmt.Printf("  Trabajador %d: completó cálculo de Fase %d (%dms)\n",
					id, fase, dur.Milliseconds())

				barrera.Esperar() // sincronización: esperar a todos

				if fase == N_FASES {
					fmt.Printf("  Trabajador %d: [Fase %d] iniciando agregación final.\n", id, fase)
				} else {
					fmt.Printf("  Trabajador %d: todos listos, avanzando a Fase %d.\n", id, fase+1)
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("\n[OK] Todas las fases completadas. Barrera reutilizable funcionando correctamente.")
}
