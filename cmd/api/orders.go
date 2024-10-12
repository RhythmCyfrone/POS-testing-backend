package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) getOrderByTableId(w http.ResponseWriter, r *http.Request) {
	order, err := app.store.Orders.GetOrderByTableId(r.Context(), chi.URLParam(r, "tableId"))
	if err != nil {
		app.internalServerError(w, err)
		return
	}
	if err = writeJSON(w, 200, order); err != nil {
		app.internalServerError(w, err)
		return
	}
}

func (app *application) getStatusById(w http.ResponseWriter, r *http.Request) {
	orderStatusID, _ := strconv.Atoi(chi.URLParam(r, "orderStatusId"))

	status, err := app.store.Status.GetOrderStatusById(r.Context(), orderStatusID)
	if err != nil {
		app.internalServerError(w, err)
		return
	}
	if err = writeJSON(w, 200, status); err != nil {
		app.internalServerError(w, err)
		return
	}
}

func (app *application) updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "orderId")
	orderStatusID, _ := strconv.Atoi(chi.URLParam(r, "orderStatusId"))

	err := app.store.Status.UpdateOrderStatusByOrderId(r.Context(), orderID, orderStatusID)
	if err != nil {
		app.internalServerError(w, err)
		return
	}
	if err = writeJSON(w, 200, nil); err != nil {
		app.internalServerError(w, err)
		return
	}
}
