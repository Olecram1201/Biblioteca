package gestionBiblioteca

import (
	_ "bufio"
	_ "os"
	_ "strings"
)

//creacion del objeto

type Autor struct {
	nombre   string
	apellido string
}

// zona de set's
func (a *Autor) SetNombre(nombre string) {
	a.nombre = nombre
}
func (a *Autor) SetApellido(apellido string) {
	a.apellido = apellido
}

//zona de get's

func (a *Autor) GetNombre() string {
	return a.nombre
}
func (a *Autor) GetApellido() string {
	return a.apellido
}


//mostrar los libros pertenecientes a cada autor 
