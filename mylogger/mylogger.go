package mylogger

import (
	"fmt"
	"log"
)

// InfoLog ...
func InfoLog(template string, args ...interface{}) {
	if len(args) == 0 {
		log.Printf("[ INFO] %v\n", template)
	} else {
		log.Printf("[ INFO] %v\n", fmt.Sprintf(template, args...))
	}
}

// DebugLog ...
func DebugLog(template string, args ...interface{}) {
	if len(args) == 0 {
		log.Printf("[DEBUG] %v\n", template)
	} else {
		log.Printf("[DEBUG] %v\n", fmt.Sprintf(template, args...))
	}
}

// ErrorLog ...
func ErrorLog(template string, args ...interface{}) {
	if len(args) == 0 {
		log.Printf("[ERROR] %v\n", template)
	} else {
		log.Printf("[ERROR] %v\n", fmt.Sprintf(template, args...))
	}
}
