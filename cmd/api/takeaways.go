package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) getAllTakeawayOrders(w http.ResponseWriter, r *http.Request) {
	takeawayOrders, err := app.store.Takeaways.GetTakeawayOrders(r.Context())
	if err != nil {
		app.internalServerError(w, err)
		return
	}
	if err = writeJSON(w, 200, takeawayOrders); err != nil {
		fmt.Printf("Error: %v\n", err)
		app.internalServerError(w, err)
		return
	}
}

type CustomerDetails struct {
	CustomerName  string  `json:"customerName"`
	CustomerPhone *string `json:"customerPhone"`
}

func (app *application) createTakeAwayOrder(w http.ResponseWriter, r *http.Request) {

	var order CustomerDetails

	// Parse the JSON request body into the struct
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		app.invalidPayloadError(w, err)
		return
	}

	newOrder, err := app.store.Takeaways.CreateTakeawayOrder(r.Context(), order.CustomerName, order.CustomerPhone)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		app.internalServerError(w, err)
		return
	}

	if err = writeJSON(w, 200, newOrder); err != nil {
		fmt.Printf("Error: %v\n", err)
		app.internalServerError(w, err)
		return
	}
}

func (app *application) getTakeAwayOrdersById(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "orderId")
	fmt.Println(orderId)
	takeawayOrder, err := app.store.Takeaways.GetTakeawayOrdersById(r.Context(), orderId)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		app.internalServerError(w, err)
		return
	}
	if err = writeJSON(w, 200, takeawayOrder); err != nil {
		fmt.Printf("Error: %v\n", err)
		app.internalServerError(w, err)
		return
	}
}
