package service

import (
	"strconv"
)

var counter = 1

func GenerateShortCode() string {

	code := "url" + strconv.Itoa(counter)

	counter++

	return code
}