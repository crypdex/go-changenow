package changenow

import (
	resty "gopkg.in/resty.v1"
)

var baseURL = "https://changenow.io/api/v1"

// ChangeNow is the root client object
type ChangeNow struct {
	apikey string
	client *resty.Client
}

// New creates a configured instance of ChangeNow
func New(apikey string) *ChangeNow {
	return &ChangeNow{
		apikey: apikey,
		client: resty.New().SetHostURL(baseURL),
	}
}
