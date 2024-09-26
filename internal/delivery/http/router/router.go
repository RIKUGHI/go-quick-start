package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rikughi/go-quick-start/internal/delivery/http/controller"
)

type Router struct {
	App             *gin.Engine
	HelloController *controller.HelloController
}

func (r *Router) Setup() {
	r.App.GET("/", r.HelloController.Hello)
}
