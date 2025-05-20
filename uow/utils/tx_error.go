package utils

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func HandleTXError(tx *sqlx.Tx, method string, err error) error {
	if txErr := tx.Rollback(); txErr != nil {
		return fmt.Errorf("tx.Rollback: %w. %s: %s", txErr, method, err)
	}

	return fmt.Errorf("%s: %w", method, err)
}
