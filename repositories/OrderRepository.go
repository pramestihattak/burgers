package repositories

import (
	"fmt"

	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/models"
)

type OrderRepository struct {
	interfaces.IDBHandler
}

func (repository *OrderRepository) GetMenuIngredients(idMenu int) ([]models.MenuIngredient, error) {

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

func (repository *OrderRepository) GetIngredientStock(ingredient string) (int, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT stock FROM ingredients where name = '%s'", ingredient))

	if err != nil {
		return 0, err
	}

	var stock int

	row.Next()
	row.Scan(&stock)

	return stock, nil
}

func (repository *OrderRepository) GetIngredients() ([]models.Ingredient, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM ingredients"))

	if err != nil {
		return []models.Ingredient{}, err
	}

	var (
		ingredient  models.Ingredient
		ingredients []models.Ingredient
	)

	for row.Next() {
		row.Scan(&ingredient.IdIngredient, &ingredient.Name, &ingredient.Stock)
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (repository *OrderRepository) SubstractIngredientStock(ingredient string, quantity int) error {

	_, err := repository.Execute(fmt.Sprintf("UPDATE ingredients SET stock = stock - %d WHERE name = '%s'", quantity, ingredient))
	if err != nil {
		return err
	}
	return nil
}

func (repository *OrderRepository) GetMenu(idMenu int) (models.Menu, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM menus WHERE id_menu = %d", idMenu))

	if err != nil {
		return models.Menu{}, err
	}

	var menu models.Menu

	row.Next()
	row.Scan(&menu.IdMenu, &menu.Name, &menu.Price)

	return menu, nil
}

func (repository *OrderRepository) CreateBill(order models.Order) (int64, error) {

	res, err := repository.Execute(fmt.Sprintf("INSERT INTO orders(buyer_name, total_price, id_discount) VALUES ('%s', %d, '%s')", order.BuyerName, order.TotalPrice, order.IdDiscount))

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		println("Error:", err.Error())
	}
	return id, nil
}

func (repository *OrderRepository) StoreOrderedMenu(orderedID int64, orderedMenu []models.IdMenus) error {

	for _, menu := range orderedMenu {
		_, err := repository.Execute(fmt.Sprintf("INSERT INTO order_menus(id_order, id_menu) VALUES (%d, %d)", orderedID, menu.IdMenu))
		if err != nil {
			return err
		}
	}

	return nil
}
