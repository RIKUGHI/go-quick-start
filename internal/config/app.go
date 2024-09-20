package config

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rikughi/go-quick-start/internal/delivery/http/controller"
	"github.com/rikughi/go-quick-start/internal/delivery/http/router"
	"github.com/rikughi/go-quick-start/internal/service"
	"github.com/spf13/viper"
)

type App struct {
	Config *viper.Viper
}

func NewApp(app *App) *App {
	return &App{
		Config: app.Config,
	}
}

func (a *App) Run() error {
	// setup servie
	helloService := service.NewHelloService()

	// setup controller
	helloController := controller.NewHelloController(helloService)

	router := router.Router{
		HelloController: helloController,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.Config.GetInt("PORT")),
		Handler:      router.Handler(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("🚀 Server is listening on port %d", a.Config.GetInt("PORT"))

	return server.ListenAndServe()
}
