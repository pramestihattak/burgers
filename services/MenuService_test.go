package services

import (
	"testing"

	"github.com/opam22/burgers/interfaces/mocks"
	"github.com/opam22/burgers/models"
	"github.com/stretchr/testify/assert"
)

func TestGetMenu(t *testing.T) {
	menuRepository := new(mocks.IMenuRepository)

	menu := models.Menu{}
	menu.IdMenu = 1
	menu.Name = "Burger Blenger"
	menu.Price = 30000

	menuRepository.On("FetchMenu", 1).Return(menu, nil)

	menuService := MenuService{menuRepository}

	expectedResult := menu

	actualResult, _ := menuService.FetchMenu(1)

	assert.Equal(t, expectedResult, actualResult)
}

func TestGetMenus(t *testing.T) {
	menuRepository := new(mocks.IMenuRepository)

	var menus = []models.Menu{}
	menus = append(menus, models.Menu{IdMenu: 1, Name: "Burger Blenger", Price: 22000})
	menus = append(menus, models.Menu{IdMenu: 2, Name: "Burger King", Price: 18000})
	menus = append(menus, models.Menu{IdMenu: 1, Name: "Cars Jr", Price: 32000})

	menuRepository.On("FetchAllMenu").Return(menus, nil)

	menuService := MenuService{menuRepository}

	expectedResult := menus

	actualResult, _ := menuService.FetchAllMenu()

	assert.Equal(t, expectedResult, actualResult)
}

func TestGetMenuReceipt(t *testing.T) {
	menuRepository := new(mocks.IMenuRepository)

	var menuReceipt = []models.MenuIngredient{}
	menuReceipt = append(menuReceipt, models.MenuIngredient{Name: "Beef", Quantity: 3})
	menuReceipt = append(menuReceipt, models.MenuIngredient{Name: "Lettuce", Quantity: 3})

	menuRepository.On("FetchMenuIngredients", 1).Return(menuReceipt, nil)

	menuService := MenuService{menuRepository}

	expectedResult := menuReceipt

	actualResult, _ := menuService.FetchMenuIngredients(1)

	assert.Equal(t, expectedResult, actualResult)
}
