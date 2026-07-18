package main

import (
	"fmt"
	"sync"
)

func producer(id int, ch chan string) {
	c := 0
	for {
		c++
		ch <- fmt.Sprintf("Producto %d producido por productor %d", c, id)
	}
}

func consumer(id int, ch chan string) {
	for {
		fmt.Printf("Consumidor %d consumiendo %s\n", id, <-ch)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(4)
	for i := 0; i < 4; i++ {
		go producer(i, ch)
		go consumer(i, ch)
	}

	wg.Wait()

}
