package gestionBiblioteca

import (
	"fmt"

	"gorm.io/gorm"
)

type Autor struct {
	gorm.Model
	Nombre   string  `gorm:"not null"`
	Apellido string  `gorm:"not null"`
	Libros   []Libro // Relación uno a muchos con Libro
}

// Constructor que permite el ingreso de un autor nuevo
func (a *Autor) IngresoAutor(db *gorm.DB) error {
	result := db.Create(&a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CrearAutor(db *gorm.DB, nombre, apellido string) (*Autor, error) {
	autor := Autor{
		Nombre:   nombre,
		Apellido: apellido,
	}

	result := db.Create(&autor)
	if result.Error != nil {
		return nil, result.Error
	}

	return &autor, nil
}

func ObtenerAutorYLibros(db *gorm.DB, id uint) (*Autor, error) {
	var autor Autor
	result := db.Preload("Libros").First(&autor, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &autor, nil
}

func ActualizarAutor(db *gorm.DB, id uint, nombre, apellido string) error {
	var autor Autor
	result := db.First(&autor, id)
	if result.Error != nil {
		return result.Error
	}

	autor.Nombre = nombre
	autor.Apellido = apellido

	result = db.Save(&autor)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func EliminarAutor(db *gorm.DB, id uint) error {
	result := db.Delete(&Autor{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Mostrar los libros pertenecientes a cada autor
func (a *Autor) MostrarLibros(db *gorm.DB) error {
	var autor Autor
	if err := db.Preload("Libros").First(&autor, a.ID).Error; err != nil {
		return err
	}

	fmt.Printf("Libros de %s %s:\n", autor.Nombre, autor.Apellido)
	for _, libro := range autor.Libros {
		fmt.Printf("- %s (Fecha de Publicación: %s)\n", libro.Titulo, libro.FechaPublicacion.Format("2006-01-02"))
	}
	return nil
}

// zona de Set's
func (a *Autor) SetNombre(nombre string) {
	a.Nombre = nombre
}

func (a *Autor) SetApellido(apellido string) {
	a.Apellido = apellido
}

// zona de Get's
func (a *Autor) GetNombre() string {
	return a.Nombre
}

func (a *Autor) GetApellido() string {
	return a.Apellido
}
