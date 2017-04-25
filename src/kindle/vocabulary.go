package kindle

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Word is a representation of a word in a Kindle database
type Word struct {
	ID   string `gorm:"primary_key"`
	Word string
	Stem string
	Lang string
}

// TableName returns a name of the table in a Kindle database
func (Word) TableName() string {
	return "WORDS"
}

// GetWords returns a list of the words from a Kindle database
func GetWords(dbFile string) []Word {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalln("Failed to connect database: " + err.Error())
	}

	var wordsList []Word

	if err := db.Find(&wordsList, &Word{Lang: "en"}).GetErrors(); err != nil {
		log.Fatalf("%v \n", err)
	}

	return wordsList
}
