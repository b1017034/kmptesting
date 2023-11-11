/*
 * ToDo API
 *
 * A simple ToDo API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	ToDoApiService := openapi.NewToDoApiService()
	ToDoApiController := openapi.NewToDoApiController(ToDoApiService)

	router := openapi.NewRouter(ToDoApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
