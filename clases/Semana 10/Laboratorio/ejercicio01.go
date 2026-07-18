// buffer de los canales
package main

import "fmt"

func main() {
	//canales asincronos
	messages := make(chan string, 2)
	messages <- "mensaje 1"
	messages <- "mensaje 2"

	fmt.Println(<-messages) //recepción de datos a traves del canal messages
	fmt.Println(<-messages)
}
