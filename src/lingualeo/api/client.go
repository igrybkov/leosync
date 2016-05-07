package api

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strings"
)

type linguaLeoApi struct {
	connection Client
}

type Client struct {
	token   *token
	request *gorequest.SuperAgent
}

func (c Client) authorize(email string, password string) []error {

	errs := c.validateCredentials(email, password)

	if errs != nil {
		return errs
	}

	resp, body, errs := c.request.Get(loginUrl).Query("email=" + email).Query("password=" + password).End()
	if resp.StatusCode != http.StatusOK {
		errs = append(errs, errors.New("Failed login: status code is "+resp.Status))
	}
	var loginResp loginResp
	err := json.NewDecoder(strings.NewReader(body)).Decode(&loginResp)
	if err != nil {
		errs = append(errs, errors.New("Failed decode: "+body))
	}
	if strings.TrimSpace(loginResp.ErrorMsg) != "" {
		errs = append(errs, errors.New("Failed login: "+loginResp.ErrorMsg))
	}
	c.token = &loginResp.Token

	return errs
}

func (c Client) validateCredentials(email string, password string) []error {
	var errs []error = nil

	if strings.TrimSpace(email) == "" {
		errs = append(errs, errors.New("Username should not be empty"))
	}
	if strings.TrimSpace(password) == "" {
		errs = append(errs, errors.New("Password should not be empty"))
	}
	return errs
}

func (c Client) GetTranslations(word string) ([]error, Word) {
	_, body, errs := c.request.Get(translateUrl).Query("word=" + word).End()

	var translations Word
	//translations.Value = word
	decodedError := json.NewDecoder(strings.NewReader(body)).Decode(&translations)
	if decodedError != nil {
		errs = append(errs, errors.New(decodedError.Error()))
	}
	return errs, translations
}

func (c Client) AddWord(word, translation string) []error {
	_, _, errs := c.request.Get(addWordUrl).Query("word=" + word).Query("tword=" + translation).End()
	return errs
}
