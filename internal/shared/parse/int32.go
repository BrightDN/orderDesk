package parse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
)

var ErrEmptyValue = errors.New("Value cannot be empty")
var ErrUnexpectedValue = errors.New("The value could not be converted")

func Int32(input string) (int32, *errorHandling.AppError) {

	if len(strings.TrimSpace(input)) == 0 {
		return 0, &errorHandling.AppError{
			Action:    "Parsing string to int32",
			LogError:  fmt.Errorf("Received an empty value"),
			UserError: ErrEmptyValue,
		}
	}
	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, &errorHandling.AppError{
			Action:    "Parsing string to int32",
			LogError:  fmt.Errorf("Received an unparseable value"),
			UserError: ErrUnexpectedValue,
		}
	}
	return int32(id), nil
}
