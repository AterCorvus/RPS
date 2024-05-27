package restapi

import (
	"github.com/gin-gonic/gin"

	adapter "github.com/AterCorvus/RPS/Server/src/Adapters"
)

var router *gin.Engine

func StartServer() {
	router = gin.Default()

	// user api
	router.GET("/user/usersonline", adapter.GetUsersOnline)
	router.GET("/user/funds", adapter.GetUserFunds)
	router.POST("/user/register", adapter.RegisterUser)
	router.POST("/user/login", adapter.LoginUser)
	router.POST("/user/logout", adapter.LogoutUser)
	router.POST("/user/addmoney", adapter.AddUserFunds)
	router.POST("/user/widthdrawmoney", adapter.WidthdrawUserFunds)

	// challenge api
	router.GET("/challenge/challenges", adapter.GetChallenges)
	router.POST("/challenge/create", adapter.AddChallenge)
	router.POST("/challenge/accept", adapter.AcceptChallenge)
	router.POST("/challenge/decline", adapter.DeclineChallenge)

	router.Run("localhost:9090")
}
