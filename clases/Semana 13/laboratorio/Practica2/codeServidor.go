package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var addrs []string                //Bitácora de direcciones
var dirLocal string               //ubicación del servidor
const portRegistro = "9000"       //es para el registro de clientes
const portProcesoNumeros = "9001" //es para el proceso de números

func main() {
	dirLocal = fmt.Sprintf("localhost:%s", portRegistro)
	//iniciar el modo de servidor
	ln, err := net.Listen("tcp", dirLocal)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Servidor iniciado en", dirLocal)

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
	// Aquí puedes manejar la conexión con el cliente
	addr, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer del cliente:", err)
		return
	}
	addr = strings.TrimSpace(addr) // Eliminar espacios en blanco y saltos de línea
	fmt.Println("Mensaje recibido:", addr)
	//agregar a la bitácora de direcciones
	addrs = append(addrs, addr)
	fmt.Println(addrs)
	//respuesta al cliente
	fmt.Fprintf(conn, "Cliente conectado al servidor %s\n", dirLocal)
}
