package models

type Ingredient struct {
	IdIngredient int    `json:"id_ingredient"`
	Name         string `json:"name"`
	Stock        int    `json:"stock"`
}
