package wheniwork

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const iSO8601 = "2006-01-02 15:04:05"

type doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	Token      string
	HttpClient doer
	baseURL    string
}

func (c *Client) url(path string) *url.URL {
	u, err := url.Parse(c.baseURL + path)

	if err != nil {
		panic(err)
	}

	return u
}

// Listing Shifts
// http://dev.wheniwork.com/#listing-shifts
func (c *Client) ListShifts(params *ListShiftParams) (*ListShiftsResponse, error) {
	u := c.url("/shifts")
	q := u.Query()

	if !params.Start.IsZero() {
		q.Set("start", params.Start.Format(iSO8601))
	}

	if !params.End.IsZero() {
		q.Set("end", params.End.Format(iSO8601))
	}

	if len(params.LocationId) > 0 {
		q.Set("location_id", strings.Join(params.LocationId, ","))
	}

	u.RawQuery = q.Encode()

	listShiftsResponse := ListShiftsResponse{}
	err := c.request(u, &listShiftsResponse)
	return &listShiftsResponse, err
}

// Get Existing Shift
// http://dev.wheniwork.com/#get-existing-shift
func (c *Client) GetShift(id int64) (*GetShiftResponse, error) {
	u := c.url("/shift/" + string(id))
	getShiftResponse := GetShiftResponse{}
	err := c.request(u, &getShiftResponse)
	return &getShiftResponse, err
}

func (c *Client) request(url *url.URL, responseHolder interface{}) error {
	req, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return err
	}

	req.Header.Add("W-Token", c.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.HttpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(contents, responseHolder)
	return err
}
