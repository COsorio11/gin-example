package main

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
	"strconv"
)

type User struct {
	Id        int    `db:"id json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}
