// topología anillo o circular //TCP
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// variables globales
var dirLocal string ///IP:PORT
var dirRemoto string
var num int

// rol servidor / Cliente
func main() {
	//ingresar dirección local
	fmt.Print("Ingrese el puerto local:")
	portLocal, err := bufio.NewReader(os.Stdin).ReadString('\n')
	//manejo de errores

	errorLectura(err)

	portLocal = strings.TrimSpace(portLocal)
	dirLocal = fmt.Sprintf("localhost:%s", portLocal)
	//////////////////////////////////////////////////
	//ingresar dirección remota
	fmt.Print("Ingrese el puerto remoto:")
	portRemoto, err1 := bufio.NewReader(os.Stdin).ReadString('\n')

	errorLectura(err1)

	portRemoto = strings.TrimSpace(portRemoto)
	dirRemoto = fmt.Sprintf("localhost:%s", portRemoto)
	////////////////////////////////////////////////
	fmt.Println(dirLocal)
	fmt.Println(dirRemoto)
	///////////////////////////////////////////////
	//Rol de servidor
	//////////////////////////////////////////////
	//exponemos un puerto -> puerto local
	ls, err2 := net.Listen("tcp", dirLocal)
	errorLectura(err2)
	defer ls.Close()
	//escucha constante
	for {
		conn, err3 := ls.Accept()
		errorLectura(err3)
		go handle(conn)
	}
}

// Funciones propias
func errorLectura(err error) {
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(-1) //finalizar el programa
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	///////////////////////
	//lógica del servicio
	//////////////////////
	strNum, _ := bufio.NewReader(conn).ReadString('\n')
	strNum = strings.TrimSpace(strNum)
	num, _ = strconv.Atoi(strNum)
	fmt.Printf("Llegó el número %d\n", num)
	//comparación
	if num == 0 {
		fmt.Printf("Boommmm!\n")
	} else {
		//enviar el nro-1 al nodo remoto
		enviar(num - 1)
	}
}

func enviar(num int) {
	conn, _ := net.Dial("tcp", dirRemoto)
	defer conn.Close()
	//envío
	fmt.Fprintln(conn, num)
}
