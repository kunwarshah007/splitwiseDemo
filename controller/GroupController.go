package controller

import (
	"fmt"
	"splitwiseDemo/models"
	"splitwiseDemo/service"
)

type GroupController struct {
	GroupService service.IGroupService
}

func (c *GroupController) CreateGroup(name string, users []models.User) (models.Group, error) {
	group, err := c.GroupService.CreateGroup(name, users)
	if err != nil {
		fmt.Println("Error while creating the Group")
		return models.Group{}, err
	}
	fmt.Printf("Group Created successfully with GroupID: %v\n", group.GroupId)
	return group, nil
}
