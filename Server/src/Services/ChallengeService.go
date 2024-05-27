package services

import (
	"errors"

	controllers "github.com/AterCorvus/RPS/Server/src/Controllers"
	gamelogic "github.com/AterCorvus/RPS/Server/src/GameLogic"
)

func GetChallenges(username string, password string) ([]controllers.Challenge, error) {
	userID := GetLoggedUserID(username, password)
	if userID > 0 {
		return controllers.GetChallengesByRecipientUserID(userID)
	}
	return nil, errors.New("user not logged in")
}

func AddChallenge(challengerUsername string, opponentUsername string, bet float64, move gamelogic.Move, password string) error {
	//Check if user logged in
	id := GetLoggedUserID(challengerUsername, password)
	if id == 0 {
		return errors.New("wrong username or password")
	}
	challenger := controllers.GetUserByUsername(challengerUsername)
	if challenger == nil || challenger.Password != password {
		return errors.New("wrong username or password")
	}
	opponent := controllers.GetUserByUsername(opponentUsername)
	if opponent == nil {
		return errors.New("user not found")
	}
	if challenger.Funds < bet {
		return errors.New("not enough funds")
	}

	err := controllers.CreateChallenge(challenger.UserID, opponent.UserID, bet, move)
	if err != nil {
		return err
	}

	err = controllers.WidthdrawUserFunds(challenger.Username, bet)
	if err == nil {
		challenger.Funds -= bet
	}
	return err
}

func ResolveChallenge(challengeID int, oponentMove gamelogic.Move, username string, password string) (gamelogic.Result, error) {
	id := GetLoggedUserID(username, password)
	if id == 0 {
		return 3, errors.New("wrong username or password")
	}
	challenge, err := controllers.GetChallengeByID(challengeID)
	if err != nil {
		return 3, err
	}
	if challenge == nil {
		return 3, errors.New("challenge not found")
	}
	oponent := controllers.GetUserByUserID(challenge.RecipientUserID)
	if oponent == nil {
		return 3, errors.New("user not found")
	}
	if oponent.Funds < challenge.Amount {
		return 3, errors.New("not enough funds")
	}
	result := gamelogic.Play(challenge.Move, oponentMove)
	if result == gamelogic.Win {
		err = controllers.AddUserFundsByUserID(challenge.RecipientUserID, challenge.Amount)
		if err != nil {
			return 3, err
		}
	} else if result == gamelogic.Lose {
		err = controllers.AddUserFundsByUserID(challenge.SenderUserID, challenge.Amount*2)
		if err != nil {
			return 3, err
		}
		err = controllers.WidthdrawUserFunds(oponent.Username, challenge.Amount)
		if err != nil {
			return 3, err
		}
		oponent.Funds -= challenge.Amount
	} else {
		err = controllers.AddUserFundsByUserID(challenge.SenderUserID, challenge.Amount)
		if err != nil {
			return 3, err
		}
	}
	controllers.DeleteChallenge(challengeID)
	return result, nil
}

func DeclineChallenge(challengeID int, username string, password string) error {
	id := GetLoggedUserID(username, password)
	if id == 0 {
		return errors.New("wrong username or password")
	}
	challenge, err := controllers.GetChallengeByID(challengeID)
	if err != nil {
		return err
	}
	if challenge == nil {
		return errors.New("challenge not found")
	}
	if challenge.RecipientUserID != id {
		return errors.New("you are not the recipient of this challenge")
	}
	err = controllers.AddUserFundsByUserID(challenge.SenderUserID, challenge.Amount)
	if err != nil {
		return err
	}
	controllers.DeleteChallenge(challengeID)
	return nil
}
