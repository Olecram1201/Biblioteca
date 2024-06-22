//@autor: Israael Espinoza
//@version: 1.0
//@fecha: 13/05/2024
//@descripcion: Esta aplicacion tiene el objetivo de gestionar un sistema de libros electrónicos.

package main

import (
	"database/sql"
	"fmt"
	gestionBiblioteca "gestionBiblioteca/gestion"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	//Importamos el paquene que nos ayudara a conectar con la BDD
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm" //Importamos el driver de  Postgres
)

var autores = []gestionBiblioteca.Autor{
	{
		Nombre:   "Gabriel",
		Apellido: "García Márquez",
		Libros: []gestionBiblioteca.Libro{
			{
				Titulo:           "Cien años de soledad",
				FechaPublicacion: time.Date(1967, time.May, 30, 0, 0, 0, 0, time.UTC),
				Archivo:          "/archivos/cien_anos_de_soledad.pdf",
			},
			{
				Titulo:           "El amor en los tiempos del cólera",
				FechaPublicacion: time.Date(1985, time.November, 8, 0, 0, 0, 0, time.UTC),
				Archivo:          "/archivos/el_amor_en_los_tiempos_del_colera.pdf",
			},
		},
	},
	{
		Nombre:   "J.K.",
		Apellido: "Rowling",
		Libros: []gestionBiblioteca.Libro{
			{
				Titulo:           "Harry Potter y la piedra filosofal",
				FechaPublicacion: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
				Archivo:          "/archivos/harry_potter_y_la_piedra_filosofal.pdf",
			},
			{
				Titulo:           "Harry Potter y la cámara secreta",
				FechaPublicacion: time.Date(1998, time.July, 2, 0, 0, 0, 0, time.UTC),
				Archivo:          "/archivos/harry_potter_y_la_camara_secreta.pdf",
			},
		},
	},
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var db *gorm.DB

// Main
func main() {
	// Crear una nueva instancia de ItemLibro
	libro := gestionBiblioteca.ItemLibro{}

	// Establecer valores utilizando los métodos Set
	//libro.SetCodigoBarras("9781234567890")
	//libro.SetNumeroPaginas(300)
	libro.SetFormato("Paperback")
	libro.SetFechaDevolucion(time.Now().AddDate(0, 0, 14)) // Ejemplo: Fecha de devolución en 14 días desde ahora

	// Obtener valores utilizando los métodos Get
	//fmt.Println("Código de Barras:", libro.GetCodigoBarras())
	//fmt.Println("Número de Páginas:", libro.GetNumeroPaginas())
	//fmt.Println("Formato:", libro.GetFormato())
	// db := conexionBdd()
	// defer db.Close()
	// menuInicio(db)

	cfg := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "Biblioteca",
	}

	db, err := conexionGORM(cfg)
	if err != nil {
		fmt.Printf("Error al conectar con la base de datos: %s\n", err)
		os.Exit(1)
	}
	fmt.Print(db)
	// defer db.Close()

	// CONFIGURACION PARA LA BASE DE DATOS
	// Auto migración para crear automáticamente las tablas en la base de datos
	err = db.AutoMigrate(&gestionBiblioteca.Autor{}, &gestionBiblioteca.Libro{})
	if err != nil {
		log.Fatalf("Error al realizar la migración de tablas: %v", err)
	}

	// CREAR UN AUTOR CON LIBROS MEDIANTE SIN ENDPOINTS

	// autor, err := gestionBiblioteca.CrearAutor(db, "Gabriel", "García Marquez")
	// if err != nil {
	// 	log.Fatalf("Error al crear el autor: %v", err)
	// }

	// libro1 := gestionBiblioteca.Libro{
	// 	Titulo:           "Cien años de soledad",
	// 	FechaPublicacion: time.Date(1967, time.May, 30, 0, 0, 0, 0, time.UTC),
	// 	Archivo:          "ruta/al/cien_anios_de_soledad.pdf",
	// 	AutorID:          autor.ID,
	// }
	// db.Create(&libro1)

	// libro2 := gestionBiblioteca.Libro{
	// 	Titulo:           "El amor en los tiempos del cólera",
	// 	FechaPublicacion: time.Date(1985, time.January, 1, 0, 0, 0, 0, time.UTC),
	// 	Archivo:          "ruta/al/amor_en_los_tiempos_del_colera.pdf",
	// 	AutorID:          autor.ID,
	// }
	// db.Create(&libro2)

	// // Mostrar información del autor y sus libros
	// fmt.Printf("Autor: %s %s\n", autor.Nombre, autor.Apellido)
	// fmt.Println("Libros:")
	// for _, libro := range autor.Libros {
	// 	fmt.Printf("- %s (Fecha de Publicación: %s)\n", libro.Titulo, libro.FechaPublicacion.Format("2006-01-02"))
	// }

	// // Ejemplo de obtener todos los libros de un autor
	// libros, err := ObtenerLibrosDeAutor(db, autor.ID)
	// if err != nil {
	// 	log.Fatalf("Error al obtener los libros del autor: %v", err)
	// }

	// // Mostrar los libros del autor
	// fmt.Printf("Libros de %s %s:\n", autor.Nombre, autor.Apellido)
	// for _, libro := range libros {
	// 	fmt.Printf("- %s (Fecha de Publicación: %s)\n", libro.Titulo, libro.FechaPublicacion.Format("2006-01-02"))
	// }

	// // Ejemplo de actualización de un autor
	// err = gestionBiblioteca.ActualizarAutor(db, autor.ID, "Gabriel", "G. Marquez")
	// if err != nil {
	// 	log.Fatalf("Error al actualizar el autor: %v", err)
	// }

	// Ejemplo de eliminación de un autor (esto también eliminará sus libros debido a la restricción CASCADE)
	// err = gestionBiblioteca.EliminarAutor(db, autor.ID)
	// if err != nil {
	// 	log.Fatalf("Error al eliminar el autor: %v", err)
	// }

	// autorss := gestionBiblioteca.Autor{
	// 	Nombre:   "Gabriel García",
	// 	Apellido: "Marquez",
	// 	Libros: []gestionBiblioteca.Libro{
	// 		{
	// 			Titulo:           "Cien años de soledad",
	// 			FechaPublicacion: time.Date(1967, time.May, 30, 0, 0, 0, 0, time.UTC),
	// 			Archivo:          "ruta/al/cien_anios_de_soledad.pdf",
	// 		},
	// 		{
	// 			Titulo:           "El amor en los tiempos del cólera",
	// 			FechaPublicacion: time.Date(1985, time.January, 1, 0, 0, 0, 0, time.UTC),
	// 			Archivo:          "ruta/al/amor_en_los_tiempos_del_colera.pdf",
	// 		},
	// 	},
	// }

	// // Crear en la base de datos
	// err = autor.IngresoAutor(db)
	// if err != nil {
	// 	log.Fatalf("Error al crear el autor y sus libros: %v", err)
	// }

	// // Mostrar los libros
	// err = autor.MostrarLibros(db)
	// if err != nil {
	// 	log.Fatalf("Error al mostrar los libros del autor: %v", err)
	// }

	// Configuración de las rutas y el servidor web
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/autores", listaAutoresHandler).Methods("GET") // Ruta para listar autores
	router.HandleFunc("/autores/nuevo", nuevoAutorFormHandler).Methods("GET")
	router.HandleFunc("/autores/nuevo", nuevoAutorHandler).Methods("POST")
	router.HandleFunc("/libros", listaLibrosHandler).Methods("GET")
	router.HandleFunc("/libros/nuevo", nuevoLibroFormHandler).Methods("GET")
	router.HandleFunc("/libros/nuevo", nuevoLibroHandler).Methods("POST")

	// Carpeta de archivos estáticos (CSS, JS, imágenes, etc.)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Configurar el servidor HTTP
	http.Handle("/", router)
	// Iniciar el servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func conexionGORM(cfg Config) (*gorm.DB, error) {
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

func ObtenerLibrosDeAutor(db *gorm.DB, autorID uint) ([]gestionBiblioteca.Libro, error) {
	var libros []gestionBiblioteca.Libro

	// Cargar todos los libros del autor con el ID especificado
	result := db.Where("autor_id = ?", autorID).Find(&libros)
	if result.Error != nil {
		return nil, result.Error
	}

	return libros, nil
}

// Handlers

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
	// fmt.Fprintf(w, "Bienvenido a la biblfasdfadsioteca virtual")

}

