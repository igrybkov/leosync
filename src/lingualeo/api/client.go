package api

import (
	"errors"
	"github.com/franela/goreq"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	cookie http.CookieJar
}

func (c Client) get(url string, requestData interface{}, result interface{}) []error {
	var errs []error = nil

	resp, err := goreq.Request{
		Uri:         url,
		QueryString: requestData,
		CookieJar:   c.cookie,
	}.Do()

	if err != nil {
		errs = append(errs, errors.New(err.Error()))
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errs = append(errs, errors.New("Failed login: status code is "+resp.Status))
	}

	err = resp.Body.FromJsonTo(&result)
	if err != nil {
		errs = append(errs, errors.New(err.Error()))
		log.Println(err.Error())
	}

	return errs
}

func (c Client) authorize(email string, password string) []error {
	req := LoginRequest{
		Email:    email,
		Password: password,
	}

	var loginResp LoginResponse
	errs := c.get(loginUrl, req, &loginResp)
	if strings.TrimSpace(loginResp.ErrorMsg) != "" {
		errs = append(errs, errors.New("Failed login: "+loginResp.ErrorMsg))
	}

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
	req := TranslationRequest{
		Word: word,
	}

	translations := Word{}
	errs := c.get(translateUrl, req, &translations)
	if strings.TrimSpace(translations.ErrorMsg) != "" {
		errs = append(errs, errors.New("Something went wrong: "+translations.ErrorMsg))
	}

	return errs, translations
}

func (c Client) AddWord(word, translation string) ([]error, Word) {
	req := AddWordRequest{
		Word:        word,
		Translation: translation,
	}

	var result Word
	errs := c.get(addWordUrl, req, &result)
	if strings.TrimSpace(result.ErrorMsg) != "" {
		errs = append(errs, errors.New("Something went wrong: "+result.ErrorMsg))
	}

	return errs, result
}
