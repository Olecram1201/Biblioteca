package gestionBiblioteca

import "time"

type Cuenta struct {
	Numero  string
	Abierta time.Time
	Estado  string
}

func (c *Cuenta) SetNumero(numero string) {
	c.Numero = numero
}

func (c *Cuenta) SetAbierta(fecha time.Time) {
	c.Abierta = fecha
}

func (c *Cuenta) SetEstado(estado string) {
	c.Estado = estado
}

func (c *Cuenta) GetNumero() string {
	return c.Numero
}

func (c *Cuenta) GetAbierta() time.Time {
	return c.Abierta
}

func (c *Cuenta) GetEstado() string {
	return c.Estado
}
