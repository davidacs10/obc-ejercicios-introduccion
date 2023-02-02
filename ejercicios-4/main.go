package main

import "fmt"

type Persona struct {
	nombre   string
	edad     int
	telefono int
}

type Cliente struct {
	Persona
	credito float32
}

type Trabajador struct {
	Persona
	Salario float32
}

func main() {
	c := Cliente{Persona{"David", 25, 123456789}, 432.140}
	fmt.Println(c)

	t := Trabajador{Persona{"David", 25, 123456789}, 8000.00}
	fmt.Println(t)
}
