package repository

import (
	"database/sql"
	"fmt"

	"goerrtag/errtag"
)

func UpdateUser() error {
	err := errtag.New(errtag.Unexpected, sql.ErrConnDone)

	return fmt.Errorf("failed to update user entity: %w", err)
}
