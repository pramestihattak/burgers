package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/opam22/burgers/interfaces"
	"github.com/opam22/burgers/payloads"
)

type MenuController struct {
	interfaces.IMenuService
}

func (controller *MenuController) GetAllMenu(res http.ResponseWriter, req *http.Request) {

	menus, err := controller.GetMenus()
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(payloads.SuccessGetAllMenu{Code: http.StatusOK, Response: menus})
}
