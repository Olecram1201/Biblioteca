//@autor: Erick Andrade
//@version: 2.0
//@fecha: 23/06/2024
//@description: Library Students Application for Oriented programming language

package main

import (
	"log"
	"net/http"

	"library_system/controllers"
	connect "library_system/db"

	"github.com/gorilla/mux"
)

func main() {
	connect.Connector()

	// Configuración de las rutas y el servidor web
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")

	router.HandleFunc("/autores", controllers.ListaAutoresHandler).Methods("GET") // Ruta para listar autores
	router.HandleFunc("/autores/nuevo", controllers.NuevoAutorFormHandler).Methods("GET")
	router.HandleFunc("/autores/nuevo", controllers.NuevoAutorHandler).Methods("POST")
	router.HandleFunc("/libros", controllers.ListaLibrosHandler).Methods("GET")
	router.HandleFunc("/libros/nuevo", controllers.NuevoLibroFormHandler).Methods("GET")
	router.HandleFunc("/libros/nuevo", controllers.NuevoLibroHandler).Methods("POST")

	router.HandleFunc("/error", controllers.ErrorHandler).Methods("GET")
	router.HandleFunc("/not-found", controllers.NotFoundHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(controllers.MethodNotAllowed)

	// Carpeta de archivos estáticos (CSS, JS, imágenes, etc.)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Configurar el servidor HTTP
	http.Handle("/", router)
	// Iniciar el servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
