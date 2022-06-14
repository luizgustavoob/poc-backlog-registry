package handlers

import (
	"encoding/json"
	"github.com/mercadolibre/lego-backlog-registry/internal/core/entities"
	"github.com/mercadolibre/lego-backlog-registry/internal/core/ports"
	"io"
	"net/http"
)

type createWorkOrderHandler struct {
	wo ports.WorkOrder
}

func NewCreateWorkOrderHandler(wo ports.WorkOrder) *createWorkOrderHandler {
	return &createWorkOrderHandler{
		wo: wo,
	}
}

func (h *createWorkOrderHandler) Method() string {
	return http.MethodPost
}

func (h *createWorkOrderHandler) Pattern() string {
	return "/commands"
}

func (h *createWorkOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var workOrder entities.Create

	err := json.NewDecoder(r.Body).Decode(&workOrder)
	if err != nil {
		result := make(map[string]string)
		result["error"] = err.Error()
		json, _ := json.Marshal(result)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}

	h.wo.Create()

	body, _ := io.ReadAll(r.Body)
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
