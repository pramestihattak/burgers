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

	r := chi.NewRouter()
	r.HandleFunc("/menus", menuController.AllMenu)
	r.HandleFunc("/menu/{id}/receipt", menuController.MenuReceipt)

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
