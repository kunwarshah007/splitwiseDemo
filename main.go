package main

import (
	"fmt"
	"splitwiseDemo/controller"
	"splitwiseDemo/models"
	"splitwiseDemo/serviceImplementation"
)

func main() {
	serviceImplementation.UserInitializer()
	serviceImplementation.GroupInitializer()
	usersDataset := &serviceImplementation.Users{UsersDataset: []models.User{}}
	groupsDataset := &serviceImplementation.Groups{GroupsDataset: []models.Group{}, Transactions: []models.Expense{}}
	userController := controller.UserController{UserService: usersDataset}
	groupController := controller.GroupController{GroupService: groupsDataset}
	expenseController := controller.ExpenseController{ExpenseService: groupsDataset}

	var users []models.User
	user1, err := userController.CreateUser("Harshit", "7905009505", "harshit@gmail.com")
	users = append(users, user1)
	user2, err := userController.CreateUser("Aditya", "0000000000", "aditya@gmail.com")
	users = append(users, user2)
	user3, err := userController.CreateUser("Shahid", "1111111111", "shahid@gmail.com")
	users = append(users, user3)

	if err != nil {
		fmt.Println("Error Occered")
	}
	group1, _ := groupController.CreateGroup("Friends", users)
	fmt.Printf("Value to be paid by User1 in group1 is %v", group1.GroupTransactionDetails[group1.GroupId+user1.UserId])
	cond := true
	for cond {
		fmt.Println("Enter your Choice:")
		fmt.Println("A: for making an Expense")
		fmt.Println("B: for checking all the Expense of a User in a Group")
		fmt.Println("C. for checking Expenses by all User in a Group")
		fmt.Println("Others for exit")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "A":
			fmt.Println("Enter in the Format : [u1 1000 4 u1 u2 u3 u4 EQUAL]")
			var paidBy, type1, userId string
			var paidTo []string
			var exactShare []float64
			var amount, share float64
			var numberOfUsers int
			fmt.Scanln(&paidBy)
			fmt.Scanln(&amount)
			fmt.Scanln(&numberOfUsers)
			for i := 1; i <= numberOfUsers; i++ {
				fmt.Scanln(&userId)
				paidTo = append(paidTo, userId)
			}
			fmt.Scanln(&type1)
			if type1 != "EQUAL" && type1 != "EXACT" && type1 != "PERCENT" {
				fmt.Println("Invalid Type")
				break
			}
			if type1 == "EXACT" || type1 == "PERCENT" {
				for i := 1; i <= numberOfUsers; i++ {
					fmt.Scanln(&share)
					exactShare = append(exactShare, share)
				}
			}
			var expense = models.Expense{
				GroupId:             group1.GroupId,
				PaidByUserId:        paidBy,
				Amount:              amount,
				UserIdsThatOwsShare: paidTo,
				Type:                type1,
				Share:               exactShare,
			}
			fmt.Println(expense)
			expenseController.SplitExpense(expense, &group1)
			fmt.Printf("%v", group1.GroupTransactionDetails)
			break
		case "B":
			break
		case "C":
			break
		default:
			cond = false
			break
		}
	}

}
