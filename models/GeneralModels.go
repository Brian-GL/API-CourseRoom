package models

type ResponseStatus int

const (
	SUCCESS = 1
	ALERT   = 2
	ERROR   = 3
)

type ResponseInfrastructure struct {
	Status ResponseStatus `json:"status"`
	Data   any            `json:"data"`
}
