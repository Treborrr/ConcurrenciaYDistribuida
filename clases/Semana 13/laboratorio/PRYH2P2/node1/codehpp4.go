// algoritmo P2P
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

var addrs []string //bitácora de direcciones de los nodos de la red
var hostIP string  //dirección del nodo actual

// servicios
const (
	hotpotatoPort = 9002
)

func main() {
	//obtener la IP del nodo
	hostIP = devolverIP() //"172.20.0.2"
	fmt.Printf("Mi IP es %s\n", hostIP)
	//lista de direcciones
	addrs = []string{"172.20.0.3", "172.20.0.4", "172.20.0.5"}

	//servicio de registro de nodos: modo Server
	hotpotatoservice()

}

func devolverIP() string {
	var ipresp string = "127.0.0.1"
	interfaces, _ := net.Interfaces()
	for _, valInterface := range interfaces {
		//fmt.Println(valInterface.Name)
		if valInterface.Name == "eth0" { //cuando el SO es linux
			dir, _ := valInterface.Addrs()
			for _, valDireccion := range dir {
				//fmt.Println(valDireccion.String())
				switch d := valDireccion.(type) {
				case *net.IPNet:
					if d.IP.To4() != nil {
						ipresp = d.IP.To4().String()
						//fmt.Println(d.IP.To4().String())
						//fmt.Println(d.IP.String())
						// case *net.IP:
						// 	fmt.Println(d.IP.String())
					}
				}
			}
		}
	}

	return ipresp
}

// ///////////////////////////////////////
// HP
// //////////////////////////////
func hotpotatoservice() {
	//servicio de escucha
	hostpotatoremote := fmt.Sprintf("%s:%d", hostIP, hotpotatoPort)
	ls, _ := net.Listen("tcp", hostpotatoremote)
	defer ls.Close()
	for {
		conn, _ := ls.Accept()
		go handleHotPotato(conn)

	}
}
func handleHotPotato(conn net.Conn) {
	//lógica del servicio
	defer conn.Close()
	strNum, _ := bufio.NewReader(conn).ReadString('\n')
	strNum = strings.TrimSpace(strNum)
	num, _ := strconv.Atoi(strNum)
	fmt.Printf("Recibido el nro = %d\n", num)
	if num == 0 {
		fmt.Println("Boommmmmm!")
	} else {
		enviarHP(num - 1)
	}
}
func enviarHP(num int) {
	//enivío random
	inx := rand.Intn(len(addrs)) //seleccionamos el indice de la IP a enviar
	remoteHostP := fmt.Sprintf("%s:%d", addrs[inx], hotpotatoPort)
	conn, _ := net.Dial("tcp", remoteHostP)
	fmt.Printf("Enviando %d hacia el nodo %s\n", num, addrs[inx])
	defer conn.Close()
	fmt.Fprintln(conn, num)

}
