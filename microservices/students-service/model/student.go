package model

import (
	"github.com/gocql/gocql"
)

type Student struct {
	Id        gocql.UUID `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Age       int        `json:"age"`
	Marks     int        `json:"marks"`
}
