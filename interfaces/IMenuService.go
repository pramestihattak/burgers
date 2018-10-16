package interfaces

import "github.com/opam22/burgers/models"

type IMenuService interface {
	GetMenus() ([]models.Menu, error)
}
