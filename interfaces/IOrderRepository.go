package interfaces

import "github.com/opam22/burgers/models"

type IOrderRepository interface {
	GetMenuIngredients(idMenu int) ([]models.MenuIngredient, error)
	GetIngredientStock(ingredient string) (int, error)
	GetIngredients() ([]models.Ingredient, error)
	SubstractIngredientStock(ingredient string, quantity int) error
	GetMenu(idMenu int) (models.Menu, error)
	CreateBill(order models.Order) (int64, error)
	StoreOrderedMenu(orderedID int64, orderedMenu []models.IdMenus) error
}
