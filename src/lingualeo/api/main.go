package api

import (
	"net/http/cookiejar"
)

func NewClient(email string, password string) ([]error, Client) {
	cookieJar, _ := cookiejar.New(nil)
	client := Client{
		cookie: cookieJar,
	}
	errs := client.authorize(email, password)

	return errs, client
}
