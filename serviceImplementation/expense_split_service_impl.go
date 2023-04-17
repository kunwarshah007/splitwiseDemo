package serviceImplementation

import (
	"errors"
	"splitwiseDemo/models"
)

func (g *Groups) SplitEqual(expense models.Expense, group *models.Group) error {

	err := checkIfUsersBelongToTheGroup(expense.PaidByUserId, expense.UserIdsThatOwsShare, group.GroupTransactionDetails, group.GroupId)
	if err != nil {
		return err
	}

	numberOfUser := float64(len(expense.UserIdsThatOwsShare))
	equalShare := float64(int((expense.Amount/numberOfUser)*100)) / 100
	// extract till first two decimal place
	equalShare = float64(int(equalShare*100)) / 100
	// if all share doesn't add up to amount
	if equalShare*numberOfUser < expense.Amount {
		userID := expense.UserIdsThatOwsShare[0]
		key := group.GroupId + userID
		group.GroupTransactionDetails[key][expense.PaidByUserId] = group.GroupTransactionDetails[key][expense.PaidByUserId] + 0.01
		key = group.GroupId + expense.PaidByUserId
		group.GroupTransactionDetails[key][userID] = group.GroupTransactionDetails[key][userID] - 0.01
	}
	for _, userID := range expense.UserIdsThatOwsShare {

		//Skipping the share to self
		if expense.PaidByUserId == userID {
			continue
		}

		key := group.GroupId + userID
		group.GroupTransactionDetails[key][expense.PaidByUserId] = group.GroupTransactionDetails[key][expense.PaidByUserId] + equalShare
		key = group.GroupId + expense.PaidByUserId
		group.GroupTransactionDetails[key][userID] = group.GroupTransactionDetails[key][userID] - equalShare
	}
	return nil
}

func (g *Groups) SplitExact(expense models.Expense, group *models.Group) error {

	err := checkIfUsersBelongToTheGroup(expense.PaidByUserId, expense.UserIdsThatOwsShare, group.GroupTransactionDetails, group.GroupId)
	if err != nil {
		return err
	}

	numberOfUser := len(expense.UserIdsThatOwsShare)
	//check if all share sums up to amount
	sum := 0.0
	for i := 0; i < numberOfUser; i++ {
		sum += expense.Share[i]
	}
	if sum != expense.Amount {
		return errors.New("exact share sum is not equal to amount")
	}
	for i, userID := range expense.UserIdsThatOwsShare {

		//Skipping the share to self
		if expense.PaidByUserId == userID {
			continue
		}

		key := group.GroupId + userID
		group.GroupTransactionDetails[key][expense.PaidByUserId] = group.GroupTransactionDetails[key][expense.PaidByUserId] + expense.Share[i]
		key = group.GroupId + expense.PaidByUserId
		group.GroupTransactionDetails[key][userID] = group.GroupTransactionDetails[key][userID] - expense.Share[i]
	}
	return nil
}

func (g *Groups) SplitPercentage(expense models.Expense, group *models.Group) error {

	err := checkIfUsersBelongToTheGroup(expense.PaidByUserId, expense.UserIdsThatOwsShare, group.GroupTransactionDetails, group.GroupId)
	if err != nil {
		return err
	}

	numberOfUser := len(expense.UserIdsThatOwsShare)
	//check if all share sums up to amount
	sum := 0.0
	for i := 0; i < numberOfUser; i++ {
		sum += expense.Share[i]
	}
	if sum != 100.0 {
		return errors.New("percentage sum is not equal to 100")
	}
	for i, userID := range expense.UserIdsThatOwsShare {

		//Skipping the share to self
		if expense.PaidByUserId == userID {
			continue
		}

		key := group.GroupId + userID
		share := float64(int(expense.Share[i]*expense.Amount)) / 100
		group.GroupTransactionDetails[key][expense.PaidByUserId] = group.GroupTransactionDetails[key][expense.PaidByUserId] + share
		key = group.GroupId + expense.PaidByUserId
		group.GroupTransactionDetails[key][userID] = group.GroupTransactionDetails[key][userID] - share
	}
	return nil
}

func checkIfUsersBelongToTheGroup(paidBy string, paidTo []string, groupTransactionDetails map[string]map[string]float64, groupId string) error {
	key := groupId + paidBy
	if _, userExistInGroup := groupTransactionDetails[key]; !userExistInGroup {
		return errors.New("the user who is paying the bill does not belongs to the group")
	}
	for _, userID := range paidTo {
		key = groupId + userID
		if _, userExistInGroup := groupTransactionDetails[key]; !userExistInGroup {
			return errors.New("the user who is paying the bill does not belongs to the group")
		}
	}
	return nil
}
