package helper

import (
	"github.com/go-resty/resty/v2"
)

func RestyRequest(c *resty.Client, uri string) (req *resty.Request) {
	req = c.SetBaseURL(uri).R().
		SetHeader("Content-Type", "application/json")
	return
}
