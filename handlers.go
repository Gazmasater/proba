// handler.go
package main

import (
	"fmt"
	"net/http"
)

// HandleTest1 обрабатывает POST-запросы по пути /test1
// @Summary Handle POST requests at /test1
// @Description Returns "привет Лев" for POST requests to /test1
// @Tags test
// @Accept  json
// @Produce  plain
// @Success 200 {string} string "привет Лев"
// @Failure 405 {object} ResponseError "Method not allowed"
// @Router /test1 [post]
func HandleTest1(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Отправляем ответ "привет Лев"
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "привет Лев")
}
