package main

import (
	"fmt"

	"github.com/danubiobwm/pfa-go/internal/order/entity"
)

func main() {
	order, err := entity.NewOrder("123", 10, 2)
	if err != nil {
		panic(err)
	}
	err = order.CalculteFinalPrice()
	if err != nil {
		panic(err)
	}

	fmt.Printf("The final price is: %f", order.FinalPrice)
}
