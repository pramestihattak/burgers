package payloads

import "github.com/opam22/burgers/models"

type SuccessGetMenu struct {
	Code     int         `json:"code"`
	Response models.Menu `json:"response"`
}

type SuccessGetAllMenu struct {
	Code     int           `json:"code"`
	Response []models.Menu `json:"response"`
}
