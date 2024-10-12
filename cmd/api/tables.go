package main

import (
	"fmt"
	"net/http"
)

func (app *application) getAllTables(w http.ResponseWriter, r *http.Request) {
	tables, err := app.store.Tables.GetAllTables(r.Context())
	if err != nil {
		app.internalServerError(w, err)
		return
	}
	if err = writeJSON(w, 200, tables); err != nil {
		fmt.Printf("Error: %v\n", err)
		app.internalServerError(w, err)
		return
	}
}
