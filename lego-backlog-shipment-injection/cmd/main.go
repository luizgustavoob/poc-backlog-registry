package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mercadolibre/lego-backlog-registry/internal/adapters/handlers"
	"github.com/mercadolibre/lego-backlog-registry/internal/services"
	"github.com/mercadolibre/lego-backlog-registry/server"
)

func main() {
	registry := services.NewRegistry(&http.Client{Timeout: 500 * time.Millisecond})
	workOrderService := services.NewWorkOrderService()

	// handlers
	pingHandler := handlers.NewPingHandler()
	createWorkOrderHandler := handlers.NewCreateWorkOrderHandler(workOrderService)
	setStatusWorkOrderHandler := handlers.NewSetStatusWorkOrderHandler(workOrderService)
	addAssigneeWorkOrderHandler := handlers.NewAddAssigneeWorkOrderHandler(workOrderService)
	addFragmentWorkOrderHandler := handlers.NewAddFragmentWorkOrderHandler(workOrderService)

	// server
	srv := server.NewServer(registry, pingHandler, createWorkOrderHandler,
		setStatusWorkOrderHandler, addAssigneeWorkOrderHandler, addFragmentWorkOrderHandler)
	srv.ListenAndServe()

	// shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	srv.Shutdown()
}
