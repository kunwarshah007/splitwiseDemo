package controller

import (
	"fmt"
	"splitwiseDemo/models"
	"splitwiseDemo/service"
)

type ExpenseController struct {
	ExpenseService service.IExpenseSplitService
}

func (c *ExpenseController) SplitExpense(expensse models.Expense, group *models.Group) {
	err := c.ExpenseService.SplitEqual(expensse, group)
	if err != nil {
		fmt.Println("Error while creating the Group")
	}
}
