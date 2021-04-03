package mylogger

import (
	"fmt"
)

// InfoLog ...
func InfoLog(template string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Printf("[ INFO] %s\n", template)
	} else {
		fmt.Printf("[ INFO] %s\n", fmt.Sprintf(template, args...))
	}
}

// DebugLog ...
func DebugLog(template string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Printf("[DEBUG] %v\n", template)
	} else {
		fmt.Printf("[DEBUG] %v\n", fmt.Sprintf(template, args...))
	}
}

// ErrorLog ...
func ErrorLog(template string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Printf("[ERROR] %v\n", template)
	} else {
		fmt.Printf("[ERROR] %v\n", fmt.Sprintf(template, args...))
	}
}
