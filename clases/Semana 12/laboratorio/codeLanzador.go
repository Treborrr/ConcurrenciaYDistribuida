// Centinela
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var dirRemoto string
var num int

func main() {

	//ingresar dirección remota
	fmt.Print("Ingrese el puerto remoto:")
	portRemoto, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	portRemoto = strings.TrimSpace(portRemoto)
	dirRemoto = fmt.Sprintf("localhost:%s", portRemoto)

	fmt.Print("Ingrese el número entero a lanzar:")
	strNum, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strNum = strings.TrimSpace(strNum)
	num, _ = strconv.Atoi(strNum)
	enviar(num)
	fmt.Println(num)
}

func enviar(num int) {
	conn, _ := net.Dial("tcp", dirRemoto)
	defer conn.Close()
	fmt.Fprintln(conn, num)
}
