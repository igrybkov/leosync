package model

import (
	"strings"

	"github.com/pkg/errors"
)

type DictionaryInterface interface {
	Add(word *Word) error
	Get(word string) (*Word, error)
	Has(word string) bool
	Remove(word string) error
}

type DictionaryStorageInterface interface {
	DictionaryInterface
	GetList() (WordList, error)
	SaveList(list *WordList) error
}

type Word struct {
	Word          string
	Transcription string
	Context       string
	Stem          string
	Translation   string
	ImageURL      string
	SoundURL      string
}

type WordList map[string]*Word

func (wl WordList) Add(word *Word) error {
	var err error

	key := strings.ToLower(word.Word)
	_, ok := wl[key]
	if ok {
		return errors.New("Word " + word.Word + " already exist")
	}

	wl[key] = word
	return err
}

func (wl *WordList) Remove(word string) error {
	delete(*wl, strings.ToLower(word))
	return nil
}

func (wl WordList) Get(word string) (*Word, error) {
	var err error
	wordData, ok := wl[strings.ToLower(word)]
	if !ok {
		err = errors.New("Word " + word + " doesn't exist")
	}

	return wordData, err
}

func (wl WordList) Has(word string) bool {
	_, ok := wl[strings.ToLower(word)]
	return ok
}
