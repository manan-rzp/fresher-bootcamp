package Service

import (
	"fmt"
	"store-api/Config"
	"store-api/Models"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)
func CoolDown (order *Models.Orders) bool{

	lastOrderTime := order.CreatedAt
	nextOrderTime := lastOrderTime.Add(5*time.Minute)
	currentTime := time.Now()
	return currentTime.Before(nextOrderTime)

}

func PlaceOrder(order *Models.Orders) (err error) {
	mu.Lock()

	var order1 Models.Orders

	if err = Config.DB.Where("user_id = ?", order.UserID).Last(&order1).Error ; err!=nil {
		fmt.Println("ds")
	}

	if CoolDown(&order1){
		mu.Unlock()
		order.Status = "Failed"
		return nil
	}

	var product Models.Product
	//fmt.Println(strconv.FormatUint(uint64(order.ProductID),10))
	err = GetProductbyID(&product, strconv.FormatUint(uint64(order.ProductID), 10))
	if err != nil {
		mu.Unlock()
		panic(err)
	}

	fmt.Println(product.Quantity)
	product.Quantity-=1
	Config.DB.Save(&product)

	if err != nil {
		mu.Unlock()
		panic(err)
	}

	if err = Config.DB.Set("gorm:auto_preload",true).Create(order).Error; err != nil {
		mu.Unlock()
		return err
	}
	mu.Unlock()
	return nil
}

func GetOrderByID(order *Models.Orders,id string) (err error) {
	if err = Config.DB.Set("gorm:auto_preload",true).Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *Models.User)(err error){
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}