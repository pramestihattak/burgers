package interfaces

import "github.com/opam22/burgers/models"

type IOrderService interface {
	MakeOrder(order models.Order) (models.Order, error)
}
