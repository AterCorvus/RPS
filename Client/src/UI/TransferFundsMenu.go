package ui

import (
	"bufio"
	"fmt"
	"os"

	utils "github.com/AterCorvus/RPS/CLient/src/Utils"
	validators "github.com/AterCorvus/RPS/CLient/src/Validators"
	funds "github.com/AterCorvus/RPS/CLient/src/Funds"
)

func AddMoney() {
	reader := bufio.NewReader(os.Stdin)

	amount, _ := utils.GetInput("Enter amount: ", reader)
	valid_amount, err := validators.FundsValidator(amount, reader)
	if err != nil {
		AddMoney()
	}
	funds.AddMoney(valid_amount)
	MainMenu()
}

func WidthdrawMoney() {
	reader := bufio.NewReader(os.Stdin)

	amount, _ := utils.GetInput("Enter amount: ", reader)
	valid_amount, err := validators.FundsValidator(amount, reader)
	if err != nil {
		WidthdrawMoney()
	}
	funds.WidthdrawMoney(valid_amount)
	MainMenu()
}

func TransferFunds() {
	reader := bufio.NewReader(os.Stdin)
	funds.GetMyFunds()

	opt, _ := utils.GetInput("Choose option (1 -AddMoney to your account, 2 - WidthdrawMoney 3 - Back): ", reader)
	switch opt {
	case "1":
		AddMoney()
	case "2":
		WidthdrawMoney()
	case "3":
		MainMenu()
	default:
		fmt.Println("Invalid option")
		TransferFunds()
	}
}
