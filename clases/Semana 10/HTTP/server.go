package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bienvenido a mi Api")
}

func saludoHandler(w http.ResponseWriter, r *http.Request) {
	nombre := r.URL.Query().Get("nombre")
	if nombre == "" {
		nombre = "desconocido"
	}
	fmt.Fprintf(w, "Hola, %s!", nombre)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	fmt.Println("Escuchando en puerto 8080")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/saludo", saludoHandler)
	http.HandleFunc("/status", statusHandler)

	server := http.Server{Addr: ":8080"}
	server.ListenAndServe()
}
