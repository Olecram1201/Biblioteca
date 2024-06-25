package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	// "github.com/joho/godotenv"
	// "database/sql"
	// "gorm.io/gorm"
)

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

	fmt.Printf("connection in time is:")
	fmt.Printf(dns)
	fmt.Printf("connection in time is:")

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	return nil, err
	// }

	// if err := db.Ping(); err != nil {
	// 	return nil, err
	// }
	// log.Println("Conexion abierta a la base de datos")
	// return db, nil
}
