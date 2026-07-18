package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Mensaje struct {
	Dni    string
	Nombre string
	Edad   int
}

func main() {
	// preparar el arrelgo de mensae
	mensajes := []Mensaje{
		{"12345678", "Juan", 25},
		{"12345678", "Juan", 25},
		{"12345678", "Juan", 25},
	}

	// lamar al server
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}
	// sino
	defer conn.Close()
	// convertir a json el arreglo de mensajes
	jsonData, _ := json.Marshal(mensajes)
	// vamos a serealizar
	jsonString := string(jsonData) + "\n"

	fmt.Println("json: ", jsonString)

	// enviamos al server
	fmt.Fprint(conn, jsonString)

	// imprimimos mensaje por consola de que se envio correctamente
	fmt.Println("mensaje enviado correctamente")
}
