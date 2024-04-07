package main

import (
	"az/app/controllers"
	"az/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@example.com"
		u.PassWord = "testest"
		fmt.Println(u)

		u.CreateUser()
	*/
	/*
		u, _ := models.GetUser(1)

		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/
}
