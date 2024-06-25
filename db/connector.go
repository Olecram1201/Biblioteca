package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/joho/godotenv"
	// "database/sql"
	// "gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func Connector() {
	err := godotenv.Load()
	if err != nil {
		// return nil, err
		fmt.Print("failed to connect to the databse")
	}
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	println(dns)

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	return nil, err
	// }

	// if err := db.Ping(); err != nil {
	// 	return nil, err
	// }
	// log.Println("Conexion abierta a la base de datos")
	// return db, nil

	// // Auto migración para crear automáticamente las tablas en la base de datos
	// err = db.AutoMigrate(&management.Autor{}, &management.Libro{})
	// if err != nil {
	// 	log.Fatalf("Error al realizar la migración de tablas: %v", err)
	// }

}

func GORMConnection(cfg Config) (*gorm.DB, error) {
	// Cadena de conexión
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// Conectar a la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %s", err)
	}

	// Comprobar la conexión
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error al obtener el DB object: %s", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al hacer ping a la base de datos: %s", err)
	}

	fmt.Println("Conexión a la base de datos realizada con éxito")

	return db, nil
}

// Funcion que nos da la conexion con la BDD
func conexionBdd() *sql.DB {
	//conexion con la BDD
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "Biblioteca"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Print("PSQLINFO", psqlInfo)
	//Establecemos la conexion
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error al conectar con la base de datos", err)
		os.Exit(1) // el valor de 1 indica q el codigo a generado un error por lo que se usa osExit para una  salida inmediata
	}

	//Comprobacion de la conexion
	err = db.Ping()
	if err != nil {
		fmt.Println("Error al conectar con la base de datos", err)
		os.Exit(1) // el valor de 1 indica q el codigo a generado un error por lo que se usa os.Exit para una  salida inmediata
	}

	fmt.Println("Conexion a la BASE DE DATOS realizada con EXITO")
	return db
}
