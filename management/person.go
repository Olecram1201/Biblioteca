package main

import (
	"time"
)

type Person struct {
	CI          string
	FullName    string
	Age         int
	Phone       string
	Birthdate   time.Time
	Ethnicity   string
	Nationality string
	CivilStatus string
}
