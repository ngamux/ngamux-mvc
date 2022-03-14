package controller

import (
	"net/http"

	"github.com/ngamux/ngamux"
	"github.com/ngamux/ngamux-mvc/model"
)

func (c *Controller) GetAllUsers(rw http.ResponseWriter, r *http.Request) error {
	users, err := c.model.GetAllUsers(r.Context())
	if err != nil {
		return ngamux.JSONWithStatus(rw, http.StatusInternalServerError, ngamux.Map{
			"error": err.Error(),
		})
	}

	return ngamux.JSON(rw, ngamux.Map{
		"users": users,
	})
}

func (c *Controller) CreateUser(rw http.ResponseWriter, r *http.Request) error {
	user := model.User{
		Email:    ngamux.GetQuery(r, "email"),
		Password: ngamux.GetQuery(r, "password"),
	}
	userCreated, err := c.model.CreateUser(r.Context(), user)
	if err != nil {
		return ngamux.JSONWithStatus(rw, http.StatusInternalServerError, ngamux.Map{
			"error": err.Error(),
		})
	}

	return ngamux.JSON(rw, ngamux.Map{
		"user": userCreated,
	})
}
