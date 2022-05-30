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

	//work-order
	r.Post("/work-order/create", RemoteServiceInterceptor(app.registry, Gateway))
	r.Put("/work-order/{id}/set-state", RemoteServiceInterceptor(app.registry, Gateway))
	r.Put("/work-order/{id}/add-assignee", RemoteServiceInterceptor(app.registry, Gateway))
	r.Put("/work-order/{id}/add-fragment", RemoteServiceInterceptor(app.registry, Gateway))
	//...another commands

	return r
}
