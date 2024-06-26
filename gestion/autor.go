package gestionBiblioteca

import (
	_ "bufio"
	"database/sql"
	"fmt"
	_ "os"
	_ "strings"
)

//creacion del objeto

type Autor struct {
	nombre   string
	apellido string
	IdA      int
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
func (a *Autor) GetIdAutor() int {
	return a.IdA
}

//hacemos el ingreso del autor a la bdd

func (a *Autor) IngresoAutor(db *sql.DB) {

	// Insertar un nuevo autor
	_, err := db.Exec("INSERT INTO Autor (nombre, apellido) VALUES ($1, $2)",
		a.nombre, a.apellido)
	if err != nil {
		fmt.Println(err)
	}
}

// obtener la ID del autor nueuvo ingresado

func (a *Autor) IDAutor(db *sql.DB) {
	//recuperar el ID del autor

	_, err := db.Exec("SELECT id_autor FROM autor", a.IdA)
	if err != nil {
		fmt.Println(err)
	}
}

//mostrar los libros pertenecientes a cada autor
