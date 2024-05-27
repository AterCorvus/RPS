package controllers

type User struct {
	UserID   int
	Username string
	Password string
	Funds    float64
}

func GetUserByUsername(username string) *User {
	var user User
	query := `SELECT id, username, password, funds FROM users WHERE username = ?`
	err := db.QueryRow(query, username).Scan(&user.UserID, &user.Username, &user.Password, &user.Funds)
	if err != nil {
		return nil
	}
	return &user
}

func GetUserByUserID(userID int) *User {
	var user User
	query := `SELECT id, username, password, funds FROM users WHERE id = ?`
	err := db.QueryRow(query, userID).Scan(&user.UserID, &user.Username, &user.Password, &user.Funds)
	if err != nil {
		return nil
	}
	return &user
}

func CreateUser(username string, password string) (*User, error) {
	query := `INSERT INTO users (username, password, funds) VALUES (?, ?, 0.0)`
	result, err := db.Exec(query, username, password)
	if err != nil {
		return nil, err
	}
	//TODO check if this is concurency safe
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &User{UserID: int(id), Username: username, Password: password, Funds: 0.0}, nil
}

func AddUserFunds(username string, additionalFunds float64) error {
	query := `UPDATE users SET funds = funds + ? WHERE username = ?`
	_, err := db.Exec(query, additionalFunds, username)
	if err != nil {
		return err
	}
	return nil
}

func AddUserFundsByUserID(userID int, additionalFunds float64) error {
	query := `UPDATE users SET funds = funds + ? WHERE id = ?`
	_, err := db.Exec(query, additionalFunds, userID)
	if err != nil {
		return err
	}
	return nil
}

func WidthdrawUserFunds(username string, WidthdrawAmmount float64) error {
	query := `UPDATE users SET funds = funds - ? WHERE username = ?`
	_, err := db.Exec(query, WidthdrawAmmount, username)
	if err != nil {
		return err
	}
	return nil
}
