package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rikughi/go-quick-start/internal/service"
	"github.com/sirupsen/logrus"
)

type HelloController struct {
	Service *service.HelloService
	Log     *logrus.Logger
}

func NewHelloController(service *service.HelloService, log *logrus.Logger) *HelloController {
	return &HelloController{
		Service: service,
		Log:     log,
	}
}

func (h *HelloController) Hello(c *gin.Context) {
	h.Log.Info("hello")
	c.JSON(200, gin.H{
		"message": h.Service.Hello(),
	})
}
