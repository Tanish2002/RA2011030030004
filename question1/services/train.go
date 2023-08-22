package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const URL = "http://20.244.56.144/train"

type TrainAuthRequest struct {
	CompanyName  string `json:"companyName"`
	ClientID     string `json:"clientId"`
	OwnerName    string `json:"ownerName"`
	OwnerEmail   string `json:"ownerEmail"`
	RollNo       string `json:"rollNo"`
	ClientSecret string `json:"clientSecret"`
}
type TrainAuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
type TrainResponse struct {
	TrainName     string `json:"trainName"`
	TrainNumber   string `json:"trainNumber"`
	DepartureTime struct {
		Hours   int `json:"Hours"`
		Minutes int `json:"Minutes"`
		Seconds int `json:"Seconds"`
	} `json:"departureTime"`
	SeatsAvailable struct {
		Sleeper int `json:"sleeper"`
		AC      int `json:"AC"`
	} `json:"seatsAvailable"`
	Price struct {
		Sleeper int `json:"sleeper"`
		AC      int `json:"AC"`
	} `json:"price"`
	DelayedBy int `json:"delayedBy"`
}

func (t TrainAuthResponse) isExpired() bool {
	expiresIn := int64(1692710096)

	// Convert "expires_in" value to a time.Duration
	expiresDuration := time.Duration(expiresIn) * time.Second

	// Calculate the expiration time by adding the duration to the current time
	expirationTime := time.Now().Add(expiresDuration)

	// Check if the token is expired
	if time.Now().After(expirationTime) {
		return true
	} else {
		return false
	}
}

// Generate New Token
func NewToken() *TrainAuthResponse {
	url := URL + "/auth"
	method := "POST"

	payload := TrainAuthRequest{
		CompanyName:  "Train Central",
		ClientID:     "a0ff42a2-44eb-4899-addb-581045fca9e0",
		OwnerName:    "Tanish Khare",
		OwnerEmail:   "tk8351@srmist.edu.in",
		RollNo:       "RA2011030030004",
		ClientSecret: "NuiUgbXhycXDtAIh",
	}
	payloadBytes, err := json.Marshal(payload)
	client := &http.Client{}
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return nil
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	// Decode the response body into TrainAuthResponse struct
	var authResponse TrainAuthResponse
	err = json.NewDecoder(res.Body).Decode(&authResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}
	return &authResponse
}

// Get all Trains
func (t TrainAuthResponse) GetTrains() *[]TrainResponse {
	url := URL + "/trains"

	if t.isExpired() {
		t = *NewToken()
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Authorization", "Bearer "+t.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	var trains []TrainResponse
	err = json.NewDecoder(res.Body).Decode(&trains)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}
	return &trains
}

// Get Single Trains
func (t TrainAuthResponse) GetTrain(id int) *TrainResponse {
	url := fmt.Sprintf("%s/trains/%d", URL, id)

	if t.isExpired() {
		t = *NewToken()
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Authorization", "Bearer "+t.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	var trains TrainResponse
	err = json.NewDecoder(res.Body).Decode(&trains)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}
	return &trains
}
