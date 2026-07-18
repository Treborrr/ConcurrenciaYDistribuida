package main

import (
	"fmt"
	"net"
)

const servidorAddr = "127.0.0.1:8000"

func main() {
	con, err := net.Dial("tcp", servidorAddr)
	if err != nil {
		fmt.Println("Error al conectar:", err)
		return
	}
	defer con.Close()

	fmt.Fprintln(con, "jalados!")
	fmt.Println("Mensaje enviado al servidor.")
}
