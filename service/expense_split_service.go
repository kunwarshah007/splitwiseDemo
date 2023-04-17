package service

import "splitwiseDemo/models"

type IExpenseSplitService interface {
	SplitEqual(expense models.Expense, group *models.Group) error
	SplitExact(expense models.Expense, group *models.Group) error
	SplitPercentage(expense models.Expense, group *models.Group) error
}
