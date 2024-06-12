package gestionBiblioteca

import (
	_ "bufio"
	"database/sql"
	"fmt"
	_ "os"
	_ "strings"
)

// creacion de nuestro objeto Libro
type Libro struct {
	fecha   string
	titulo  string
	archivo string
}

// zona de Sett's
func (l *Libro) SetFecha(fecha string) {
	l.fecha = fecha
}

func (l *Libro) SetTitulo(titulo string) {
	l.titulo = titulo
}

func (l *Libro) SetArchivo(archivo string) {
	l.archivo = archivo
}

// zona de Get's
func (l *Libro) GetFecha() string {
	return l.fecha
}
func (l *Libro) GetTitulo() string {
	return l.titulo
}
func (l *Libro) GetArchivo() string {
	return l.archivo
}

// Constructor que permite el ingreso de un libro nuevo
func (l *Libro) IngresoLibro(db *sql.DB) {

	// Insertar un nuevo libro
	_, err := db.Exec("INSERT INTO Libro (Título, Fecha_Publicación, Archivo) VALUES ($1, $2, $3)",
		l.titulo, l.fecha, l.archivo)
	if err != nil {
		fmt.Println(err)
	}
}
