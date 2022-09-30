package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"CashRegisterDemo/controllers"
)

func main() {
	route := gin.Default()
	route.GET("/CashRegister", controllers.CashRegisterController)
	route.Run()
}
