package models

type ResposeInfrastruture struct {
	Status ResponseStatus `json:"status"`
	Data   any            `json:"data"`
}
