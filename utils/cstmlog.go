package utils

import (
	"fmt"
	"os"
	"sync"
)

var (
	mu sync.Mutex
)

// Log appends the given name to a log file.
func Log(name string) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Attempting to log:", name) // Immediate feedback

	fd, err := os.OpenFile("./log.log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer fd.Close()

	_, err = fmt.Fprintln(fd, name)
	if err != nil {
		fmt.Printf("Error writing to log file: %v\n", err)
	}
}
