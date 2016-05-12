package wheniwork

import (
	"net/http"
)

type Client struct {
	Token      string
	HttpClient *http.Client
}
