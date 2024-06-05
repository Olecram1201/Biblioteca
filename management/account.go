package main

import (
	"time"
)

type Account struct {
	Number string
	Opened time.Time
	State  string
}
