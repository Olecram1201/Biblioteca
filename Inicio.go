//@autor: Israael Espinoza
//@version: 1.0
//@fecha: 13/05/2024
//@descripcion: Esta aplicacion tiene el objetivo de gestionar un sistema de libros electrónicos.

package main

import (
	//Importamos el paquene que nos ayudara a conectar con la BDD
	"database/sql"
	"fmt"

	//Importamos el driver de postgres
	_ "github.com/lib/pq"
)

// Main
func main() {
	conexionBdd()
	a := 3
	fmt.Println("a", a)
	incrementoA(&a)
	fmt.Println("a2", a)
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

	fmt.Print("Conexion realizada")
}


//Funcion de prueba *Borrar*
func incrementoA(a *int) {
	*a++
}
