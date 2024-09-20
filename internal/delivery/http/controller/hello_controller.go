package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rikughi/go-quick-start/internal/service"
)

type HelloController struct {
	Service *service.HelloService
}

func NewHelloController(service *service.HelloService) *HelloController {
	return &HelloController{
		Service: service,
	}
}

func (h *HelloController) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello World",
	}); err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Error",
		})
	}
}
