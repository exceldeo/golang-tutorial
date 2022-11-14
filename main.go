package main

import (
	"fmt"
	"park-billing/park"
)

func main() {
	fmt.Println("Hello World")
	price1 := park.CalculateParkingBill(park.Car, 48)
	price2 := park.CalculateParkingBill(park.Motorcycle, 8)
	price3 := park.CalculateParkingBill(park.Bicycle, 8)
	price4 := park.CalculateParkingBill(park.Car, 48)
	fmt.Println("price", price1)
	fmt.Println("price", price2)
	fmt.Println("price", price3)
	fmt.Println("price", price4)
}
