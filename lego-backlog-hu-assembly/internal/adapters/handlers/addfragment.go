package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mercadolibre/lego-backlog-registry/internal/core/entities"
	"github.com/mercadolibre/lego-backlog-registry/internal/core/ports"
)

type addFragmentWorkOrderHandler struct {
	wo ports.WorkOrder
}

func NewAddFragmentWorkOrderHandler(wo ports.WorkOrder) *addFragmentWorkOrderHandler {
	return &addFragmentWorkOrderHandler{
		wo: wo,
	}
}

func (h *addFragmentWorkOrderHandler) Method() string {
	return http.MethodPut
}

func (h *addFragmentWorkOrderHandler) Pattern() string {
	return "/work-order/{id}/add-fragment"
}

func (h *addFragmentWorkOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var fragment entities.AddFragment

	err := json.NewDecoder(r.Body).Decode(&fragment)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	h.wo.AddFragment(id, fragment.Fragment)

	w.WriteHeader(http.StatusNoContent)
}
