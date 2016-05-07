package api

import (
	"github.com/parnurzeal/gorequest"
)

func NewClient(email string, password string) ([]error, Client) {
	client := Client{
		request: gorequest.New(),
	}
	errs := client.authorize(email, password)

	return errs, client

}
