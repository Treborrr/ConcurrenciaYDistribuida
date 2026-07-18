package main

import (
	"fmt"
	"sync"
)

func philosopher(name string, right, left sync.Mutex, s chan bool) {
	for {
		fmt.Println(name, "Pensando!")
		<-s
		left.Lock()
		right.Lock()
		fmt.Println(name, "Comiendo!")
		right.Unlock()
		left.Unlock()
		s <- true
	}
}

func main() {
	fork := make([]sync.Mutex, 5)
	semaphore := make(chan bool, 4)

	semaphore <- true
	go philosopher("Socrates", fork[0], fork[1], semaphore)
	semaphore <- true
	go philosopher(" Aristoteles", fork[1], fork[2], semaphore)
	semaphore <- true
	go philosopher("  Nietzsche", fork[2], fork[3], semaphore)
	semaphore <- true
	go philosopher("   Platon", fork[3], fork[4], semaphore)
	semaphore <- true
	philosopher("    Fonsi", fork[4], fork[0], semaphore)

	close(semaphore)
}
