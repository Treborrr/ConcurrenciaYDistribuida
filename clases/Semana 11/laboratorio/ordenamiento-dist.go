// ordenamiento distribuido
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// variables
var dirLocal string
var dirRemoto string
var chCont chan int
var num int //tope de valores a recibir
var min int //valor mínimo almacenado por nodo

func main() {
	//ingresar valores por consola de port local
	fmt.Print("Ingrese puerto local:")
	br := bufio.NewReader(os.Stdin)
	portLocal, _ := br.ReadString('\n')
	portLocal = strings.TrimSpace(portLocal)          //eliminar espacios en los bordes
	dirLocal = fmt.Sprintf("localhost:%s", portLocal) // ejemplo localhost:9080

	//ingresa valores de port remoto
	fmt.Print("Ingrese puerto remoto:")
	portRemoto, _ := br.ReadString('\n')
	portRemoto = strings.TrimSpace(portRemoto)
	dirRemoto = fmt.Sprintf("localhost:%s", portRemoto)

	//# máximo de valores a recibir
	fmt.Print("Ingrese el # máximo de valores a recibir:")
	strNum, _ := br.ReadString('\n')
	strNum = strings.TrimSpace(strNum)
	num, _ = strconv.Atoi(strNum)

	//inicializar el contador de # que se van llegando
	//sincronizar
	chCont = make(chan int, 1)
	//dato inicial
	chCont <- 0

	//rol servidor
	ls, _ := net.Listen("tcp", dirLocal)
	defer ls.Close()
	//escucha constante
	for {
		conn, _ := ls.Accept()
		//manejar millones de conexiones
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	//lógica del servicio
	//recuperar el número
	br := bufio.NewReader(conn)
	dato, _ := br.ReadString('\n')
	dato = strings.TrimSpace(dato)
	dNum, _ := strconv.Atoi(dato) //convertimos en valor numérico

	fmt.Printf("Número recibido = %d\n", dNum)

	//lógica de menor
	//sincronizar el número de recibidos
	cont := <-chCont

	if cont == 0 {
		//es el primero
		min = dNum
	} else if dNum < min {
		enviar(min) //enviamos el número al siguiente nodo remoto
		min = dNum  //ya tenemos un nuevo valor mínimo
	} else {
		enviar(dNum) //enviamos porque el dnum es mayor al min
	}

	cont++           //actualizamos el # recibidos
	if cont == num { //si llegamos al máximo de val recibidos
		fmt.Printf("# = %d\n", min)
		cont = 0
	}

	//sincronizamos con el # valores recibidos
	chCont <- cont
}

func enviar(numero int) {
	//lógica de enviar al sgte nodo
	conn, _ := net.Dial("tcp", dirRemoto)
	defer conn.Close()
	//enviamos
	fmt.Fprintln(conn, numero)
}
