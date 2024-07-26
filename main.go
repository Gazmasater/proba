// main.go
package main

import (
	"log"
	"net/http"

	"github.com/Gazmasater/api"
	"github.com/Gazmasater/docs"
)

func main() {
	// Создаем новый маршрутизатор
	r := api.NewRouter()

	docs.SwaggerInfo.Title = "API TEST"
	docs.SwaggerInfo.Description = "Это пример API для отправки сообщений."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Запускаем HTTP сервер на порту 8080
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
