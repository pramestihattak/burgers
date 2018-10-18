package services

import (
	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/models"
)

type OrderService struct {
	interfaces.IOrderRepository
}

func (service *OrderService) MakeOrder(order models.Order) (models.Order, error) {

	for _, orderMenu := range order.Orders {
		menuReceipt, errMenu := service.GetMenuIngredients(orderMenu.IdMenu)
		if errMenu != nil {
			return models.Order{}, errMenu
		}

		errCook := service.cook(menuReceipt)
		if errCook != nil {
			return models.Order{}, errCook
		}

	}

	orderedID, totalPrice, err := service.generateBill(order)
	if err != nil {
		return models.Order{}, err
	}

	order.TotalPrice = totalPrice
	order.Status = "Your order is success"
	order.OrderCode = orderedID

	return order, nil
}

func (service *OrderService) generateBill(order models.Order) (int64, int, error) {
	totalPrice, err := service.getTotalPrice(order.Orders)
	if err != nil {
		return 0, 0, err
	}
	order.TotalPrice = totalPrice
	orderedID, err := service.CreateBill(order)
	if err != nil {
		return 0, 0, err
	}

	errStore := service.StoreOrderedMenu(orderedID, order.Orders)
	if errStore != nil {
		return 0, 0, errStore
	}
	return orderedID, totalPrice, nil
}

func (service *OrderService) getTotalPrice(ordersMenu []models.IdMenus) (int, error) {
	var totalPrice = 0
	for _, orderMenu := range ordersMenu {
		menu, err := service.GetMenu(orderMenu.IdMenu)
		if err != nil {
			return 0, err
		}
		totalPrice = totalPrice + menu.Price
	}

	return totalPrice, nil
}

func (service *OrderService) cook(menuReceipt []models.MenuIngredient) error {
	for _, ingredient := range menuReceipt {
		err := service.SubstractIngredientStock(ingredient.Name, ingredient.Quantity)
		if err != nil {
			return err
		}

	}

	return nil
}
