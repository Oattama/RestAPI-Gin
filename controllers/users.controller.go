package controllers

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	cdb "gin/restAPI/config"
	m "gin/restAPI/models"
	//d "gin/restAPI/data"
)

var db = cdb.Connection()

func GetUsers(c *gin.Context) {
	var result m.User

	query := `SELECT * FROM user WHERE ID = ?`

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	data, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}
	
	defer data.Close()

	for data.Next() {
		var user m.User

		err := data.Scan(&user.ID, 
			&user.Name, 
			&user.Age, 
			&user.Gender, 
			&user.Address.Country, 
			&user.Address.City, 
			&user.Address.PinCode)

		if err != nil {
			fmt.Println(err)
		}
		result = user
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func ListUsers(c *gin.Context) {
	var userList []m.User

	query := `SELECT * FROM user`

	data,err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	for data.Next() {
		var user m.User

		err := data.Scan(&user.ID, 
			&user.Name, 
			&user.Age, 
			&user.Gender, 
			&user.Address.Country, 
			&user.Address.City, 
			&user.Address.PinCode)

		if err != nil {
			fmt.Println(err)
		}
		userList = append(userList, user)
	}

	c.JSON(200, gin.H{
		"result": userList,
	})
}

func EditUsers(c *gin.Context) {
	var user m.User

	user.Name = c.PostForm("name")
	user.Age, _ = strconv.Atoi(c.PostForm("age"))
	user.Gender = c.PostForm("gender")
	user.Address.Country = c.PostForm("country")
	user.Address.City = c.PostForm("city")
	user.Address.PinCode, _ = strconv.Atoi(c.PostForm("pincode"))

	queryId := c.Query("id")

	query := `UPDATE user SET Name=?, Age=?, Gender=?, Country=?, City=?, PinCode=? WHERE ID = ?`

	data, err := db.Query(query, user.Name, user.Age, user.Gender, user.Address.Country, user.Address.City, user.Address.PinCode, queryId)
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	c.JSON(200, gin.H{
		"result": "update successful",
	})
	
}

func AddUsers(c *gin.Context) {
	/* var user m.User

	user.ID = c.PostForm("id")
	user.Name = c.PostForm("name")
	user.Age, _ = strconv.Atoi(c.PostForm("age"))
	user.Gender = c.PostForm("gender")
	user.Address.Country = c.PostForm("country")
	user.Address.City = c.PostForm("city")
	user.Address.PinCode, _ = strconv.Atoi(c.PostForm("pincode"))

	newuser := user

	d.Users = append(d.Users, newuser)
	c.JSON(200, gin.H{
		"result": d.Users,
	}) */

	var user m.User

	user.Name = c.PostForm("name")
	user.Age, _ = strconv.Atoi(c.PostForm("age"))
	user.Gender = c.PostForm("gender")
	user.Address.Country = c.PostForm("country")
	user.Address.City = c.PostForm("city")
	user.Address.PinCode, _ = strconv.Atoi(c.PostForm("pincode"))

	query := `INSERT INTO user (Name, Age, Gender, Country, City, PinCode) VALUES (?,?,?,?,?,?)`

	data, err := db.Query(query, user.Name, user.Age, user.Gender, user.Address.Country, user.Address.City, user.Address.PinCode)
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	c.JSON(200, gin.H{
		"result": "add user successful",
	})
}