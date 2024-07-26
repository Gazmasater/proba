// models.go
package main

// ResponseError представляет ошибку ответа
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
