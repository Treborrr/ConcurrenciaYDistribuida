package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// const localAddr = "localhost:8000" // su propia IP aquí
var localAddr string

const (
	cnum = iota // iota genera valores en secuencia y se reinicia en cada bloque const
	opa         //voto 1
	opb         //voto 2
)

type tmsg struct {
	Code int
	Addr string
	Op   int
}

// Las IP de los demás participantes acá, todos deberían usar el puerto 8000
var addrs []string

var chInfo chan map[string]int //almacenamiento de las votaciones de los nodos
var chOpc chan int

func main() {
	var opc int
	var strOpc string
	var n int //nro de nodos

	fmt.Print("Ingrese port local:")
	localAddr, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	localAddr = fmt.Sprintf("localhost:%s", strings.TrimSpace(localAddr))

	fmt.Printf("Soy el nodo %s\n", localAddr)

	// (1) Solicitamos direcciones de los otros nodos
	fmt.Print("Ingrese cantidad de nodos de la red: ")
	fmt.Scanf("%d\n", &n)

	addrs = make([]string, n)
	for i := range addrs {
		fmt.Printf("Port nodo %d: ", i+1)
		dirnode, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		addrs[i] = fmt.Sprintf("localhost:%s", strings.TrimSpace(dirnode))
		//fmt.Scanf("localhost:%s\n", &(addrs[i]))
	}

	fmt.Println(addrs)

	chInfo = make(chan map[string]int)
	chOpc = make(chan int)

	go func() { chInfo <- map[string]int{} }() //mensaje inicio
	go server()                                //servicio de escucha de recepción de mensajes
	time.Sleep(time.Millisecond * 100)

	for {
		fmt.Print("Your option (1 o 2): ")
		//fmt.Scanf("%d\n", &opc)
		strOpc, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		strOpc = strings.TrimSpace(strOpc)

		opc, _ = strconv.Atoi(strOpc)

		//fmt.Println(opc)

		msg := tmsg{cnum, localAddr, opc} //armado de mensaje

		//fmt.Println(msg)

		for _, addr := range addrs { //notificar a todos los nodos el mensaje armado
			//fmt.Println(addr)
			send(addr, msg)
		}
	}
}
func server() {
	if ln, err := net.Listen("tcp", localAddr); err != nil {
		log.Panicln("Can't start listener on", localAddr)
	} else {
		defer ln.Close()
		fmt.Println("Listeing on", localAddr)
		for {
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				go handle(conn)
			}
		}
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var msg tmsg
	if err := dec.Decode(&msg); err != nil {
		log.Println("Can't decode from", conn.RemoteAddr())
	} else {
		fmt.Println(msg)
		switch msg.Code {
		case cnum:
			concensus(conn, msg) //aplicando el consenso
		}
	}
}
func concensus(conn net.Conn, msg tmsg) {
	//Inicio: Aplicando lógica de consenso
	info := <-chInfo        //recibir los datos q va almacenando el nodo
	info[msg.Addr] = msg.Op //por cada dirección de nodo guarda su opinión
	//evalua el consenso: cuando el nodo recibe los mensajes de los nodos de la red
	if len(info) == len(addrs) {

		//asumiendo que la cantidad de votantes son impares
		ca, cb := 0, 0
		//contabilizar los votos
		for _, op := range info {
			//lógica
			if op == opa {
				ca++
			} else {
				cb++
			}
		}
		//evaluar consenso por mayoría
		if ca > cb {
			fmt.Println("Vamos por a!!")
		} else {
			fmt.Println("Vamos por b!!")
		}
		//inicializar el arreglo de votos
		info = map[string]int{}

	}
	go func() {
		chInfo <- info
	}()
	//Fin
}
func send(remoteAddr string, msg tmsg) {
	if conn, err := net.Dial("tcp", remoteAddr); err != nil {
		log.Println("Can't dial", remoteAddr)
	} else {
		defer conn.Close()
		fmt.Println("Sending to", remoteAddr)
		enc := json.NewEncoder(conn)
		enc.Encode(msg)
	}
}
