// cierre de canales
package main

import "fmt"

func envios(chDatos chan int) {
	for i := 0; i < 10; i++ {
		chDatos <- i //salida
	}

	close(chDatos) //cerrar el canal
}

func main() {
	chDatos := make(chan int) //sincrona

	go envios(chDatos)

	//fmt.Println(<-chDatos)

	for num := range chDatos {
		fmt.Println(num)
	}

}
