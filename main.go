package main

import (
	"errors"
	"fmt"

	"goerrtag/errtag"
	"goerrtag/repository"
	"goerrtag/user"
)

func main() {
	var tag *errtag.Tag

	err := user.Update()

	if errors.As(err, &tag) {
		fmt.Println(tag.Code)        // custom tag attribute
		fmt.Println(tag.Err.Error()) // tagged error message
		fmt.Println(err.Error())     // full error message
	}

	err = repository.UpdateUser()

	if errors.As(err, &tag) {
		fmt.Println(tag.Code)
		fmt.Println(tag.Error())
		fmt.Println(err.Error())
	}
}
