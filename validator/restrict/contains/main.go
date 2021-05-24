package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type User struct {
	Name string `validate:"containsrune=☻"`
	Age  int    `validate:"min=18"`
}

func main() {
	validate := validator.New()

	u1 := User{"d☻j", 18}
	err := validate.Struct(u1)
	if err != nil {
		fmt.Println(err)
	}

	u2 := User{"dj", 18}
	err = validate.Struct(u2)
	if err != nil {
		fmt.Println(err)
	}
}
