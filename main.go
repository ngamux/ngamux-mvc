package main

import (
	"log"
	"net/http"

	"github.com/ngamux/ngamux"
	"github.com/ngamux/ngamux-mvc/config"
	"github.com/ngamux/ngamux-mvc/controller"
	"github.com/ngamux/ngamux-mvc/model"
)

func main() {
	mux := ngamux.New()

	database, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	models := model.New(database)
	controller.New(mux, models)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
