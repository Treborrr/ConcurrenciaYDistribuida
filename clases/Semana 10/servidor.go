package main

import (
	"bufio"
	"fmt"
	"net"
)

const escucharAddr = "0.0.0.0:8000"

func main() {
	ln, err := net.Listen("tcp", escucharAddr)
	if err != nil {
		fmt.Println("Error al escuchar:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Servidor escuchando en", escucharAddr)

	for {
		con, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		fmt.Println("Cliente conectado:", con.RemoteAddr())
		go func(c net.Conn) {
			defer c.Close()
			r := bufio.NewReader(c)
			msg, _ := r.ReadString('\n')
			fmt.Print("Mensaje recibido: ", msg)
		}(con)
	}
}
