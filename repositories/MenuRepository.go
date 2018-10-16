package repositories

import (
	"fmt"

	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/models"
)

type MenuRepository struct {
	interfaces.IDBHandler
}

func (repository *MenuRepository) FetchAllMenu() ([]models.Menu, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM menus"))
	if err != nil {
		return []models.Menu{}, err
	}

	var (
		menu  models.Menu
		menus []models.Menu
	)

	for row.Next() {
		row.Scan(&menu.Id, &menu.Name, &menu.Price)

		menus = append(menus, menu)
	}

	return menus, nil
}
