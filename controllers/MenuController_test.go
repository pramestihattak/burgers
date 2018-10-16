package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/opam22/burgers/interfaces/mocks"
	"github.com/opam22/burgers/models"
	"github.com/opam22/burgers/payloads"
	"github.com/stretchr/testify/assert"
)

func TestAllMenu(t *testing.T) {
	menuService := new(mocks.IMenuService)

	var menus = []models.Menu{}
	menus = append(menus, models.Menu{IdMenu: 1, Name: "Burger Blenger", Price: 22000})
	menus = append(menus, models.Menu{IdMenu: 2, Name: "Burger King", Price: 18000})
	menus = append(menus, models.Menu{IdMenu: 1, Name: "Cars Jr", Price: 32000})

	menuService.On("GetMenus").Return(menus, nil)

	menuController := MenuController{menuService}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:6969/menus", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/menus", menuController.AllMenu)

	r.ServeHTTP(w, req)

	expectedResult := payloads.SuccessGetAllMenu{Code: http.StatusOK, Response: menus}

	actualResult := payloads.SuccessGetAllMenu{}

	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMenuReceipt(t *testing.T) {
	menuService := new(mocks.IMenuService)
	menu := models.Menu{IdMenu: 1, Name: "Burger Super", Price: 22000}

	var menuReceipt = []models.MenuIngredient{}
	menuReceipt = append(menuReceipt, models.MenuIngredient{Name: "Beef", Quantity: 3})
	menuReceipt = append(menuReceipt, models.MenuIngredient{Name: "Lettuce", Quantity: 3})

	menuService.On("GetMenuReceipt", 1).Return(menuReceipt, nil)
	menuService.On("GetMenu", 1).Return(menu, nil)

	menuController := MenuController{menuService}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:6969/menus/1/receipt", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/menus/{id}/receipt", menuController.MenuReceipt)

	r.ServeHTTP(w, req)

	expectedResult := payloads.SuccessGetMenuReceipt{Code: http.StatusOK, Response: payloads.MenuReceipt{Menu: menu, Receipt: menuReceipt}}

	actualResult := payloads.SuccessGetMenuReceipt{Response: payloads.MenuReceipt{}}

	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, expectedResult, actualResult)

}
