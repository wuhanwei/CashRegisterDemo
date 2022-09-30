package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "CashRegisterDemo/models"
	// "CashRegisterDemo/services"
)

type structCashRegister struct {
	NestedStruct structRedeemPoints
	Order_price  string `form:"order_price" binding:"required"`
	Member_level string `form:"member_level" binding:"required"`
}

type structRedeemPoints struct {
	Field_points string `form:"points" json:"points" binding:"required"`
	Field_coins  string `form:"coins" json:"coins" binding:"required"`
}

type structResponseOrderData struct {
	redeems           structRedeemPoints
	Field_order_price string `json:"final_order_price"`
}

type structResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	// Data    string `json:"data"`
	Data *structResponseOrderData `json:"data"`
}

func CashRegisterController(context *gin.Context) {
	var url_param structCashRegister
	var response_data structResponse

	fmt.Println(url_param)
	err := context.ShouldBindQuery(&url_param)
	fmt.Println(err)
	if err != nil {
		response_data = structResponse{
			Status:  "error",
			Message: err.Error(),
		}
	} else {
		response_order := &structResponseOrderData{
			Field_order_price: "0",
		}
		response_data = structResponse{
			Status:  "OK",
			Message: "",
			Data:    response_order,
		}
	}
	context.JSON(http.StatusOK, response_data)
}
