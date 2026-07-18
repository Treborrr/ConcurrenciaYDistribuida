package main

import (
	"fmt"
	"net"
)

func main() {
	enviar(9)
	enviar(12)
	enviar(10)
	enviar(2)
}
func enviar(numero int) {
	//lógica de enviar al sgte nodo
	conn, _ := net.Dial("tcp", "localhost:9080")
	defer conn.Close()
	//enviamos
	fmt.Fprintln(conn, numero)
}
