package main

import (
	"td-project/configs"
	"td-project/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main()  {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/", handlers.Create)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.List)
	router.Get("/{id}", handlers.GetTask)	

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}