package ports

import (
	"encoding/json"

	"github.com/mercadolibre/lego-backlog-registry/internal/core/entities"
)

type WorkOrder interface {
	Create()
	SetState(idWorkOrder string, state entities.State)
	AddAssignee(idWorkOrder string, rep entities.Rep)
	AddFragment(idWorkOrder string, fragment map[string]json.RawMessage)
}