func listaAutoresHandler(w http.ResponseWriter, r *http.Request) {
	// Renderizar la plantilla HTML con los datos ficticios de autores
	if err := renderTemplate(w, "lista_autores.html", autores); err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}

func nuevoAutorFormHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "nuevo_autor.html", nil)
}

func nuevoAutorHandler(w http.ResponseWriter, r *http.Request) {
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")

	autor := gestionBiblioteca.Autor{
		Nombre:   nombre,
		Apellido: apellido,
	}

	db.Create(&autor)
	http.Redirect(w, r, "/autores", http.StatusSeeOther)
}

func listaLibrosHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí podrías usar la variable global autores para renderizar la lista de libros
	// en una plantilla HTML o devolverla como JSON, dependiendo de tus necesidades.
	// Por simplicidad, aquí solo imprimimos los títulos de los libros.
	for _, autor := range autores {
		for _, libro := range autor.Libros {
			_, _ = w.Write([]byte("Libro: " + libro.Titulo + "\n"))
		}
	}
}

// var libros []gestionBiblioteca.Libro
// if err := db.Find(&libros).Error; err != nil {
// 	http.Error(w, "Error al buscar libros en la base de datos", http.StatusInternalServerError)
// 	return
// }

// // Renderizar la plantilla HTML con los datos de libros
// if err := renderTemplate(w, "lista_libros.html", libros); err != nil {
// 	http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
// 	return
// }

func nuevoLibroFormHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "nuevo_libro.html", nil)
}

func nuevoLibroHandler(w http.ResponseWriter, r *http.Request) {
	titulo := r.FormValue("titulo")
	archivo := r.FormValue("archivo")
	fechaPublicacion, _ := time.Parse("2006-01-02", r.FormValue("fecha_publicacion"))

	libro := gestionBiblioteca.Libro{
		Titulo:           titulo,
		Archivo:          archivo,
		FechaPublicacion: fechaPublicacion,
	}

	db.Create(&libro)
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

// funcion para renderizar la plantilla localizada en ./templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	// Cargar la plantilla desde el sistema de archivos
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		return err
	}

	// Ejecutar la plantilla y escribir el resultado en la respuesta HTTP
	err = t.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

/////////////////////////////////////////////////////

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

// Funcion para ingresar un libro
// En este caso no se solicita un ID puesto que la BDD
// se encargara de crear uno mediante parametros de identidad.
