package management

import (
	"time"

	"gorm.io/gorm"
)

type Libro struct {
	gorm.Model
	Titulo           string    `gorm:"not null"`
	FechaPublicacion time.Time `gorm:"not null"`
	Archivo          string    `gorm:"not null"`
	AutorID          uint      `gorm:"not null"` // Clave externa para la relación con Autor
}

// Constructor que permite el ingreso de un libro nuevo
func (l *Libro) IngresoLibro(db *gorm.DB) error {
	result := db.Create(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// zona de Sett's
func (l *Libro) SetFecha(fecha string) {
	l.FechaPublicacion, _ = time.Parse("2006-01-02", fecha)
}

func (l *Libro) SetTitulo(titulo string) {
	l.Titulo = titulo
}

func (l *Libro) SetArchivo(archivo string) {
	l.Archivo = archivo
}

// zona de Get's
func (l *Libro) GetFecha() string {
	return l.FechaPublicacion.Format("2006-01-02")
}

func (l *Libro) GetTitulo() string {
	return l.Titulo
}

func (l *Libro) GetArchivo() string {
	return l.Archivo
}
