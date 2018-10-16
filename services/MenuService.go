package services

import (
	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/models"
)

type MenuService struct {
	interfaces.IMenuRepository
}

func (service *MenuService) GetMenus() ([]models.Menu, error) {

	menus, err := service.FetchAllMenu()
	if err != nil {
		//Handle error
	}

	return menus, nil
}
