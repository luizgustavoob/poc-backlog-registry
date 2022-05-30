package application

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func BuildRoutes(app *App) http.Handler {
	r := chi.NewRouter()

	//registry
	r.Get("/services", app.HandleListServices)
	r.Post("/services", app.HandleAddService)
	r.Delete("/services/{service_name}", app.HandleDeleteService)

	//work-order proxy
	r.Handle("/work-order/*", RemoteServiceInterceptor(app.registry, Gateway))

	return r
}
