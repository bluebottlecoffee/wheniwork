package wheniwork

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
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
