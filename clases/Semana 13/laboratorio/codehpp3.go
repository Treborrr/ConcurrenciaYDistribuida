// algoritmo P2P
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

var addrs []string //bitácora de direcciones de los nodos de la red
var hostIP string  //dirección del nodo actual

// servicios
const (
	registroPort  = 9000
	notificaPort  = 9001
	hotpotatoPort = 9002
)

func main() {
	//obtener la IP del nodo
	hostIP = devolverIP()
	fmt.Printf("Mi IP es %s\n", hostIP)
	//servicio de registro de nodos: modo Server
	go registrarService() //port 9000
	go hotpotatoService() //port 9002

	//solicitar unirse a la red: modo Cliente
	fmt.Print("Ingrese la Ip de nodo remoto a unirse: ")
	ipRemoto, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ipRemoto = strings.TrimSpace(ipRemoto)
	//condición cuando es el primer nodo de la red
	if ipRemoto != "" {
		//registrar el nodo
		registrarCliente(ipRemoto)
	}

	notificarService() //port 9001
}

func devolverIP() string {
	var ipresp string = "127.0.0.1"
	interfaces, _ := net.Interfaces()
	for _, valInterface := range interfaces {
		//fmt.Println(valInterface.Name)
		if valInterface.Name == "Ethernet" {
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
func registrarService() { //rol servidor
	//modo escucha del servicio para registrar un nuevo nodo a la red
	hostdir := fmt.Sprintf("%s:%d", hostIP, registroPort)
	ln, _ := net.Listen("tcp", hostdir) //escucha
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handlerRegistro(con)
	}
}
func handlerRegistro(con net.Conn) {
	defer con.Close()
	//capturamos la IP del nodo cliente
	ip, _ := bufio.NewReader(con).ReadString('\n')
	//limpiar dato ip de espacios en los extremos
	ip = strings.TrimSpace(ip)

	//1.- Devolver al nodo cliente la bitácora de direcciones
	jsonBytes, _ := json.Marshal(addrs) //formateo json
	//serializar
	fmt.Fprintln(con, string(jsonBytes))

	//2.- Comunicar al resto de nodo la incorporación de un nuevo nodo
	comunicarTodos(ip)

	//3.- Actualiza la bitácora
	addrs = append(addrs, ip)

	//4.- Imprimir bitácora
	fmt.Println(addrs)
}
func comunicarTodos(ip string) {
	for _, valIp := range addrs {
		enviar(valIp, ip)
	}
}
func enviar(valIp, ip string) {
	remoteHost := fmt.Sprintf("%s:%d", valIp, notificaPort)
	con, _ := net.Dial("tcp", remoteHost) //llamada al nodo remoto
	defer con.Close()
	fmt.Fprintf(con, "%s\n", ip) //enviar la IP de nuevo nodo al resto de nodos de la red
}
func registrarCliente(ipRemoto string) {
	//Realizar llamado a host remoto
	dirRemota := fmt.Sprintf("%s:%d", ipRemoto, registroPort) //IP:PORT
	//1.- Enviar su IP al nodo remoto
	conn, _ := net.Dial("tcp", dirRemota)
	defer conn.Close()
	fmt.Fprintln(conn, hostIP) //el nodo cliente envía su IP al nodo remoto

	//2.- Recepciona bitácora como respuesta de la llamada a host remoto
	strBitacora, _ := bufio.NewReader(conn).ReadString('\n')
	strBitacora = strings.TrimSpace(strBitacora)
	var tempAddr []string
	json.Unmarshal([]byte(strBitacora), &tempAddr)

	//3.- actualizar la bitácora del nodo actual
	//agregando el nodo remoto
	addrs = append(tempAddr, ipRemoto) //actualiza la bitácora

	//4.- imprimir bitácora
	fmt.Println(addrs)
}
func notificarService() {
	//modo escucha del servicio de notificación de la llegada del nuevo nodo
	hostLocal := fmt.Sprintf("%s:%d", hostIP, notificaPort)
	ln, _ := net.Listen("tcp", hostLocal) //esucha por el puerto 9001
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handlerNotificador(con)
	}
}
func handlerNotificador(con net.Conn) {
	defer con.Close()
	//se recibe la IP del nuevo nodo q ingresa a la red
	ipCliente, _ := bufio.NewReader(con).ReadString('\n')
	ipCliente = strings.TrimSpace(ipCliente) //limpiar espacios en los extremos
	//se adiciona la IP a la bitácora de este nodo
	addrs = append(addrs, ipCliente)
	//se imprime la bitácora
	fmt.Println(addrs)
}

func hotpotatoService() {
	dirLocalHP := fmt.Sprintf("%s:%d", hostIP, hotpotatoPort)
	//listen
	ls, _ := net.Listen("tcp", dirLocalHP)
	defer ls.Close()
	for {
		conn, _ := ls.Accept()
		go handleHP(conn)
	}
}
func handleHP(conn net.Conn) {
	defer conn.Close()
	//lógica
	strNum, _ := bufio.NewReader(conn).ReadString('\n')
	strNum = strings.TrimSpace(strNum)
	num, _ := strconv.Atoi(strNum) //convierte a número
	fmt.Printf("Llegó el número %d\n", num)
	//evaluar
	if num == 0 {
		fmt.Println("Booommmmm!")
	} else {
		enviarHP(num - 1)
	}
}
func enviarHP(num int) {
	//seleccionar una IP al azar
	indice := rand.Intn(len(addrs))
	dirRemoto := fmt.Sprintf("%s:%d", addrs[indice], hotpotatoPort)
	fmt.Printf("El número %d enviado al nodo %s\n", num, addrs[indice])
	conn, _ := net.Dial("tcp", dirRemoto)
	fmt.Fprintln(conn, num)

}
