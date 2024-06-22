package gestionBiblioteca

import "time"

type ItemLibro struct {
	CodigoBarras    string
	NumeroPaginas   int
	Formato         string
	FechaDevolucion time.Time
}

func (b *ItemLibro) SetCodigoBarras(codigoBarras string) {
	b.CodigoBarras = codigoBarras
}

func (b *ItemLibro) SetNumeroPaginas(numeroPaginas int) {
	b.NumeroPaginas = numeroPaginas
}

func (b *ItemLibro) SetFormato(formato string) {
	b.Formato = formato
}

func (b *ItemLibro) SetFechaDevolucion(fechaDevolucion time.Time) {
	b.FechaDevolucion = fechaDevolucion
}

func (b *ItemLibro) GetCodigoBarras() string {
	return b.CodigoBarras
}

func (b *ItemLibro) GetNumeroPaginas() int {
	return b.NumeroPaginas
}

func (b *ItemLibro) GetFormato() string {
	return b.Formato
}

func (b *ItemLibro) GetFechaDevolucion() time.Time {
	return b.FechaDevolucion
}
