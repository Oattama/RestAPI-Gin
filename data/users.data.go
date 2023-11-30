package data

import (
	m "gin/restAPI/models"
)	

var Users = [] m.User {
	{
		ID: "1",
		Name: "Thomas",
		Age: 32,
		Gender: "Male",
		Address: m.Address{
			Country: "England",
			City: "London",
			PinCode: 12358,
		},
	},
	{
		ID: "2",
		Name: "Toji",
		Age: 28,
		Gender: "Male",
		Address: m.Address{
			Country: "Japan",
			City: "Hokkaido",
			PinCode: 15723,
		},
	},
}