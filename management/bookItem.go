package main

import (
	"time"
)

type BookItem struct {
	Barcode       string
	NumberOfPages int
	Format        string
	DueDate       time.Time
}
