package main

import (
	"fmt"
	"sync"
)

var horasp = 5
var horasb = 5

var wg sync.WaitGroup

func estudiante(m *sync.Mutex) {
	for {
		m.Lock()
		if horasp > 0 {
			horasp--
			fmt.Printf("Estudiante en sala, Hora %d\n", horasp)
		} else {
			horasb = 5
			defer wg.Done()
		}

		m.Unlock()
	}
}

func becario(m *sync.Mutex) {
	for {
		m.Lock()
		if horasb > 0 && horasp == 0 {
			horasb--
			fmt.Printf("Becario en sala, Hora %d\n", horasb)
		} else {
			horasp = 5
			defer wg.Done()
		}
		m.Unlock()
	}
}

func main() {
	m := new(sync.Mutex)
	wg.Add(2)
	go estudiante(m) //proceso 1
	go becario(m)    //proceso 2
	wg.Wait()
	//var in string
	//fmt.Scanln(&in)
}
