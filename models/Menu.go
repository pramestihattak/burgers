package models

type Menu struct {
	IdMenu int    `json:"id_menu"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}
