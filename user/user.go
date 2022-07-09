package user

import (
	"errors"
	"fmt"

	"goerrtag/errtag"
)

var ErrNotFound = errtag.New(errtag.NotFound, errors.New("user not found"))

func Update() error {
	return fmt.Errorf("failed to update a user: %w", ErrNotFound)
}
