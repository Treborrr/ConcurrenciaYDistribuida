package main

import (
	"bufio"
	"fmt"
	"net"
)

func handdleConnection(conn net.Conn) {
	defer conn.Close()
	// recibir el mensaje de clientes
	msg, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		panic(err)
	}
	// enviar mensaje de confirmacion
	fmt.Print("mensaje recibido: ", msg)

	//enviar mensaje de respuesta al cliente apovechando la omunicacion q uan no se a cerrado
	fmt.Fprintln(conn, "Recibido...")

}

func main() {
	ln, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		fmt.Println("error ", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handdleConnection(conn) // soporta concurrencia de conexiones
	}

}
