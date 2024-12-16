package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/go-quick-start/internal/config"
)

func main() {
	app := gin.Default()
	viper := config.NewViper()
	logFileRotator := config.NewLogFileRotator("./logs", 7*24*time.Hour)
	logger := config.NewLogger()

	logger.SetOutput(logFileRotator)

	app.StaticFS("/logs", http.Dir(config.LogsDirectory))

	app.GET("/", func(c *gin.Context) {
		logger.Info("Welcome")
		c.JSON(http.StatusOK, gin.H{
			"message": "🔥 Welcome to the API 🔥",
		})
	})

	config.Bootstrap(&config.App{
		App:    app,
		Config: viper,
		Log:    logger,
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("PORT")),
		Handler: app.Handler(),
	}

	log.Printf("🚀 Server is listening on port %d", viper.GetInt("PORT"))

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %+v", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

	// log.Printf("🚀 Server is listening on port %d", viper.GetInt("PORT"))

	// if err := app.Run(); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
}
