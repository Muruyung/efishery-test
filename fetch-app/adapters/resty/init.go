package resty

import "github.com/go-resty/resty/v2"

// Resty is wrapper for resty client
type Resty struct {
	*resty.Client
}

// Init create new resty client
func Init() *Resty {
	return &Resty{
		resty.New(),
	}
}
