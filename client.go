package wheniwork

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Token      string
	HttpClient *http.Client
	BaseURL    string
}

func (c *Client) ListShifts(*ListShiftParams) (*ListShiftsResponse, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/shifts", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("W-Token", c.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	listShiftsResponse := ListShiftsResponse{}
	err = json.Unmarshal(contents, &listShiftsResponse)

	if err != nil {
		return nil, err
	}

	return &listShiftsResponse, nil
}
