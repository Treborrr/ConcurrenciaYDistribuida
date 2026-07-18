// problema de conway
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX = 9
	K   = 4
)

func compress(inC, pipe chan rune) {
	n := 0
	previous := <-inC
	for {
		c := <-inC
		if (c == previous) && (n < MAX-1) {
			n++
		} else {
			if n > 0 {
				pipe <- rune(n + 49) //envío del número de repeticiones como un carácter
				n = 0
			}
			pipe <- previous
			previous = c
		}
	}
}

func output(pipe, outC chan rune) {
	m := 0
	for {
		outC <- <-pipe
		m++
		if m >= K {
			outC <- '\n'
			m = 0
		}
	}
}

func main() {
	inC := make(chan rune)
	pipe := make(chan rune)
	outC := make(chan rune)
	go compress(inC, pipe)
	go output(pipe, outC)
	rand.Seed(time.Now().UTC().UnixNano())

	//lanzador
	go func() {
		for {
			inC <- rune(rand.Intn(26) + 65) //envío de caracteres aleatorios (A-Z)
		}
	}()

	for {
		fmt.Printf("%c", <-outC) //imprime los caracteres comprimidos
	}
}
