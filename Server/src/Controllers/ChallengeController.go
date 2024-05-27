package controllers

import (
	"log"

	gamelogic "github.com/AterCorvus/RPS/Server/src/GameLogic"
)

type Challenge struct {
	ID              int
	SenderUserID    int
	RecipientUserID int
	Amount          float64
	Move            gamelogic.Move
}

func CreateChallenge(senderUserId int, recipientUserId int, bet float64, move gamelogic.Move) error {
	query := `INSERT INTO pending_challenges (sender_user_id, recipient_user_id, bet, move) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, senderUserId, recipientUserId, bet, move)
	if err != nil {
		return err
	}
	return nil
}

func GetChallengesByRecipientUserID(recipientUserID int) ([]Challenge, error) {
	query := `SELECT * FROM pending_challenges WHERE recipient_user_id = ?`
	rows, err := db.Query(query, recipientUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var challenges []Challenge
	for rows.Next() {
		var c Challenge
		err = rows.Scan(&c.ID, &c.SenderUserID, &c.RecipientUserID, &c.Amount, &c.Move)
		if err != nil {
			log.Fatal(err)
		}
		c.Move = 4 // Hide the move from the recipient
		challenges = append(challenges, c)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return challenges, nil
}

func GetChallengeByID(id int) (*Challenge, error) {
	query := `SELECT * FROM pending_challenges WHERE id = ?`
	var c Challenge
	err := db.QueryRow(query, id).Scan(&c.ID, &c.SenderUserID, &c.RecipientUserID, &c.Amount, &c.Move)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func DeleteChallenge(id int) error {
	query := `DELETE FROM pending_challenges WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
