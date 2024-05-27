package adapters

import (
	"net/http"

	"github.com/gin-gonic/gin"

	gamelogic "github.com/AterCorvus/RPS/Server/src/GameLogic"
	services "github.com/AterCorvus/RPS/Server/src/Services"
)

type ChallengeAdapter struct {
	ID         int            `json:"id"`
	Challenger string         `json:"sender_userId"`
	Opponent   string         `json:"recipient_userId"`
	Bet        float64        `json:"amount"`
	Move       gamelogic.Move `json:"move"`
	Password   string         `json:"passwords"`
}

func GetChallenges(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	challenges, err := services.GetChallenges(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"challenges": challenges})
}

func AddChallenge(c *gin.Context) {
	var challengeAdapter ChallengeAdapter
	if err := c.ShouldBindJSON(&challengeAdapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if challengeAdapter.Challenger == "" || challengeAdapter.Opponent == "" || challengeAdapter.Bet <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please fill all fields"})
		return
	}

	err := services.AddChallenge(challengeAdapter.Challenger, challengeAdapter.Opponent, challengeAdapter.Bet, challengeAdapter.Move, challengeAdapter.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func AcceptChallenge(c *gin.Context) {
	var challengeAdapter ChallengeAdapter
	if err := c.ShouldBindJSON(&challengeAdapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if challengeAdapter.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please fill all fields correctly"})
		return
	}

	result, err := services.ResolveChallenge(challengeAdapter.ID, challengeAdapter.Move, challengeAdapter.Opponent, challengeAdapter.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func DeclineChallenge(c *gin.Context) {
	var challengeAdapter ChallengeAdapter
	if err := c.ShouldBindJSON(&challengeAdapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if challengeAdapter.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please fill all fields correctly"})
		return
	}

	err := services.DeclineChallenge(challengeAdapter.ID, challengeAdapter.Opponent, challengeAdapter.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
