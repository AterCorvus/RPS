package ui

import (
	"bufio"
	"fmt"
	"os"

	authentification "github.com/AterCorvus/RPS/CLient/src/Authentification"
	utils "github.com/AterCorvus/RPS/CLient/src/Utils"
)

func MainMenu() {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := utils.GetInput("Choose option (1 -Transfer funds, 2 - Play 3 - Logout): ", reader)
	switch opt {
	case "1":
		TransferFunds()
	case "2":
		Play()
	case "3":
		authentification.Logout()
		LoginOrRegisterUser()
	default:
		fmt.Println("Invalid option")
		MainMenu()
	}
}
