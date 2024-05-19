//@autor: Israael Espinoza
//@version: 1.0
//@fecha: 13/05/2024
//@descripcion: Esta aplicacion tiene el objetivo de gestionar un sistema de libros electrónicos.

package main

import (
	"bufio"
	"os"

	//Importamos el paquene que nos ayudara a conectar con la BDD
	"database/sql"
	"fmt"
	"strings"

	//Importamos el driver de MSQL server
	_ "github.com/lib/pq"
)

// Main
func main() {
	conexionBdd()
	menuInicio()

}

// Funcion que me permite navegar en un menu
func menuInicio() {
	var opcion int

	for opcion != 4 {
		fmt.Println("MENU")
		fmt.Println("1. Ingresar un libro")
		fmt.Println("4. Salir")
		fmt.Println("Que accion desea hacer: ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			ingresoLibro()
		case 4:
			os.Exit(0)
		}

	}
}

// Funcion que nos da mla conexion con la BDD
func conexionBdd() {
	//conexion con la BDD
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "mibe2001"
		dbname   = "Biblioteca"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//Establecemos la conexion
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion a la BASE DE DATOS realizada con EXITO")
}

// Funcion para ingresar un libro
// En este caso no se solicita un ID puesto que la BDD
// se encargara de crear uno mediante parametros de identidad.
func ingresoLibro() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el nombre del libro: ")
	titulo, _ := reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	fmt.Println("Ingrese la tematica del libro: ")
	tema, _ := reader.ReadString('\n')
	tema = strings.TrimSpace(tema)

	fmt.Println("Ingrese el autor del libro: ")
	autor, _ := reader.ReadString('\n')
	autor = strings.TrimSpace(autor)

	fmt.Println("Ingrese el precio del libro: ")
	var precio int
	fmt.Scanln(&precio)

	nuevoLibro := libro{
		titulo:   titulo,
		tematica: tema,
		autor:    autor,
		precio:   precio,
	}

	fmt.Println("El libro ingresado presenta las siguientes caracteristicas:\n", nuevoLibro)
	var listaLibros []libro
	listaLibros = append(listaLibros, nuevoLibro)

	fmt.Println("Listado de Libros: ")
	for _, libro := range listaLibros {
		fmt.Println(libro)
	}
}

// Zona de creacion de clases
type libro struct {
	titulo   string
	tematica string
	autor    string
	precio   int
	libroID  int
}
