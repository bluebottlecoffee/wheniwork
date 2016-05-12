package wheniwork

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Credentials struct {
	Username string
	Password string
	Key      string
	BaseURL  string
}

// POST https://api.wheniwork.com/2/login
func Login(creds *Credentials) (*Client, error) {
	if creds.BaseURL == "" {
		creds.BaseURL = "https://api.wheniwork.com/2"
	}

	httpClient := &http.Client{}

	loginRequest := LoginRequest{
		Username: creds.Username,
		Password: creds.Password,
	}

	loginRequestBody, err := json.Marshal(loginRequest)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", creds.BaseURL+"/login", bytes.NewReader(loginRequestBody))

	if err != nil {
		return nil, err
	}

	req.Header.Add("W-Key", creds.Key)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	loginResponse := &LoginResponse{}
	err = json.Unmarshal(contents, &loginResponse)

	return &Client{Token: loginResponse.Login.Token, HttpClient: httpClient}, nil
}
