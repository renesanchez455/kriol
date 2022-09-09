// Filename: cmd/api/main.go
package main

import "log"

// The application version number
const version = "1.0.0"

// The configuration settings
type config struct {
	port int
	env  string // development, staging, production, etc.
}

// Dependency injection
type application struct {
	config config
	logger *log.Logger
}

func main() {

}
