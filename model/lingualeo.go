package model

import (
	"errors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/franela/goreq"
)

func NewLingualeoStorage(email string, password string) DictionaryStorageInterface {
	leoStorage := &linguaLeoStorage{}

	leoStorage.api.credentials.Email = email
	leoStorage.api.credentials.Password = password

	urls := lingiaLeoApiURLs{
		Login:           "https://api.lingualeo.com/api/login",
		UserDictionary:  "http://lingualeo.com/userdict/json",
		Translate:       "https://api.lingualeo.com/gettranslates",
		AddWord:         "https://api.lingualeo.com/addword",
		SetPicture:      "https://lingualeo.com/userdict3/setPicture",
		DownloadPicture: "http://lingualeo.com/userdict3/downloadPicture",
	}
	leoStorage.api.urls = urls

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	leoStorage.api.cookie = cookieJar

	return leoStorage
}

type lingiaLeoApiURLs struct {
	Login           string
	UserDictionary  string
	Translate       string
	AddWord         string
	SetPicture      string
	DownloadPicture string
}

type linguaLeoStorage struct {
	api struct {
		urls        lingiaLeoApiURLs
		credentials struct {
			Email    string
			Password string
		}
		isAuthorized bool
		cookie       *cookiejar.Jar
	}
}

type leoWord struct {
	ID           uint   `json:"word_id"`
	Value        string `json:"word_value"`
	Transcript   string `json:"transcription"`
	Translations []struct {
		Value      string `json:"value"`
		PictureURL string `json:"pic_url"`
	} `json:"translate"`
	TranslateID int    `json:"translate_id"`
	SoundURL    string `json:"sound_url"`
	PictureURL  string `json:"pic_url"`
}

func (leo *linguaLeoStorage) ensureAuthorized() (err error) {
	if !leo.api.isAuthorized {
		err = leo.authorize(leo.api.credentials.Email, leo.api.credentials.Password)
	}
	return err
}

func (leo *linguaLeoStorage) authorize(email string, password string) error {
	var loginReq = struct {
		Email    string `url:"email"`
		Password string `url:"password"`
	}{
		Email:    email,
		Password: password,
	}

	var loginResp = struct {
		ErrorMsg string `json:"error_msg"`
		User     struct {
			Username string `json:"nickname"`
			ID       int    `json:"user_id"`
			Key      string `json:"autologin_key"`
		} `json:"user"`
	}{}

	err := leo.get(leo.api.urls.Login, loginReq, &loginResp)
	if strings.TrimSpace(loginResp.ErrorMsg) != "" {
		err = errors.New("Failed login: " + loginResp.ErrorMsg)
	}

	if err == nil {
		leo.api.isAuthorized = true
	}

	return err
}

func (leo *linguaLeoStorage) get(url string, requestData interface{}, result interface{}) error {
	var err error

	resp, err := goreq.Request{
		Uri:         url,
		QueryString: requestData,
		CookieJar:   leo.api.cookie,
	}.Do()
	if err != nil {
		return err
	}

	// Todo: check is there logical error. Looks like may happen the case when the body will never be closed
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = resp.Body.FromJsonTo(result)

	if resp.StatusCode != http.StatusOK {
		return errors.New("Request wasn't successful: response status code is " + resp.Status)
	}

	return err
}

func (leo *linguaLeoStorage) post(url string, requestData interface{}, result interface{}) error {
	var err error

	resp, err := goreq.Request{
		Uri:       url,
		Method:    http.MethodPost,
		Body:      requestData,
		CookieJar: leo.api.cookie,
	}.Do()
	if err != nil {
		return err
	}

	// Todo: check is there logical error. Looks like may happen the case when the body will never be closed
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Fatalln(resp.Header.Get("Location"))

	err = resp.Body.FromJsonTo(result)

	if resp.StatusCode != http.StatusOK {
		return errors.New("Request wasn't successful: response status code is " + resp.Status)
	}

	return err
}

func (leo *linguaLeoStorage) Add(word *Word) error {
	err := leo.ensureAuthorized()
	if err != nil {
		return err
	}

	addWordRequest := struct {
		Word        string `url:"word"`
		Translation string `url:"tword"`
		Context     string `url:"context"`
	}{
		Word:        word.Word,
		Translation: word.Translation,
	}

	resp := struct {
		ErrorMsg   string `json:"error_msg"`
		ID         uint   `json:"word_id"`
		Value      string `json:"word_value"`
		Transcript string `json:"transcription"`
		//	Created     time.Time       `json:"created_at"`
		//	LastUpdated time.Time       `json:"last_updated_at"`
		//Translations []UserTranslate `json:"translate"`
		TranslateID int    `json:"translate_id"`
		SoundURL    string `json:"sound_url"`
		PictureURL  string `json:"pic_url"`
	}{}

	err = leo.get(leo.api.urls.AddWord, addWordRequest, &resp)
	if strings.TrimSpace(resp.ErrorMsg) != "" {
		err = errors.New("Can't add the word \"" + word.Word + "\": " + resp.ErrorMsg)
	}

	return err
}

func (leo *linguaLeoStorage) Get(word string) (*Word, error) {
	panic("implement me")
}

func (leo *linguaLeoStorage) Has(word string) bool {
	panic("implement me")
}

func (leo *linguaLeoStorage) Remove(word string) error {
	panic("implement me")
}

func (leo *linguaLeoStorage) GetList() (WordList, error) {
	wl := WordList{}
	err := leo.ensureAuthorized()
	if err != nil {
		return wl, err
	}

	page := 1
	for {
		getPageRequest := struct {
			GroupId  string `url:"groupId"`
			SortBy   string `url:"sortBy"`
			Filter   string `url:"filter"`
			WordType int    `url:"wordType"`
			Page     int    `url:"page"`
		}{
			GroupId:  "dictionary",
			SortBy:   "date",
			Filter:   "all",
			WordType: 0,
			Page:     page,
		}

		resp := struct {
			ErrorMsg string `json:"error_msg"`
			UserDict []leoWord
		}{}

		err = leo.post(leo.api.urls.UserDictionary, getPageRequest, &resp)
		if strings.TrimSpace(resp.ErrorMsg) != "" {
			return wl, errors.New("Can't get the page from the API: " + resp.ErrorMsg)
		}
		page++
		break
	}

	return wl, err
}

func (leo *linguaLeoStorage) SaveList(list *WordList) error {
	panic("implement me")
}
