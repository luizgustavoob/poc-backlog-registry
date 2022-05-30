package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/mercadolibre/lego-backlog-registry/internal/core/entities"
	"github.com/mercadolibre/lego-backlog-registry/internal/core/ports"
	"net/http"
)

type addAssigneeWorkOrderHandler struct {
	wo ports.WorkOrder
}

func NewAddAssigneeWorkOrderHandler(wo ports.WorkOrder) *addAssigneeWorkOrderHandler {
	return &addAssigneeWorkOrderHandler{
		wo: wo,
	}
}

func (h *addAssigneeWorkOrderHandler) Method() string {
	return http.MethodPut
}

func (h *addAssigneeWorkOrderHandler) Pattern() string {
	return "/work-order/{id}/add-assignee"
}

func (h *addAssigneeWorkOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var assignee entities.AddAssignee

	err := json.NewDecoder(r.Body).Decode(&assignee)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	h.wo.AddAssignee(id, assignee.Rep)

	w.WriteHeader(http.StatusNoContent)
}
