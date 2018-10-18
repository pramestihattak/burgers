package main

import (
	"sync"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	menuController := ServiceContainer().InjectMenuController()
	orderController := ServiceContainer().InjectOrderController()

	r := chi.NewRouter()
	r.Get("/menus", menuController.AllMenu)
	r.Get("/menu/{id}/receipt", menuController.MenuReceipt)

	r.Post("/order", orderController.Order)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
