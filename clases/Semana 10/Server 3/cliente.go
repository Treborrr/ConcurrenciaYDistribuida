package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const servidorAddr = "localhost:9001"

func main() {
	conn, err := net.Dial("tcp", servidorAddr)
	if err != nil {
		fmt.Println("Error al conectar:", err)
		return
	}
	defer conn.Close()

	//leer stdin y enviar al servidor
	fmt.Print("Escribe un mensaje: ")
	msg, err1 := bufio.NewReader(os.Stdin).ReadString('\n')
	if err1 != nil {
		panic(err1)
	}

	fmt.Fprintln(conn, msg) // enviamos al server el msg leido por la consola
	//recibir respuesta del server
	respuesta, err2 := bufio.NewReader(conn).ReadString('\n')
	if err2 != nil {
		panic(err2)
	}
	fmt.Print("Respuesta recibida del servidor: ", respuesta)

}
