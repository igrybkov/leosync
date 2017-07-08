package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

func NewKindleStorage(dbFile string) DictionaryStorageInterface {
	storage := kindleStorage{
		dbFile: dbFile,
	}
	return &storage
}

type kindleStorage struct {
	dbFile    string
	tableName string
}

func (*kindleStorage) Add(word *Word) error {
	panic("implement me")
}

func (*kindleStorage) Get(word string) (*Word, error) {
	panic("implement me")
}

func (*kindleStorage) Has(word string) bool {
	panic("implement me")
}

func (*kindleStorage) Remove(word string) error {
	panic("implement me")
}

func (*kindleStorage) GetList() (WordList, error) {
	panic("implement me")
}

func (*kindleStorage) SaveList(list *WordList) error {
	panic("implement me")
}

// Word is a representation of a word in a Kindle database
type kindleWord struct {
	ID   string `gorm:"primary_key"`
	Word string
	Stem string
	Lang string
}

// TableName returns a name of the table in a Kindle database
func (kindleWord) TableName() string {
	return "WORDS"
}

// GetWords returns a list of the words from a Kindle database
func (kindle *kindleStorage) GetWords(dbFile string) []kindleWord {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalln("Failed to connect database: " + err.Error())
	}

	var wordsList []kindleWord

	if err := db.Find(&wordsList, &kindleWord{Lang: "en"}).GetErrors(); err != nil {
		log.Fatalf("%v \n", err)
	}

	return wordsList
}
