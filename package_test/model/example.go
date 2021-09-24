package model

import (
	"strconv"
)

var (
	a    = 1
	b    = true
	port = "8080"
)

var s int = 0

func Show() []string {
	return []string{strconv.Itoa(a), string(strconv.FormatBool(b)), port}
}

// fmt.Println(show())
