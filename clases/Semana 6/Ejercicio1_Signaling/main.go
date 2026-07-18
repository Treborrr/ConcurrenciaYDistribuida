package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Semáforo inicializado en 0: bloquea a B hasta que A haga signal
	sem := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	// Proceso A: descarga el paquete de datos
	go func() {
		defer wg.Done()
		fmt.Println("Proceso_A: iniciando descarga del paquete de datos...")
		time.Sleep(600 * time.Millisecond) // simula tiempo de descarga
		fmt.Println("Proceso_A: descarga FINALIZADA.")
		sem <- struct{}{} // signal → desbloquea a Proceso_B
	}()

	// Proceso B: procesa y visualiza los datos (debe esperar a A)
	go func() {
		defer wg.Done()
		fmt.Println("Proceso_B: esperando señal de descarga completa...")
		<-sem // wait → se bloquea hasta recibir signal de A
		fmt.Println("Proceso_B: datos recibidos. Procesando y visualizando...")
		time.Sleep(200 * time.Millisecond) // simula procesamiento
		fmt.Println("Proceso_B: visualización completa.")
	}()

	wg.Wait()
	fmt.Println("\n[OK] Ambos procesos han finalizado. B nunca ejecutó antes que A.")
}
