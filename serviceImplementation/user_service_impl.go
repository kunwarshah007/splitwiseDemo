package serviceImplementation

import (
	"splitwiseDemo/models"
	"strconv"
)

type Users struct {
	UsersDataset []models.User
}

var userId int

func UserInitializer() {
	userId = 1
	return
}

func (u *Users) CreateUser(name string, mobileNumber string, emailId string) (models.User, error) {
	user := models.User{
		UserId:       "U" + strconv.Itoa(userId),
		Name:         name,
		MobileNumber: mobileNumber,
		Email:        emailId,
	}
	u.UsersDataset = append(u.UsersDataset, user)
	userId = userId + 1
	return user, nil
}
