package kindle

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"log"
)

type Word struct {
	Id   string `gorm:"primary_key"`
	Word string
	Stem string
	Lang string
}

func (Word) TableName() string {
	return "WORDS"
}

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
