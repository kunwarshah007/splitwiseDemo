package controller

import (
	"fmt"
	"splitwiseDemo/models"
	"splitwiseDemo/service"
)

type UserController struct {
	UserService service.IUserService
}

func (c *UserController) CreateUser(name string, mob string, email string) (models.User, error) {
	user, err := c.UserService.CreateUser(name, mob, email)
	if err != nil {
		fmt.Println("Error while creating the user")
		return models.User{}, err
	}
	fmt.Printf("User Created successfully with UserID: %v\n", user.UserId)
	return user, nil
}
