package main

import "fmt"

type Persona struct {
	_nombre   string
	_edad     int
	_telefono int
}

func (p *Persona) setNombre(value string) {
	p._nombre = value
}

func (p Persona) getNombre() string {
	return p._nombre
}

func (p *Persona) setEdad(value int) {
	p._edad = value
}

func (p Persona) getEdad() int {
	return p._edad
}

func (p *Persona) setTelefono(value int) {
	p._telefono = value
}

func (p Persona) getTelefono() int {
	return p._telefono
}

func main() {
	p := Persona{}
	p.setNombre("David")
	p.setEdad(25)
	p.setTelefono(123456789)
	fmt.Println("Nombre de la persona =", p.getNombre(), "|", "Edad =", p.getEdad(), "|", "Numero de telefono =", p.getTelefono())
}
