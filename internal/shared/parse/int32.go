package parse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrEmptyValue = errors.New("Value cannot be empty")
var ErrUnexpectedValue = errors.New("The value could not be converted")

func Int32(input string) (int32, error) {

	if len(strings.TrimSpace(input)) == 0 {
		fmt.Printf("Error: failed to parse: %v", ErrEmptyValue)
		return 0, ErrEmptyValue
	}
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error: failed to parse: %v", ErrUnexpectedValue)
		return 0, ErrUnexpectedValue
	}
	return int32(id), nil
}
