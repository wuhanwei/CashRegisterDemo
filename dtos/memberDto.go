package dtos

type StructMember struct {
	Field_level       string `json:"level"`
	Field_level_bonus string `json:"bonus"`
	Field_points      int    `json:"poins"`
}

type StructCashRegister struct {
	Redeem_points string `form:"redeem_points" binding:"required"`
	Order_price   string `form:"order_price" binding:"required"`
	Name          string `form:"name" binding:"required"`
	Sale          string `form:"sale"`
}
