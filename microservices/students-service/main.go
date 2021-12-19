package main

import (
	"go-projects/microservices/students-service/api"
	"go-projects/microservices/students-service/core"
)

func main() {
	router := core.Router{}
	router.Init(core.RouterConfig{StrictSlash: true})

	router.HandleFunc("/", api.Health)
	router.HandleFunc("/students", api.AddNewStudent).Methods("POST")
	router.HandleFunc("/students", api.UpdateStudent).Methods("PUT")

	router.HttpListenAndServe()
}
