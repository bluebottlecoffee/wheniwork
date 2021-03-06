package wheniwork

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("testdata/login_response.json")

		if err != nil {
			t.Error(err)
		}

		if r.Header.Get("W-Key") != "iworksoharditsnotfunny" {
			t.Error("W-Key header was improperly set to", r.Header.Get("W-Key"))
		}

		contents, err := ioutil.ReadAll(r.Body)
		j := string(contents)

		if j != `{"username":"gregg@bluebottlecoffee.com","password":"notthemagicword"}` {
			t.Error("Request body was:", j)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}))
	defer ts.Close()

	client, _ := Login(&Credentials{
		Username: "gregg@bluebottlecoffee.com",
		Password: "notthemagicword",
		Key:      "iworksoharditsnotfunny",
		baseURL:  ts.URL,
	})

	if client.Token != "ilovemyboss" {
		t.Error("Unexpected \"ilovemyboss\", but got", client.Token)
	}
}
