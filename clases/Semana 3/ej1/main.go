package main

import (
	"fmt"
	"sync"
)

func main() {
	balance := 100
	var wg sync.WaitGroup

	// Cada goroutine hará varias operaciones para que sea más visible la carrera.
	const ops = 1000

	wg.Add(2)

	var mu sync.Mutex

	// Stingy: suma 10 en cada iteración.
	go func() {
		defer wg.Done()
		for i := 0; i < ops; i++ {
			mu.Lock()
			balance += 10
			mu.Unlock()
		}
	}()

	// Spendy: resta 10 en cada iteración.
	go func() {
		defer wg.Done()
		for i := 0; i < ops; i++ {
			mu.Lock()
			balance -= 10
			mu.Unlock()
		}
	}()

	wg.Wait()

	// Con 100 inicial y ops iguales, el resultado "matemático" sería 100.
	// Sin mutex/atomic, el valor suele ser incorrecto (data race).
	fmt.Println("balance final:", balance)
}
