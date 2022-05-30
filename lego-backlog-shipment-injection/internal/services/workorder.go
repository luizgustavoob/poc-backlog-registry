package services

import (
	"encoding/json"
	"log"

	"github.com/mercadolibre/lego-backlog-registry/internal/core/entities"
)

type WorkOrder struct{}

func NewWorkOrderService() *WorkOrder {
	return &WorkOrder{}
}

func (w *WorkOrder) Create() {
	log.Print("Faz de conta que estamos criando uma work-order aqui em shipment-injection garotinho...")
}

func (w *WorkOrder) SetState(idWorkOrder string, state entities.State) {
	log.Printf("Faz de conta que estamos atualiando o status da work-order %s para %s aqui em shipment-injection garotinho...",
		idWorkOrder, state)
}

func (w *WorkOrder) AddAssignee(idWorkOrder string, rep entities.Rep) {
	log.Printf("Faz de conta que estamos adicionando o rep %s na work-order %s aqui em shipment-injection garotinho...",
		rep.ID, idWorkOrder)
}

func (w *WorkOrder) AddFragment(idWorkOrder string, fragment map[string]json.RawMessage) {
	log.Printf("Faz de conta que estamos adicionando o fragment %s na work-order %s aqui em shipment-injection garotinho...",
		fragment, idWorkOrder)
}
