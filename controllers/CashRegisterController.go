package controllers

import (
	"CashRegisterDemo/dtos"
	"CashRegisterDemo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CashRegisterController(context *gin.Context) {
	var url_param dtos.StructCashRegister

	if err := context.ShouldBindQuery(&url_param); err != nil {
		context.JSON(http.StatusOK, services.Response_error(err, "4101"))
	} else {
		if response_order, err, status_code := services.GetRedeemPrice(&url_param); err != nil {
			context.JSON(http.StatusOK, services.Response_error(err, status_code))
		} else {
			context.JSON(http.StatusOK, services.Response_success(*response_order))
		}
	}
}
