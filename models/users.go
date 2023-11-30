package models

type Address struct {
	Country string	`json:"country"`
	City string	`json:"city"`
	PinCode int	`json:"pincode"`
}

type User struct {
	ID string `json: "id"`
	Name string	`json:"name"`
	Age int	`json:"age"`
	Gender string	`json:"gender"`
	Address Address	`json:"address"`
}