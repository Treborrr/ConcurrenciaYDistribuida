package main

import (
	"fmt"
	"sync"
)

const NUM_HILOS = 100

// sinMutex muestra la race condition: el resultado final puede ser < 100
func sinMutex() int {
	saldo := 0
	var wg sync.WaitGroup

	for i := 0; i < NUM_HILOS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			saldo++ // NO es atómico: read-increment-write puede solaparse
		}()
	}

	wg.Wait()
	return saldo
}

// conMutex protege la sección crítica: garantiza saldo == 100
func conMutex() int {
	saldo := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < NUM_HILOS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			saldo++ // sección crítica protegida
			mu.Unlock()
		}()
	}

	wg.Wait()
	return saldo
}

func main() {
	fmt.Println("=== Exclusión Mutua - Saldo Bancario ===")
	fmt.Printf("Hilos concurrentes: %d | Incremento por hilo: 1 | Esperado: %d\n\n", NUM_HILOS, NUM_HILOS)

	// Ejecutar varias veces sin mutex para mostrar race condition
	fmt.Println("--- Sin Mutex (puede perder actualizaciones) ---")
	for i := 0; i < 5; i++ {
		resultado := sinMutex()
		status := "OK"
		if resultado != NUM_HILOS {
			status = "RACE CONDITION detectada!"
		}
		fmt.Printf("  Intento %d: saldo = %d  → %s\n", i+1, resultado, status)
	}

	fmt.Println("\n--- Con Mutex (siempre correcto) ---")
	for i := 0; i < 5; i++ {
		resultado := conMutex()
		fmt.Printf("  Intento %d: saldo = %d  → OK\n", i+1, resultado)
	}

	fmt.Printf("\nEjecuta con: go run -race main.go  para detectar la race condition explícitamente.\n")
}
