package interfaces

import "github.com/opam22/burgers/models"

type IMenuService interface {
	GetMenu(idMenu int) (models.Menu, error)
	GetMenus() ([]models.Menu, error)
	GetMenuReceipt(idMenu int) ([]models.MenuIngredient, error)
}
