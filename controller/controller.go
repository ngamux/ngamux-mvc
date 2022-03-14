package controller

import (
	"github.com/ngamux/ngamux"
	"github.com/ngamux/ngamux-mvc/model"
)

type Controller struct {
	model *model.Model
}

func New(mux *ngamux.Ngamux, model *model.Model) *Controller {
	controller := &Controller{
		model: model,
	}

	mux.Get("/", controller.Index)

	users := mux.Group("/users")
	users.Get("/", controller.GetAllUsers)
	users.Get("/create", controller.CreateUser)

	return controller
}
