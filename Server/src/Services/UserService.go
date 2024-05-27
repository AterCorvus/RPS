package services

import (
	"errors"

	controllers "github.com/AterCorvus/RPS/Server/src/Controllers"
)

var userService *UserService = NewUserService()

type UserService struct {
	//TODO Upadte from DB after changes
	usersOnline map[string]*controllers.User
}

func NewUserService() *UserService {
	return &UserService{
		usersOnline: make(map[string]*controllers.User),
	}
}

func RegisterUser(username string, password string) error {
	var user *controllers.User = controllers.GetUserByUsername(username)
	var err error
	if user != nil {
		return errors.New("user with this username already exists")
	} else {
		user, err = controllers.CreateUser(username, password)
		userService.usersOnline[username] = user
	}
	//TODO think about error handling
	return err
}

func LoginUser(username string, password string) bool {
	if local_user, ok := userService.usersOnline[username]; ok {
		if local_user.Password == password {
			return true
		}
	} else {
		var local_user *controllers.User = controllers.GetUserByUsername(username)
		if local_user != nil && local_user.Password == password {
			userService.usersOnline[username] = local_user
			return true
		}
	}

	return false
}

func LogoutUser(username string, password string) bool {
	if local_user, ok := userService.usersOnline[username]; ok {
		if local_user.Password == password {
			delete(userService.usersOnline, username)
			return true
		}
	}
	return false
}

func AddUserFunds(username string, password string, additionalFunds float64) error {
	userID := GetLoggedUserID(username, password)
	if userID == 0 {
		return errors.New("wrong username or password")
	}
	err := controllers.AddUserFunds(username, additionalFunds)
	//TODO think about this
	if err == nil {
		userService.usersOnline[username].Funds += additionalFunds
	}
	return err
}

func WidthdrawUserFunds(username string, password string, WidthdrawAmmount float64) error {
	userID := GetLoggedUserID(username, password)
	if userID == 0 {
		return errors.New("wrong username or password")
	}
	if userService.usersOnline[username].Funds < WidthdrawAmmount {
		return errors.New("not enough funds")
	}
	err := controllers.WidthdrawUserFunds(username, WidthdrawAmmount)
	if err == nil {
		userService.usersOnline[username].Funds -= WidthdrawAmmount
	}
	return err
}

func GetLoggedUserID(username string, password string) int {
	if local_user, ok := userService.usersOnline[username]; ok {
		if local_user.Password == password {
			return local_user.UserID
		}
	}
	return 0
}

func GetUsersOnline(username string, password string) ([]controllers.User, error) {
	if local_user, ok := userService.usersOnline[username]; ok {
		if local_user.Password == password {
			var users []controllers.User
			for _, user := range userService.usersOnline {
				newUser := controllers.User{}
				newUser.UserID = user.UserID
				newUser.Username = user.Username
				users = append(users, newUser)
			}
			return users, nil
		}
	}
	return nil, errors.New("user not logged in")
}

func GetUserFunds(username string, password string) (float64, error) {
	if local_user, ok := userService.usersOnline[username]; ok {
		if local_user.Password == password {
			return local_user.Funds, nil
		}
	}
	return 0, errors.New("user not logged in")
}
