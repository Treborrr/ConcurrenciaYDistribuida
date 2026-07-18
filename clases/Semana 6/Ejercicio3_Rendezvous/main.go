package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Dos semáforos (canales con buffer 1) inicializados en 0
	jugador1Llego := make(chan struct{}, 1)
	jugador2Llego := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	// Jugador 1
	go func() {
		defer wg.Done()

		// Fase 1: preparación (duración aleatoria para demostrar sincronización)
		fmt.Println("Jugador 1: [Fase 1] preparando equipo...")
		time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)
		fmt.Println("Jugador 1: llegó al punto de control.")

		jugador1Llego <- struct{}{} // signal: aviso que llegué
		<-jugador2Llego             // wait:   espero a Jugador 2

		// Fase 2: solo se ejecuta cuando AMBOS llegaron
		fmt.Println("Jugador 1: [Fase 2] ¡ambos listos! iniciando misión.")
	}()

	// Jugador 2
	go func() {
		defer wg.Done()

		// Fase 1: preparación
		fmt.Println("Jugador 2: [Fase 1] preparando equipo...")
		time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)
		fmt.Println("Jugador 2: llegó al punto de control.")

		jugador2Llego <- struct{}{} // signal: aviso que llegué
		<-jugador1Llego             // wait:   espero a Jugador 1

		// Fase 2: solo se ejecuta cuando AMBOS llegaron
		fmt.Println("Jugador 2: [Fase 2] ¡ambos listos! iniciando misión.")
	}()

	wg.Wait()
	fmt.Println("\n[OK] Misión iniciada. Ninguno avanzó a Fase 2 sin el otro.")
}
