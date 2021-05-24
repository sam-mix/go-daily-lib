package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type User struct {
	Name  string `validate:"min=2"`
	Age   int    `validate:"min=18"`
	Email string `validate:"email"`
}

func main() {
	validate := validator.New()

	u1 := User{
		Name:  "dj",
		Age:   18,
		Email: "dj@example.com",
	}
	err := validate.Struct(u1)
	if err != nil {
		fmt.Println(err)
	}

	u2 := User{
		Name:  "dj",
		Age:   18,
		Email: "djexample.com",
	}
	err = validate.Struct(u2)
	if err != nil {
		fmt.Println(err)
	}
}
