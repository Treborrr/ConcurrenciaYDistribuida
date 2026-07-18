package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedData struct {
	readersCount int
	cond         *sync.Cond
	isWriting    bool
}

func (sd *SharedData) startReading(id int) {
	sd.cond.L.Lock()
	for sd.isWriting { // Espera si hay un escritor activo
		sd.cond.Wait()
	}
	sd.readersCount++ // Incrementa el contador de lectores activos
	fmt.Printf("Reader %d starts reading. Readers count: %d\n", id, sd.readersCount)
	sd.cond.L.Unlock()
}

func (sd *SharedData) finishReading(id int) {
	sd.cond.L.Lock()
	sd.readersCount-- // Decrementa el contador de lectores activos
	fmt.Printf("Reader %d finished reading. Remaining readers: %d\n", id, sd.readersCount)
	if sd.readersCount == 0 { // Si no hay más lectores, notifica a los escritores
		sd.cond.Signal()
	}
	sd.cond.L.Unlock()
}

func (sd *SharedData) startWriting(id int) {
	sd.cond.L.Lock()
	for sd.isWriting || sd.readersCount > 0 { // Espera si hay un escritor o lectores activos
		sd.cond.Wait()
	}
	sd.isWriting = true // Indica que hay un escritor activo
	fmt.Printf("Writer %d starts writing.\n", id)
	sd.cond.L.Unlock()
}

func (sd *SharedData) finishWriting(id int) {
	sd.cond.L.Lock()
	sd.isWriting = false // Indica que ya no hay un escritor activo
	fmt.Printf("Writer %d finished writing.\n", id)
	sd.cond.Signal() // Notifica a los lectores o escritores en espera
	sd.cond.L.Unlock()
}

func reader(sd *SharedData, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	sd.startReading(id)
	time.Sleep(200 * time.Millisecond) // Simula tiempo de lectura
	sd.finishReading(id)
}

func writer(sd *SharedData, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	sd.startWriting(id)
	time.Sleep(500 * time.Millisecond) // Simula tiempo de escritura
	sd.finishWriting(id)
}

func main() {
	var wg sync.WaitGroup
	sharedData := &SharedData{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Inicia 3 goroutines de escritores
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go writer(sharedData, i, &wg)
	}

	// Inicia 5 goroutines de lectores
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go reader(sharedData, i, &wg)
	}

	wg.Wait()
	fmt.Println("All readers and writers have finished.")
}
