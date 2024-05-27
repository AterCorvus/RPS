package authentification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	utils "github.com/AterCorvus/RPS/CLient/src/Utils"
)

var username string
var password string

type Payload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUrername() string {
	return username
}

func GetPassword() string {
	return password
}

func sendRequest(url string, localUsername string, localPassword string) int {
	data := Payload{
		Username: localUsername,
		Password: localPassword,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	return resp.StatusCode
}

func Login(localUsername string, localPassword string) bool {
	respCode := sendRequest(utils.BaseURL+"/user/login", localUsername, localPassword)
	if respCode >= 200 && respCode <= 299 {
		username = localUsername
		password = localPassword
		fmt.Println("Login was successful")
		return true
	} else {
		return false
	}
}

func Registration(localUsername string, localPassword string) bool {
	respCode := sendRequest(utils.BaseURL+"/user/register", localUsername, localPassword)
	if respCode >= 200 && respCode <= 299 {
		username = localUsername
		password = localPassword
		fmt.Println("Registration was successful")
		return true
	} else {
		return false
	}
}

func Logout() {
	if username != "" && password != "" {
		sendRequest(utils.BaseURL+"/user/logout", username, password)
		username = ""
		password = ""
	}
}
