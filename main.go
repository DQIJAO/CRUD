package main

import (
	"log"
	"net/http"

	"github.com/CRUD/pkg/app"
	"github.com/CRUD/pkg/model"
)

const (
	port     = ":8080"
	mongoURL = "mongodb://127.0.0.1:27017"
)

func main() {
	mux := http.NewServeMux()

	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v ", err)
	}
	http.ListenAndServe(port, mux)
}
