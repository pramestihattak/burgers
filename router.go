package main

import (
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type IRouter interface {
	Init() *chi.Mux
}

type router struct{}

func (router *router) Init() *chi.Mux {

	menuController := ServiceContainer().InjectMenuController()
	orderController := ServiceContainer().InjectOrderController()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/menus", menuController.AllMenu)
	r.Get("/menu/{id}/receipt", menuController.MenuReceipt)

	r.Post("/order", orderController.Order)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func Router() IRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
