package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/mercadolibre/lego-backlog-registry/internal/core/entities"
	"github.com/mercadolibre/lego-backlog-registry/internal/core/ports"
	"net/http"
)

type setStatusWorkOrderHandler struct {
	wo ports.WorkOrder
}

func NewSetStatusWorkOrderHandler(wo ports.WorkOrder) *setStatusWorkOrderHandler {
	return &setStatusWorkOrderHandler{
		wo: wo,
	}
}

func (h *setStatusWorkOrderHandler) Method() string {
	return http.MethodPut
}

func (h *setStatusWorkOrderHandler) Pattern() string {
	return "/work-order/{id}/set-state"
}

func (h *setStatusWorkOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	handleError := func(w http.ResponseWriter, err error, statusCode int) {
		result := make(map[string]string)
		result["error"] = err.Error()
		js, _ := json.Marshal(result)
		w.WriteHeader(statusCode)
		w.Write(js)
	}

	if id == "" {
		handleError(w, errors.New("id work-order not informed"), http.StatusBadRequest)
		return
	}

	var state entities.SetState

	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	h.wo.SetState(id, state.State)

	w.WriteHeader(http.StatusNoContent)
}
