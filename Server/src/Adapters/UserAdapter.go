package adapters

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "github.com/AterCorvus/RPS/Server/src/Services"
)

type UserAdapter struct {
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	FundsChange float64 `json:"funds_change"`
}

func RegisterUser(c *gin.Context) {
	var userAdapter UserAdapter
	if err := c.ShouldBindJSON(&userAdapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userAdapter.Username == "" || userAdapter.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
		return
	}

	err := services.RegisterUser(userAdapter.Username, userAdapter.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func LoginUser(c *gin.Context) {
	var userAdapter UserAdapter
	if err := c.ShouldBindJSON(&userAdapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userAdapter.Username == "" || userAdapter.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
		return
	}

	result := services.LoginUser(userAdapter.Username, userAdapter.Password)
	if result {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong username or password"})
	}
}

func LogoutUser(c *gin.Context) {
	var userAdapter UserAdapter
	if err := c.ShouldBindJSON(&userAdapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userAdapter.Username == "" || userAdapter.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
		return
	}

	result := services.LogoutUser(userAdapter.Username, userAdapter.Password)
	if result {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong username or password"})
	}
}

func AddUserFunds(c *gin.Context) {
	var userAdapter UserAdapter
	if err := c.ShouldBindJSON(&userAdapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userAdapter.Username == "" || userAdapter.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
		return
	}

	err := services.AddUserFunds(userAdapter.Username, userAdapter.Password, userAdapter.FundsChange)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func WidthdrawUserFunds(c *gin.Context) {
	var userAdapter UserAdapter
	if err := c.ShouldBindJSON(&userAdapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userAdapter.Username == "" || userAdapter.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
		return
	}

	err := services.WidthdrawUserFunds(userAdapter.Username, userAdapter.Password, userAdapter.FundsChange)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetUsersOnline(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	users, err := services.GetUsersOnline(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserFunds(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	funds, err := services.GetUserFunds(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"funds": funds})
}