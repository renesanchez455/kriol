// Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createEntryHandler for the "POST /v1/entries" endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new entry..")
}

// showEntryHandler for the "GET /v1/entries/:id" endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Use the "ParamsFromContext()" function to get a request context as a slice
	params := httprouter.ParamsFromContext(r.Context())

	// GET the value of the "id" parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Display the entry id
	fmt.Fprintf(w, "show the details for entry %d\n", id)
}
