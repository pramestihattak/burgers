package models

type Order struct {
	BuyerName  string    `json:"buyer_name"`
	IdDiscount string    `json:"id_discount"`
	TotalPrice int       `json:"total_price"`
	Orders     []IdMenus `json:"orders"`
	Status     string    `json:"status"`
	OrderCode  int64     `json:"order_code"`
}

type IdMenus struct {
	IdMenu int `json:"id_menu,int"`
}
