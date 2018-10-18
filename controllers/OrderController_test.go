package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/opam22/burgers/interfaces/mocks"
	"github.com/opam22/burgers/models"
	"github.com/opam22/burgers/payloads"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {

	order := models.Order{
		BuyerName:  "Pram",
		IdDiscount: "GOPHER",
		TotalPrice: 20000,
		Orders:     []models.IdMenus{models.IdMenus{IdMenu: 1}},
		Status:     "Your order is success",
		OrderCode:  1,
	}

	orderService := new(mocks.IOrderService)
	orderService.On("MakeOrder", order).Return(order, nil)

	orderController := OrderController{orderService}

	jsonOrder, _ := json.Marshal(order)

	// call the code we are testing
	req := httptest.NewRequest("POST", "/order", bytes.NewBuffer(jsonOrder))
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Post("/order", orderController.Order)
	r.ServeHTTP(w, req)

	expectedResult := payloads.SuccessOrder{Code: http.StatusOK, Response: order}

	actualResult := payloads.SuccessOrder{Response: models.Order{}}

	json.NewDecoder(w.Body).Decode(&actualResult)

	fmt.Printf("%v\n", w.Body)

	assert.Equal(t, expectedResult, actualResult)
}
