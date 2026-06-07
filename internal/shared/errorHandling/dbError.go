package errorHandling

import (
	"errors"

	"github.com/lib/pq"
)

func mapDBError(err error) (int, string) {
	var pqErr *pq.Error

	if errors.As(err, &pqErr) {
		switch pqErr.Code.Name() {
		case "unique_violation":
			return 409, "DUPLICATE_ENTRY"

		case "foreign_key_violation":
			return 400, "FOREIGN_KEY_VIOLATION"

		case "not_null_violation":
			return 400, "NOT_NULL_VIOLATION"

		case "check_violation":
			return 400, "CHECK_CONSTRAINT_VIOLATION"
		}
	}

	return 500, "INTERNAL_ERROR"
}
