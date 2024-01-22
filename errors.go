package main

import (
	"errors"
	"fmt"
)

var (
	validationErrors = errors.New("validation errors")
	notFoundErrors   = errors.New("not found errors")
)

func Getbyid(id string) error {
	if id == "" {
		return validationErrors
	}
	if id != "ucon" {
		return notFoundErrors
	}

	return nil
}
func main() {
	err := Getbyid("lana")
	if err != nil {
		if errors.Is(err, validationErrors) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, notFoundErrors) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
