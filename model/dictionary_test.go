package model_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/igrybkov/leosync/model"
)

func TestWordList_Add(t *testing.T) {
	var err error

	wl := model.WordList{}
	wordTesting := model.Word{
		Word: "Testing",
	}

	wordMocking := model.Word{
		Word: "Mocking",
	}

	err = wl.Add(&wordTesting)
	if err != nil {
		t.Log("The error happened when adding a word for the first time: " + err.Error())
		t.Fail()
	}

	err = wl.Add(&wordMocking)
	if err != nil {
		t.Log("The error happened when adding a word for the first time: " + err.Error())
		t.Fail()
	}

	err = wl.Add(&wordTesting)
	if err == nil {
		t.Log("The error must be raised when a duplicated word added")
		t.Fail()
	}

	returnedWord, ok := wl[strings.ToLower(wordTesting.Word)]
	if !ok {
		t.Log("The word should be present in the map and available by the lower case")
		t.Fail()
	}

	if &wordTesting != returnedWord {
		t.Log("The word should be the same as a tested one")
		t.Fail()
	}

	totalWords := len(wl)
	if 2 != totalWords {
		t.Log("The map should containt both words, but the length of the map is " + strconv.Itoa(totalWords))
		t.Fail()
	}
}

func TestWordList_Has(t *testing.T) {
	wl := model.WordList{}
	wordTesting := model.Word{
		Word: "That's something STR@NG3 😊!!!11",
	}

	if wl.Has("testing") {
		t.Log("The word reported as existing, but it is not present in the list")
		t.Fail()
	}

	err := wl.Add(&wordTesting)
	if err != nil {
		t.Log("The error happened when adding a word.")
		t.Fail()
	}

	if !wl.Has(wordTesting.Word) {
		t.Log("The word as missing, but it has been added to the list")
		t.Fail()
	}

	if wl.Has("unknown_word") {
		t.Log("The word reported as existing, but it is not added to the list")
		t.Fail()
	}
}

func TestWordList_Remove(t *testing.T) {
	var err error

	wl := model.WordList{}
	wordTesting := model.Word{
		Word: "Testing",
	}

	if wl.Has("testing") {
		t.Log("The word reported as existing, but it is not present in the list")
		t.Fail()
	}

	err = wl.Remove("testing")
	if err != nil {
		t.Log("The list shouldn't raise an error when removes not existing word")
		t.Fail()
	}

	err = wl.Add(&wordTesting)
	if err != nil {
		t.Log("The error happened when adding a word.")
		t.Fail()
	}

	if !wl.Has("testing") {
		t.Log("The word not present in the list and therefore cannot be removed")
		t.Fail()
	}

	err = wl.Remove("testing")
	if err != nil {
		t.Log("The list shouldn't raise an error when removes not existing word")
		t.Fail()
	}

	if wl.Has("testing") {
		t.Log("The word should be absent in the list, but it still present")
		t.Fail()
	}

	_, ok := wl[strings.ToLower(wordTesting.Word)]
	if ok {
		t.Log("The word should not exist in the map")
		t.Fail()
	}
}

func TestWordList_Get(t *testing.T) {
	var err error

	wl := model.WordList{}
	wordTesting := model.Word{
		Word: "Testing",
	}

	_, err = wl.Get("unknown word")
	if err == nil {
		t.Log("The error be returned when a word is missing in the list.")
		t.Fail()
	}

	err = wl.Add(&wordTesting)
	if err != nil {
		t.Log("The error happened when adding a word.")
		t.Fail()
	}

	returnedWord, err := wl.Get(strings.ToUpper(wordTesting.Word))
	if err != nil {
		t.Log("The word should be present in the map and available by any case")
		t.Fail()
	}

	if &wordTesting != returnedWord {
		t.Log("The word should be the same as a tested one")
		t.Fail()
	}
}
