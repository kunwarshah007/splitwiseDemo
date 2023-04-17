package service

import "splitwiseDemo/models"

type IUserService interface {
	CreateUser(string, string, string) (models.User, error)
}
