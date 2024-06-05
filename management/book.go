package main

import (
	"time"
)

type Book struct {
	ISBN            string
	Name            string
	Subject         string
	Overview        string
	Publisher       string
	PublicationDate time.Time
	Lang            string
}
