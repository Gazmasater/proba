// router.go
package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter создает новый маршрутизатор и настраивает маршруты
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Обрабатываем GET-запросы по пути /test
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "привет")
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // путь к сгенерированной документации
	))

	// Обрабатываем POST-запросы по пути /test1
	r.Post("/test1", HandleTest1)

	return r
}
