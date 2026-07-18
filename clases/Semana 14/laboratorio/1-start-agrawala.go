package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

var addrs []string
var myNum int

//const myAddr = "localhost:8001"

var myAddr = ""

type Info struct {
	Tipo     string
	NodeNum  int
	NodeAddr string
}

type MyInfo struct {
	contMsg  int
	first    bool
	nextNum  int
	nextAddr string
}

var canStart chan bool
var chMyInfo chan MyInfo

func main() {
	var n int
	fmt.Print("Ingrese port local:")
	myAddr, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	myAddr = fmt.Sprintf("localhost:%s", strings.TrimSpace(myAddr))

	fmt.Printf("Soy el nodo %s\n", myAddr)
	// (1) Solicitamos direcciones de los otros nodos
	fmt.Print("Ingrese cantidad de nodos: ")
	fmt.Scanf("%d\n", &n)

	addrs = make([]string, n)
	for i := range addrs {
		fmt.Printf("Port nodo %d: ", i+1)
		dirnode, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		addrs[i] = fmt.Sprintf("localhost:%s", strings.TrimSpace(dirnode))
		//fmt.Scanf("localhost:%s\n", &(addrs[i]))
	}
	// (1) fin

	// (2) generar numero aleatorio = Ticket

	// (2) fin

	//(3)Iniciar los canales

	// (3) fin

	//(4)Iniciar Goroutine para enviar valor inicial de MyInfo

	// (4) fin

	// (5) starting things off
	go func() {
		fmt.Println("Press start to begin...")
		r := bufio.NewReader(os.Stdin)
		r.ReadString('\n')
		info := Info{"SENDNUM", myNum, myAddr} //enviar número
		for _, addr := range addrs {           //notificar a todos los nodos
			go send(addr, info)
		}
	}()
	// (5) fin

	server() //modo escucha de servicio de procesamiento de mensaje Info
}
func server() {
	//host := fmt.Sprintf("%s:8001", myAddr)
	host := myAddr
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	msg, _ := r.ReadString('\n')
	var info Info
	json.Unmarshal([]byte(msg), &info)
	fmt.Println(info)

	//Inicio: Lógica

	//Fin: Lógica

}
func criticalSection() {
	fmt.Println("Nos toca iniciar!")
	//Inicio: Recibiendo MyInfo y determinar el envío del mensaje

	//Fin

}
func send(addr string, msg Info) {
	//remote := fmt.Sprintf("%s:8001", addr)
	remote := addr
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	bMsg, _ := json.Marshal(msg)
	fmt.Fprintln(conn, string(bMsg))
}
