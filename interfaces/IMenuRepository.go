package interfaces

import "github.com/opam22/burgers/models"

type IMenuRepository interface {
	FetchMenu(idMenu int) (models.Menu, error)
	FetchAllMenu() ([]models.Menu, error)
	FetchMenuIngredients(idMenu int) ([]models.MenuIngredient, error)
}
