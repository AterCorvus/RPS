package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	authentification "github.com/AterCorvus/RPS/CLient/src/Authentification"
	utils "github.com/AterCorvus/RPS/CLient/src/Utils"
)

type Payload struct {
	ID         int     `json:"id"`
	Challenger string  `json:"sender_userId"`
	Opponent   string  `json:"recipient_userId"`
	Bet        float64 `json:"amount"`
	Move       int     `json:"move"`
	PassWord   string  `json:"passwords"`
}

func ChallengePlayer(oponentName string, bet float64, move int) {
	data := Payload{
		Challenger: authentification.GetUrername(),
		Opponent:   oponentName,
		Move:       move,
		Bet:        bet,
		PassWord:   authentification.GetPassword(),
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", utils.BaseURL+"/challenge/create", body)
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
}

func AcceptChallenge(id int, move int) {
	data := Payload{
		ID:       id,
		Opponent: authentification.GetUrername(),
		Move:     move,
		PassWord: authentification.GetPassword(),
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", utils.BaseURL+"/challenge/accept", body)
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

	var result struct {
		Result int `json:"result"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}
	if result.Result == 1 {
		fmt.Println("You won")
	} else if result.Result == 0 {
		fmt.Println("You lost")
	} else {
		fmt.Println("Draw")
	}
}

func DeclineChallenge(id int) {
	data := Payload{
		ID:       id,
		Opponent: authentification.GetUrername(),
		PassWord: authentification.GetPassword(),
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", utils.BaseURL+"/challenge/decline", body)
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
}

func ViewPendingChallenges() {
	baseURL, err := url.Parse(utils.BaseURL + "/challenge/challenges")
	if err != nil {
		log.Fatalln(err)
	}

	params := url.Values{}
	params.Add("username", authentification.GetUrername())
	params.Add("password", authentification.GetPassword())
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Response Body:", string(body))
}

func ShowPlayersOnline() {
	baseURL, err := url.Parse(utils.BaseURL + "/user/usersonline")
	if err != nil {
		log.Fatalln(err)
	}

	params := url.Values{}
	params.Add("username", authentification.GetUrername())
	params.Add("password", authentification.GetPassword())
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Response Body:", string(body))
}
