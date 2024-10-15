package main

import (
	"cyfrone/backend/internal/store"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   mysql.Config
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(EnableCORS)

	r.Use(middleware.Timeout(time.Second * 30))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/PointOfSale", func(r chi.Router) {
			r.Get("/GetAllPosTables", app.getAllTables)
		})
		r.Route("/orders", func(r chi.Router) {
			r.Get("/getOrderByTableId/{tableId}", app.getOrderByTableId)
			r.Get("/getStatusById/{orderStatusId}", app.getStatusById)
			r.Put("/updateOrderStatus/{orderId}/{orderStatusId}", app.updateOrderStatus)
		})
		r.Route("/takeaways", func(r chi.Router) {
			r.Get("/getAllTakeAwayOrders", app.getAllTakeawayOrders)
			r.Post("/createTakeAwayOrder", app.createTakeAwayOrder)
			r.Get("/getTakeAwayOrdersById/{orderId}", app.getTakeAwayOrdersById)
		})
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {

	serv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	log.Printf("Starting server on %s", serv.Addr)

	return serv.ListenAndServe()
}
