package serviceImplementation

import (
	"splitwiseDemo/models"
	"strconv"
)

type Groups struct {
	GroupsDataset []models.Group
	Transactions  []models.Expense
}

var groupId int

func GroupInitializer() {
	groupId = 1
	return
}

func (g *Groups) CreateGroup(name string, users []models.User) (models.Group, error) {
	group := models.Group{
		GroupId: "G" + strconv.Itoa(groupId),
		Users:   users,
		Name:    name,
	}
	g.GroupsDataset = append(g.GroupsDataset, group)
	g.GroupTransactionInitializer(&group)
	return group, nil
}

func (g *Groups) GroupTransactionInitializer(group *models.Group) {
	expenseShareMapWithOtherUsers := make(map[string]map[string]float64) // string here is combination of group_id + user_id
	groupID := group.GroupId
	for _, user := range group.Users {
		key := groupID + user.UserId
		groupTransactionDetail := make(map[string]float64)
		for _, otherUsers := range group.Users {

			if otherUsers.UserId == user.UserId {
				continue
			}
			groupTransactionDetail[otherUsers.UserId] = 0.0
		}
		expenseShareMapWithOtherUsers[key] = groupTransactionDetail
		//groupTransactionDetails[key] = 0
	}
	group.GroupTransactionDetails = expenseShareMapWithOtherUsers
	return
}
