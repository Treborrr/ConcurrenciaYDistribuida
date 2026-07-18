package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Mensaje struct {
	Dni    string
	Nombre string
	Edad   int
}

func handleConnection(conn net.Conn) {
	// logica de recuperar mensaje, almacenarlo en memora e imprimirlo ppr consola

	defer conn.Close()
	// recepcoin dle mensaje
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	//caso contrario decarar una variable tipro struct
	var mensajeStruct []Mensaje
	//decodificar el mensaje
	json.Unmarshal([]byte(msg), &mensajeStruct) //decodificando para guardarlo en la variable
	//imprimir el arreglo de mensajes
	fmt.Println("msg: ", mensajeStruct)

}

func main() {

	//crear el servicio servidor
	ln, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
