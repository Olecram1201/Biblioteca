package controllers

import (
	"html/template"
	"library_system/management"
	"net/http"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

var autores = []management.Autor{
	{
		Nombre:   "Gabriel",
		Apellido: "García Márquez",
		Libros: []management.Libro{
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
		Libros: []management.Libro{
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

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {

	t, err := template.ParseFiles("templates/base.html", "templates/"+tmpl)
	if err != nil {
		return err
	}

	// Render template
	// t.Execute(w, data)
	t.ExecuteTemplate(w, "base", nil)

	return nil
}

func ObtenerLibrosDeAutor(db *gorm.DB, autorID uint) ([]management.Libro, error) {
	var libros []management.Libro

	// Cargar todos los libros del autor con el ID especificado
	result := db.Where("autor_id = ?", autorID).Find(&libros)
	if result.Error != nil {
		return nil, result.Error
	}

	return libros, nil
}

// Handlers

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html", nil)
}
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	var libros []management.Libro

	result := db.Where("autor_id = ?", 1).Find(&libros)
	println(result)

	RenderTemplate(w, "books.html", result)
}

func AuthorsHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "authors.html", nil)
}

func ListaAutoresHandler(w http.ResponseWriter, r *http.Request) {
	// Renderizar la plantilla HTML con los datos ficticios de autores
	if err := RenderTemplate(w, "lista_autores.html", autores); err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}

func NuevoAutorFormHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "nuevo_autor.html", nil)
}

func NuevoAutorHandler(w http.ResponseWriter, r *http.Request) {
	nombre := r.FormValue("nombre")
	apellido := r.FormValue("apellido")

	autor := management.Autor{
		Nombre:   nombre,
		Apellido: apellido,
	}

	db.Create(&autor)
	http.Redirect(w, r, "/autores", http.StatusSeeOther)
}

func ListaLibrosHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí podrías usar la variable global autores para renderizar la lista de libros
	// en una plantilla HTML o devolverla como JSON, dependiendo de tus necesidades.
	// Por simplicidad, aquí solo imprimimos los títulos de los libros.
	for _, autor := range autores {
		for _, libro := range autor.Libros {
			_, _ = w.Write([]byte("Libro: " + libro.Titulo + "\n"))
		}
	}
}

func NuevoLibroFormHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "nuevo_libro.html", nil)
}

func NuevoLibroHandler(w http.ResponseWriter, r *http.Request) {
	titulo := r.FormValue("titulo")
	archivo := r.FormValue("archivo")
	fechaPublicacion, _ := time.Parse("2006-01-02", r.FormValue("fecha_publicacion"))

	libro := management.Libro{
		Titulo:           titulo,
		Archivo:          archivo,
		FechaPublicacion: fechaPublicacion,
	}

	db.Create(&libro)
	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}
