package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/models"
	"github.com/opam22/burgers/payloads"
)

type OrderController struct {
	interfaces.IOrderService
}

func (controller *OrderController) Order(res http.ResponseWriter, req *http.Request) {
	var order models.Order

	json.NewDecoder(req.Body).Decode(&order)

	bill, err := controller.MakeOrder(order)
	if err != nil {
		json.NewEncoder(res).Encode(payloads.ErrorRequest{Code: http.StatusInternalServerError, Response: err.Error()})
	}

	json.NewEncoder(res).Encode(payloads.SuccessOrder{Code: http.StatusOK, Response: bill})
}
