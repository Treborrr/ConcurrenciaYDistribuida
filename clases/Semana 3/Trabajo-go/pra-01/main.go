package main

import "fmt"

type Empleado struct {
	Nombre  string
	Salario float64
}

func aplicarAumento(e *Empleado, porcentaje float64) {
	if porcentaje > 0 && porcentaje < 30 {
		nuevo_salario := e.Salario * (1 + porcentaje/100)
		e.Salario = nuevo_salario
	} else {
		println("El porcentaje de aumento debe ser entre 0 y 30.")
	}
}

func main() {

	// emp := new(Empleado)
	// emp.Nombre = "Juan Pérez"
	// emp.Salario = 50000.00
	emp := Empleado{"Juan Pérez", 50000.00}

	func(e *Empleado) {
		println("Nombre:", e.Nombre)
		salInicial := e.Salario
		fmt.Printf("Salario Inicial: %f\n", salInicial)
		aplicarAumento(e, 20)
		fmt.Printf("Salario con Aumento: %f\n", e.Salario)
		fmt.Printf("Aumento aplicado: %f\n", e.Salario-salInicial)

	}(&emp)

}
