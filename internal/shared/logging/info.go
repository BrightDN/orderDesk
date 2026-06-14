package logging

import "fmt"

func InfoLog(action, message string) {
	fmt.Printf("Type: info\n Action: %s\nMessage: %s\n", action, message)
}
