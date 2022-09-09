// Filename: cmd/api/main.go
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

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
	var cfg config
	// read in the flags that are needed to populate our config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()
	//Create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//Create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Create our new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

}
