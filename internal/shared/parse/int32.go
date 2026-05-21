package parse

import (
	"errors"
	"strconv"
	"strings"
)

var ErrEmptyValue = errors.New("Value cannot be empty")
var ErrUnexpectedValue = errors.New("The value could not be converted")

func Int32(input string) (int32, error) {

	if len(strings.TrimSpace(input)) == 0 {
		return 0, ErrEmptyValue
	}
	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, ErrUnexpectedValue
	}
	return int32(id), nil
}
