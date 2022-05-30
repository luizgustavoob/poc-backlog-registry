package main

import (
	"github.com/mercadolibre/lego-backlog-registry/infrastructure/database"
	"github.com/mercadolibre/lego-backlog-registry/internal/application"
	"github.com/mercadolibre/lego-backlog-registry/internal/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db := database.NewInMemoryDatabase()

	app := application.BuildApp(db)
	routes := application.BuildRoutes(app)
	srv := application.NewServer(routes)

	//run
	srv.ListenAndServe()

	//hearbeat
	services.Heartbeat(db, &http.Client{Timeout: 800 * time.Millisecond}).Run()

	//shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	srv.Shutdown()
}
