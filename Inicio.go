package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	conexionBdd()
	a := 3
	fmt.Println("a", a)
	incrementoA(&a)
	fmt.Println("a2",a)
}

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


func incrementoA (a *int) {
	*a++
} 
