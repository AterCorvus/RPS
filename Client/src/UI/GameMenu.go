package ui

import (
	"bufio"
	"fmt"
	"os"

	game "github.com/AterCorvus/RPS/CLient/src/Game"
	utils "github.com/AterCorvus/RPS/CLient/src/Utils"
	validators "github.com/AterCorvus/RPS/CLient/src/Validators"
)

func ShowPlayersOnline() {
	game.ShowPlayersOnline()
	Play()
}

func ChallengePlayer() {
	reader := bufio.NewReader(os.Stdin)

	oponentName, _ := utils.GetInput("Enter player name: ", reader)
	bet, _ := utils.GetInput("Enter bet: ", reader)
	validBet, err := validators.FundsValidator(bet, reader)
	if err != nil {
		ChallengePlayer()
	}
	move, _ := utils.GetInput("Choose your move 0 - Rock, 1 - Paper, 2 - Scissors, 4 - Back): ", reader)
	if move == "4" {
		ViewPendingChallenges()
	} else {
		validMove, moveErr := validators.MoveValidator(move, reader)
		if moveErr != nil {
			ChallengePlayer()
		}
		game.ChallengePlayer(oponentName, validBet, validMove)
	}
	Play()
}

func ViewPendingChallenges() {
	reader := bufio.NewReader(os.Stdin)

	game.ViewPendingChallenges()
	id, _ := utils.GetInput("Choose challenge to accept or decline, 0 - Back): ", reader)
	//TODO think about validation
	if id == "0" {
		Play()
	} else {
		validID, err := validators.IDValidator(id, reader)
		if err != nil {
			ViewPendingChallenges()
		}
		choise, _ := utils.GetInput("Choose 1 - To Accept, 2 - To Decline, 3 - Back): ", reader)
		switch choise {
		case "1":
			move, _ := utils.GetInput("Choose your move 0 - Rock, 1 - Paper, 2 - Scissors, 4 - Back): ", reader)
			if move == "4" {
				Play()
			} else {
				validMove, moveErr := validators.MoveValidator(move, reader)
				if moveErr != nil {
					ViewPendingChallenges()
				}
				game.AcceptChallenge(validID, validMove)
			}
		case "2":
			game.DeclineChallenge(validID)
		case "3":
			ViewPendingChallenges()
		default:
			fmt.Println("Invalid option")
			ViewPendingChallenges()
		}
	}
	Play()
}

func Play() {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := utils.GetInput("Choose option (1 - Show players online, 2 - Challenge player 3 - View pending Challenges 4 - Back): ", reader)
	switch opt {
	case "1":
		ShowPlayersOnline()
	case "2":
		ChallengePlayer()
	case "3":
		ViewPendingChallenges()
	case "4":
		MainMenu()
	default:
		fmt.Println("Invalid option")
		Play()
	}
}
