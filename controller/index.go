package controller

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

func (c *Controller) Index(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.JSON(rw, ngamux.Map{
		"message": "hore",
	})
}
