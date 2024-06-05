//@autor: Israael Espinoza
//@version: 1.0
//@fecha: 13/05/2024
//@descripcion: Esta aplicacion tiene el objetivo de gestionar un sistema de libros electrónicos.

package main

import (
	_ "github.com/lib/pq"
)

// Main
func main() {
	// db := conexionBdd()
	// defer db.Close()
	// menuInicio(db)
	// p := Person{
	// 	CI:          "123456789",
	// 	FullName:    "John Doe",
	// 	Age:         30,
	// 	Phone:       "555-1234",
	// 	Birthdate:   time.Date(1993, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Ethnicity:   "Caucasian",
	// 	Nationality: "American",
	// 	CivilStatus: "Single",
	// }

	// b := Book{
	// 	ISBN:            "978-3-16-148410-0",
	// 	Name:            "Go Programming",
	// 	Subject:         "Programming",
	// 	Overview:        "A book about Go programming.",
	// 	Publisher:       "Tech Books",
	// 	PublicationDate: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Lang:            "English",
	// }

	// a := Author{
	// 	Name:      "Jane Smith",
	// 	Biography: "Author of several programming books.",
	// 	BirthDate: time.Date(1980, 5, 10, 0, 0, 0, 0, time.UTC),
	// }

	// acc := Account{
	// 	Number: "ACC123456",
	// 	Opened: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	State:  "Active",
	// }

	// bi := BookItem{
	// 	Barcode:       "1234567890123",
	// 	NumberOfPages: 300,
	// 	Format:        "Hardcover",
	// 	DueDate:       time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
	// }

	// _ = p
	// _ = b
	// _ = a
	// _ = acc
	// _ = bi

}

// // Funcion que me permite navegar en un menu
// func menuInicio(db *sql.DB) {
// 	var opcion int

// 	for opcion != 4 {
// 		fmt.Println("MENU")
// 		fmt.Println("1. Ingresar un libro")
// 		fmt.Println("4. Salir")
// 		fmt.Println("Que accion desea hacer: ")
// 		fmt.Scanln(&opcion)
// 		switch opcion {
// 		case 1:
// 			ingresoLibro(db)
// 		case 4:
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Opcion no valida")
// 		}

// 	}
// }

// // Funcion que nos da la conexion con la BDD
// func conexionBdd() *sql.DB {
// 	//conexion con la BDD
// 	const (
// 		host     = "localhost"
// 		port     = 5432
// 		user     = "postgres"
// 		password = "mibe2001"
// 		dbname   = "Biblioteca"
// 	)
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	//Establecemos la conexion
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Error al conectar con la base de datos", err)
// 		os.Exit(1) // el valor de 1 indica q el codigo a generado un error por lo que se usa osExit para una  salida inmediata
// 	}

// 	//Comprobacion de la conexion
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("Error al conectar con la base de datos", err)
// 		os.Exit(1) // el valor de 1 indica q el codigo a generado un error por lo que se usa os.Exit para una  salida inmediata
// 	}

// 	fmt.Println("Conexion a la BASE DE DATOS realizada con EXITO")
// 	return db
// }

// // Funcion para ingresar un libro
// // En este caso no se solicita un ID puesto que la BDD
// // se encargara de crear uno mediante parametros de identidad.
// func ingresoLibro(db *sql.DB) {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Título del libro: ")
// 	titulo, _ := reader.ReadString('\n')
// 	titulo = strings.TrimSpace(titulo)

// 	fmt.Print("Fecha de Publicación (YYYY-MM-DD): ")
// 	fechaPublicacion, _ := reader.ReadString('\n')
// 	fechaPublicacion = strings.TrimSpace(fechaPublicacion)

// 	fmt.Print("Archivo (ruta/al/archivo.pdf): ")
// 	archivo, _ := reader.ReadString('\n')
// 	archivo = strings.TrimSpace(archivo)

// 	// Insertar un nuevo libro
// 	_, err := db.Exec("INSERT INTO Libro (Título, Fecha_Publicación, Archivo) VALUES ($1, $2, $3)",
// 		titulo, fechaPublicacion, archivo)
// 	if err != nil {
// 		fmt.Println("Error al ingresar el Libro", err)
// 		return
// 	}

// 	fmt.Println("El libro ha sido ingresado con exito")
// }
