package service

import "splitwiseDemo/models"

type IGroupService interface {
	CreateGroup(string, []models.User) (models.Group, error)
	GroupTransactionInitializer(*models.Group)
}
