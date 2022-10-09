package dtos

type StructResponseOrderData struct {
	Field_order_price     string `json:"final_order_price"`
	Field_level_bonus     string `json:"member_level_bonus"`
	Field_discount_points string `json:"discount_points"`
	Field_discount_coins  string `json:"discount_coins"`
}

type StructResponse struct {
	Status     string                   `json:"status"`
	Message    string                   `json:"message"`
	StatusCode string                   `json:"code"`
	Data       *StructResponseOrderData `json:"data"`
}
