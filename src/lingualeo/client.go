package lingualeo

import (
	"errors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/franela/goreq"
)

type Client struct {
	config ConnectionConfig
	http   *httpClient
}
type httpClient struct {
	cookie       *cookiejar.Jar
	isAuthorized bool
}

func (c *httpClient) IsAuthorized() bool {
	return c.isAuthorized
}

// TODO: Add proper error handling
func (c *httpClient) Authorize(email string, password string) error {
	req := LoginRequest{
		Email:    email,
		Password: password,
	}

	var loginResp LoginResponse
	errs := c.get(loginURL, req, loginResp)
	if strings.TrimSpace(loginResp.ErrorMsg) != "" {
		errs = append(errs, errors.New("Failed login: "+loginResp.ErrorMsg))
	}
	return nil
}

// TODO: Add proper error handling
func (c httpClient) get(url string, requestData interface{}, result interface{}) []error {
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

// New instance of the LinguaLeo client will be returned
func New(connectionConfig ConnectionConfig) (client Client, err error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return client, err
	}

	client = Client{
		config: connectionConfig,
		http: &httpClient{
			cookie: cookieJar,
		},
	}

	return client, err
}

func (c *Client) ensureAuthorized() (err error) {
	if !c.http.isAuthorized {
		err = c.http.Authorize(c.config.Email, c.config.Password)
	}
	return err
}

func (c *Client) AddWord(req AddWordRequest) (err error, resp *Word) {
	err = c.ensureAuthorized()
	if err != nil {
		return err, resp
	}

	errs := c.http.get(addWordURL, req, &resp)
	if strings.TrimSpace(resp.ErrorMsg) != "" {
		errs = append(errs, errors.New("Something went wrong: "+resp.ErrorMsg))
	}

	return err, resp
}
