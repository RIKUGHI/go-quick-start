package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rikughi/go-quick-start/internal/delivery/http/controller"
	"github.com/rikughi/go-quick-start/internal/delivery/http/router"
	"github.com/rikughi/go-quick-start/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	Config *viper.Viper
	App    *gin.Engine
	Log    *logrus.Logger
}

func Bootstrap(app *App) {
	// setup servie
	helloService := service.NewHelloService()

	// setup controller
	helloController := controller.NewHelloController(helloService, app.Log)

	router := router.Router{
		App:             app.App,
		HelloController: helloController,
	}

	router.Setup()
}
