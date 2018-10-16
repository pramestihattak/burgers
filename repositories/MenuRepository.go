package repositories

import (
	"fmt"

	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/models"
)

type MenuRepository struct {
	interfaces.IDBHandler
}

func (repository *MenuRepository) FetchMenu(idMenu int) (models.Menu, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM menus WHERE id_menu = '%d'", idMenu))
	if err != nil {
		return models.Menu{}, err
	}

	var menu models.Menu

	row.Next()
	row.Scan(&menu.IdMenu, &menu.Name, &menu.Price)

	return menu, nil
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
		row.Scan(&menu.IdMenu, &menu.Name, &menu.Price)

		menus = append(menus, menu)
	}

	return menus, nil
}

func (repository *MenuRepository) FetchMenuIngredients(idMenu int) ([]models.MenuIngredient, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT ingredients.name, menu_ingredient.quantity FROM menu_ingredient JOIN ingredients ON menu_ingredient.id_ingredient = ingredients.id_ingredient where id_menu = '%d'", idMenu))

	if err != nil {
		return []models.MenuIngredient{}, err
	}

	var (
		menuIngredient  models.MenuIngredient
		menuIngredients []models.MenuIngredient
	)

	for row.Next() {
		row.Scan(&menuIngredient.Name, &menuIngredient.Quantity)
		menuIngredients = append(menuIngredients, menuIngredient)
	}

	return menuIngredients, nil
}
