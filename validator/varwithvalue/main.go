package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

func main() {
	name1 := "dj"
	name2 := "dj2"

	validate := validator.New()
	fmt.Println(validate.VarWithValue(name1, name2, "eqfield"))

	fmt.Println(validate.VarWithValue(name1, name2, "nefield"))
}
