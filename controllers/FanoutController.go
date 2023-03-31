package controllers

import (
	"fanout_service/models"
	"log"
	"net/http"

	service "fanout_service/services"

	"github.com/gin-gonic/gin"
)

type fanoutController struct {
	l  *log.Logger
	pS service.IFanoutService
}

func (pC fanoutController) Publish(c *gin.Context) {
	var post models.Post

	if c.ShouldBind(&post) == nil {
		err := pC.pS.Publish(post)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusCreated)
		return
	}

	c.Status(http.StatusBadRequest)
}

func NewFanoutController(l *log.Logger, pS service.IFanoutService) IController {
	return fanoutController{
		l: l,
	}
}
