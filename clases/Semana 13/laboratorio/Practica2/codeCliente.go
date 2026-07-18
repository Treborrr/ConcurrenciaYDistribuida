package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var dirLocal string //ubicación del cliente

func main() {
	fmt.Print("Ingrese su puerto: ")
	portLocal, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el puerto:", err)
		return
	}
	portLocal = strings.TrimSpace(portLocal) // Eliminar espacios en blanco y saltos de línea
	dirLocal = fmt.Sprintf("localhost:%s", portLocal)
	fmt.Println("Dirección: ", dirLocal)

	//conectar al servidor
	fmt.Print("Ingrese el puerto del servidor: ")
	portServidor, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el puerto del servidor:", err)
		return
	}
	portServidor = strings.TrimSpace(portServidor) // Eliminar espacios en blanco y saltos de línea
	dirServidor := fmt.Sprintf("localhost:%s", portServidor)
	fmt.Println("Conectando al servidor en", dirServidor)

	conn, err := net.Dial("tcp", dirServidor)
	if err != nil {
		fmt.Println("Error al conectar al servidor:", err)
		return
	}
	defer conn.Close()
	//enviar la dirección del cliente al servidor
	fmt.Fprintf(conn, "%s\n", dirLocal)
	//recibir la respuesta del servidor
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error al recibir la respuesta del servidor:", err)
		return
	}
	fmt.Println("Respuesta del servidor:", response)

	///////////////////////
	//Lógica del cliente en su puerto local
	ln, err := net.Listen("tcp", dirLocal)
	if err != nil {
		fmt.Println("Error al iniciar el servidor local:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Servidor local iniciado en", dirLocal)

	//modo escucha constante
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar la conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Aquí puedes agregar la lógica para manejar la conexión entrante
	dato, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer datos de la conexión:", err)
		return
	}
	fmt.Println("Datos recibidos:", dato)
}
