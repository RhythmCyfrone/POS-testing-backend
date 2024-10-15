package main

import (
	"cyfrone/backend/internal/store"
	"fmt"
	"net/http"
)

type tablesResponse struct {
	CountOfActualCapacity int           `json:"countOfActualCapacity"`
	AvgTableOccupancy     float64       `json:"avgTableOccupancy"`
	AvgTableTurnOverTime  float64       `json:"avgTableTurnOverTime"`
	ServingTableDetails   []store.Table `json:"servingTableDetails"`
}

func (app *application) getAllTables(w http.ResponseWriter, r *http.Request) {
	tables, err := app.store.Tables.GetAllTables(r.Context())
	if err != nil {
		app.internalServerError(w, err)
		return
	}
	data := tablesResponse{
		CountOfActualCapacity: 80,
		AvgTableOccupancy:     30.0,
		AvgTableTurnOverTime:  107.69,
		ServingTableDetails:   tables,
	}
	if err = writeJSON(w, 200, data); err != nil {
		fmt.Printf("Error: %v\n", err)
		app.internalServerError(w, err)
		return
	}
}
