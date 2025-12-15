package main

import "net/http"
import "github.com/go-chi/chi/v5"
import "github.com/go-chi/chi/v5/middleware"
import "time"
import "log"
import "github.com/mjsandagi/go-ecommerce/internal/products"

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID) // Important for rate limiting
	r.Use(middleware.RealIP)    // Also important for rate limiting, as well as analytics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // Helps recover from crashes.

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second)) // Times out any requests that are taking longer than 60s to process.

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	// productService := products.NewService()
	productHandler := products.NewHandler(nil)
	r.Get("/products", productHandler.ListProducts)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.address,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Listening to requests at %s", app.config.address)

	return srv.ListenAndServe()

}

type application struct {
	config config
}

// For endpoint signatures, runs, and graceful shutdowns
type config struct {
	address string // e.g., 8080
	db      dbConfig
}

type dbConfig struct {
	dsn string
}
