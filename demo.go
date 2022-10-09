package main

import (
	"CashRegisterDemo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/CashRegister", controllers.CashRegisterController)
	route.Run()
}
