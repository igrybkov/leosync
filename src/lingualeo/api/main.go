package api

import (
	"log"
	"net/http/cookiejar"
)

// NewClient returns new instance of the API client
func NewClient(email string, password string) ([]error, Client) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		log.Panic(err)
	}

	client := Client{
		cookie: cookieJar,
	}
	errs := client.authorize(email, password)

	return errs, client
}
