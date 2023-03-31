package controllers

import "github.com/gin-gonic/gin"

type IController interface {
	Publish(c *gin.Context)
}
