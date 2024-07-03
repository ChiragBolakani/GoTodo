package main

import (
	"fmt"
	"go_tutorials/pkg/config"
	"go_tutorials/pkg/db"
	"go_tutorials/pkg/middlewares"
	"go_tutorials/pkg/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.Init()

	database := db.NewDB()

	service := service.NewService(database)

	r := mux.NewRouter().StrictSlash(true)
	r.Use(middlewares.ContentTypeApplicationJsonMiddleware)

	r.HandleFunc("/api/v1/items", service.CreateItem).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/users/{user_id}/items", service.GetItems).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{user_id}/items/{id}", service.GetItem).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/items/{id}", service.UpdateItem).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/users/{user_id}/items/{id}", service.DeleteItem).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), r))
}
