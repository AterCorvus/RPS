package funds

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
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	FundsChange float64 `json:"funds_change"`
}

func sendRequest(url string, amount float64) {
	data := Payload{
		Username:    authentification.GetUrername(),
		Password:    authentification.GetPassword(),
		FundsChange: amount,
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
}

func AddMoney(amount float64) {
	sendRequest(utils.BaseURL+"/user/addmoney", amount)
}

func WidthdrawMoney(amount float64) {
	sendRequest(utils.BaseURL+"/user/widthdrawmoney", amount)
}

func GetMyFunds() {
	baseURL, err := url.Parse(utils.BaseURL + "/user/funds")
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

	fmt.Println("My Funds:", string(body))
}
