package services

import (
	"testing"

	"github.com/opam22/burgers/interfaces/mocks"
	"github.com/opam22/burgers/models"
	"github.com/stretchr/testify/assert"
)

func TestMakeOrder(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)

	menuReceipt := []models.MenuIngredient{
		models.MenuIngredient{Name: "beef", Quantity: 2},
		models.MenuIngredient{Name: "lettuce", Quantity: 2},
	}

	orderRepository.On("GetMenuIngredients", 1).Return(menuReceipt, nil)

	orderService := OrderService{orderRepository}

	expectedResult := menuReceipt

	actualResult, _ := orderService.GetMenuIngredients(1)

	assert.Equal(t, expectedResult, actualResult)

}

func TestGetTotalPrice(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)

	menu := models.Menu{IdMenu: 1, Name: "Burger Belenger", Price: 20000}

	orderRepository.On("GetMenu", 1).Return(menu, nil)

	orderService := OrderService{orderRepository}

	expectedResult := menu

	actualResult, _ := orderService.GetMenu(1)

	assert.Equal(t, expectedResult, actualResult)

}

func TestGenerateBill(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)

	order := models.Order{
		BuyerName:  "Pram",
		IdDiscount: "GOPHER",
		TotalPrice: 20000,
		Status:     "Success",
		OrderCode:  1,
		Orders: []models.IdMenus{
			models.IdMenus{IdMenu: 1},
		},
	}

	var orderedID int64 = 1

	orderRepository.On("CreateBill", order).Return(orderedID, nil)
	orderRepository.On("StoreOrderedMenu", orderedID, order.Orders).Return(nil)

	orderService := OrderService{orderRepository}

	expectedResult := orderedID

	actualResult, _ := orderService.CreateBill(order)

	assert.Equal(t, expectedResult, actualResult)

	var expectedResultStore error

	actualResultStore := orderService.StoreOrderedMenu(orderedID, order.Orders)

	assert.Equal(t, expectedResultStore, actualResultStore)

}

func TestCook(t *testing.T) {
	orderRepository := new(mocks.IOrderRepository)

	ingredient := models.MenuIngredient{Name: "beef", Quantity: 4}

	orderRepository.On("SubstractIngredientStock", ingredient.Name, ingredient.Quantity).Return(nil)

	orderService := OrderService{orderRepository}

	var expectedResult error

	actualResult := orderService.SubstractIngredientStock(ingredient.Name, ingredient.Quantity)

	assert.Equal(t, expectedResult, actualResult)

}
