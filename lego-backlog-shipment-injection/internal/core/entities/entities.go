package entities

import "encoding/json"

type State string

type RemoteService struct {
	ProcessName string `json:"process_name"`
	Address     string `json:"address"`
	Scope       string `json:"scope,omitempty"`
}

type WorkOrder struct {
	ID      string           `json:"id"`
	Process string           `json:"process"`
	Status  string           `json:"status"`
	Params  *json.RawMessage `json:"params,omitempty"`
	Result  *json.RawMessage `json:"result,omitempty"`
}

type Create struct {
	WorkOrder `json:"work_order"`
}

type SetState struct {
	State State `json:"state"`
}

type Rep struct {
	ID string
}

type AddAssignee struct {
	Rep `json:"rep"`
}

type AddFragment struct {
	Fragment map[string]json.RawMessage `json:"fragment"`
}
