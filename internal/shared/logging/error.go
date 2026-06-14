package logging

import "fmt"

func ErrorLog(action, message string) {
	fmt.Printf("Type: error\nAction: %s\nMessage: %s\n", action, message)
}
