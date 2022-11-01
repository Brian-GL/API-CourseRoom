package models

type ResponseInfrastructure struct {
	Status ResponseStatus `json:"status"`
	Data   any            `json:"data"`
}
