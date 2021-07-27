package Controllers

import (
	"fmt"
	"net/http"
	"store-api/Service"
	"store-api/Models"
	"github.com/gin-gonic/gin"
)

func PlaceOrderController(c *gin.Context) {

	var order Models.Orders
	c.BindJSON(&order)
	order.Status ="Order Successful !!"
	err := Service.PlaceOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		c.JSON(http.StatusOK, order)
	}

}

func GetOrderByIDController(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Orders
	c.BindJSON(&order)
	err := Service.GetOrderByID(&order,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func CreateUserController(c *gin.Context){
	var user Models.User
	c.BindJSON(&user)
	err := Service.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}