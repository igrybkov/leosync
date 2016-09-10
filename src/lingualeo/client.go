package lingualeo

import (
	"errors"
	"github.com/franela/goreq"
	"github.com/igrybkov/leosync/src/lingualeo/request"
	"github.com/igrybkov/leosync/src/lingualeo/response"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

// Client is an API client
type apiClient struct {
	cookie http.CookieJar
}

func (c apiClient) get(url string, requestData interface{}, result interface{}) []error {
	var errs []error

	resp, err := goreq.Request{
		Uri:         url,
		QueryString: requestData,
		CookieJar:   c.cookie,
	}.Do()
	if err != nil {
		errs = append(errs, errors.New(err.Error()))
		log.Fatalln(err.Error())
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

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

func (c apiClient) post(url string, requestData interface{}) []error {

	var errs []error
	req := goreq.Request{
		Method:      "POST",
		Body:        requestData,
		Uri:         url,
		ContentType: "application/x-www-form-urlencoded; charset=UTF-8",
		CookieJar:   c.cookie,
	}
	req.AddHeader("X-Requested-With", "XMLHttpRequest")
	resp, err := req.Do()

	if err != nil {
		errs = append(errs, errors.New(err.Error()))
		log.Fatalln(err.Error())
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		errs = append(errs, errors.New("Failed login: status code is "+resp.Status))
	}

	return errs
}

// DownloadPicture posts a picture to the translation
func (c apiClient) DownloadPicture(url string, translateID string) []error {
	req := "url=" + url + "&translate_id=" + translateID
	errs := c.post(downloadPictureURL, req)
	return errs
}

func (c apiClient) authorize(email string, password string) []error {
	req := request.LoginRequest{
		Email:    email,
		Password: password,
	}

	var loginResp response.LoginResponse
	errs := c.get(loginURL, req, loginResp)
	if strings.TrimSpace(loginResp.ErrorMsg) != "" {
		errs = append(errs, errors.New("Failed login: "+loginResp.ErrorMsg))
	}

	return errs
}

func (c apiClient) validateCredentials(email string, password string) []error {
	var errs []error

	if strings.TrimSpace(email) == "" {
		errs = append(errs, errors.New("Username should not be empty"))
	}
	if strings.TrimSpace(password) == "" {
		errs = append(errs, errors.New("Password should not be empty"))
	}
	return errs
}

// GetTranslations returns translations for a word
func (c apiClient) GetTranslations(word string) ([]error, response.Word) {
	req := request.TranslationRequest{
		Word: word,
	}

	translations := response.Word{}
	errs := c.get(translateURL, req, &translations)
	if strings.TrimSpace(translations.ErrorMsg) != "" {
		errs = append(errs, errors.New("Something went wrong: "+translations.ErrorMsg))
	}

	return errs, translations
}

// AddWord posts new word to the API
func (c apiClient) AddWord(word, translation string) ([]error, response.Word) {
	req := request.AddWordRequest{
		Word:        word,
		Translation: translation,
	}

	var result response.Word
	errs := c.get(addWordURL, req, &result)
	if strings.TrimSpace(result.ErrorMsg) != "" {
		errs = append(errs, errors.New("Something went wrong: "+result.ErrorMsg))
	}

	return errs, result
}

// AddWordWithContext add a word with a context
// ToDo: Rewrite the method to combine it with an AddWord call
func (c apiClient) AddWordWithContext(word, translation string, context string) ([]error, response.Word) {
	req := request.AddWordWithContextRequest{
		Word:        word,
		Translation: translation,
		Context:     context,
	}

	var result response.Word
	errs := c.get(addWordURL, req, &result)
	if strings.TrimSpace(result.ErrorMsg) != "" {
		errs = append(errs, errors.New("Something went wrong: "+result.ErrorMsg))
	}

	return errs, result
}

// NewClient returns new instance of the API client
func newClient(email string, password string) ([]error, apiClient) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		log.Panic(err)
	}

	client := apiClient{
		cookie: cookieJar,
	}
	errs := client.authorize(email, password)

	return errs, client
}
