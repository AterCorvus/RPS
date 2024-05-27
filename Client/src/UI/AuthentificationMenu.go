package ui

import (
	"bufio"
	"fmt"
	"os"

	authentification "github.com/AterCorvus/RPS/CLient/src/Authentification"
	utils "github.com/AterCorvus/RPS/CLient/src/Utils"
)

func login(reader *bufio.Reader) {
	username, _ := utils.GetInput("Enter username: ", reader)
	password, _ := utils.GetInput("Enter password: ", reader)

	result := authentification.Login(username, password)
	if result {
		MainMenu()
	} else {
		LoginOrRegisterUser()
	}
}

func registration(reader *bufio.Reader) {
	username, _ := utils.GetInput("Enter username: ", reader)
	password, _ := utils.GetInput("Enter password: ", reader)

	result := authentification.Registration(username, password)
	if result {
		MainMenu()
	} else {
		LoginOrRegisterUser()
	}
}

func LoginOrRegisterUser() {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := utils.GetInput("Choose option (1 - Login, 2 - Registrate): ", reader)
	switch opt {
	case "1":
		login(reader)
	case "2":
		registration(reader)
	default:
		fmt.Println("Invalid option")
		LoginOrRegisterUser()
	}
}
