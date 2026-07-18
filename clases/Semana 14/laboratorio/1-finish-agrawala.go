package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

var addrs []string //libreta de direcciones
var myNum int      //#ticket

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
	fmt.Print("Ingrese cantidad de nodos de la red: ")
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
	rand.Seed(time.Now().UTC().UnixNano()) //semilla
	myNum = rand.Intn(1000000)
	fmt.Printf("Mi ticket es %d\n", myNum)
	// (2) fin

	//(3)Iniciar los canales
	canStart = make(chan bool)
	chMyInfo = make(chan MyInfo)
	// (3) fin

	//(4)Iniciar Goroutine para enviar valor inicial de MyInfo
	go func() {
		chMyInfo <- MyInfo{0, true, 1000001, ""}
	}()
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
	switch info.Tipo {
	case "SENDNUM":
		myInfo := <-chMyInfo //leer la información de análisis de cada nodo
		if info.NodeNum < myNum {
			myInfo.first = false
		} else if info.NodeNum < myInfo.nextNum {
			//actualizar la información del próximo ticket a ejecutar su SC
			myInfo.nextNum = info.NodeNum
			myInfo.nextAddr = info.NodeAddr
		}
		myInfo.contMsg++
		//guardar le información de myInfo enviando al canal
		go func() {
			chMyInfo <- myInfo
		}()
		//evaluar si el nodo recibió todos los mensajes de los demás nodos
		if myInfo.contMsg == len(addrs) {
			//determinar si inicia su atención en la SC
			if myInfo.first {
				criticalSection() //si es el primero del cluster (ticket menor) ingresa a la SC
			} else {
				canStart <- true //manda el mensaje para que espere la señal de turno
			}
		}

	case "START":
		//descarta el mensaje del canal para su inicio de la atención en SC
		<-canStart
		criticalSection()
	}
	//Fin: Lógica

}
func criticalSection() {
	fmt.Println("Nos toca iniciar!")
	fmt.Println("Iniciando ejecución de SC!") //ejecuta SC y
	//Inicio: Recibiendo MyInfo y determinar el envío del mensaje
	myInfo := <-chMyInfo //recuperamos la información guardada del nodo
	if myInfo.nextAddr == "" {
		fmt.Printf("Soy el nodo último a ejecutar!!!")
	} else {
		//envía mensaje al siguiente nodo que toca su ejecución de SC
		fmt.Printf("Siguiente nodo es %s con #Ticket %d\n", myInfo.nextAddr, myInfo.nextNum)
		msg := Info{Tipo: "START"}
		send(myInfo.nextAddr, msg)
	}
	//Fin

}
func send(addr string, msg Info) {
	//remote := fmt.Sprintf("%s:8001", addr)
	remote := addr
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	bMsg, _ := json.Marshal(msg)     //codificar el mensaje
	fmt.Fprintln(conn, string(bMsg)) //serializando el mensaje antes de enviar
}
