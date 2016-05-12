package wheniwork

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestListShifts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("testdata/shifts.json")

		if err != nil {
			t.Error(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}))
	defer ts.Close()

	client := Client{Token: "faketoken", HttpClient: &http.Client{}, BaseURL: ts.URL}
	resp, err := client.ListShifts(&ListShiftParams{})

	if err != nil {
		t.Error(err)
	}

	if len(resp.Shifts) == 0 {
		t.Error("No shifts returned")
	}
}

type requestRecorder struct {
	RequestedPath string
}

func (c *requestRecorder) Do(req *http.Request) (*http.Response, error) {
	c.RequestedPath = req.URL.String()

	responseBody := `{}`
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(responseBody))),
	}, nil
}

func TestListShiftsWithQueryParams(t *testing.T) {
	recorder := requestRecorder{}
	client := Client{Token: "faketoken", HttpClient: &recorder, BaseURL: "wheniwork.com"}
	_, err := client.ListShifts(&ListShiftParams{
		Start: time.Date(2014, 3, 5, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2014, 3, 8, 23, 59, 59, 0, time.UTC),
	})

	if err != nil {
		t.Error(err)
	}

	if recorder.RequestedPath != "wheniwork.com/shifts?end=2014-03-08+23%3A59%3A59&start=2014-03-05+00%3A00%3A00" {
		t.Error("Request was made to:", recorder.RequestedPath)
	}
}
