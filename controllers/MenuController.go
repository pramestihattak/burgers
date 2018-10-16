package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/payloads"
)

type MenuController struct {
	interfaces.IMenuService
}

func (controller *MenuController) AllMenu(res http.ResponseWriter, req *http.Request) {

	menus, err := controller.GetMenus()
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(payloads.SuccessGetAllMenu{Code: http.StatusOK, Response: menus})
}

func (controller *MenuController) MenuReceipt(res http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	idMenu, errAtoi := strconv.Atoi(id)
	if errAtoi != nil {

	}

	menuReceipt, err := controller.GetMenuReceipt(idMenu)
	if err != nil {
	}

	menu, err := controller.GetMenu(idMenu)
	if err != nil {
	}

	json.NewEncoder(res).Encode(payloads.SuccessGetMenuReceipt{Code: http.StatusOK, Response: payloads.MenuReceipt{Menu: menu, Receipt: menuReceipt}})
}
