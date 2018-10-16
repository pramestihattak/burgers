package interfaces

import "github.com/opam22/burgers/models"

type IMenuRepository interface {
	FetchAllMenu() ([]models.Menu, error)
}
