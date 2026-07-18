// dirección / sentido de los datos
//
//	chPing ----->  proceso
//	      proceso -----> chPongs
//	chPing ----->  proceso  ----> chPongs
package main

import "fmt"

func ping(pings chan string, msg string) {
	//envio
	pings <- msg
}

func pong(pings chan string, pongs chan string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Texto de prueba") //salida
	pong(pings, pongs)             //entrada y salida

	fmt.Println(<-pongs) //salida
}
