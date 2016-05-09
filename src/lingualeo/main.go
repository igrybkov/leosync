package lingualeo

import (
	"github.com/igrybkov/leosync/src/configuration"
	"github.com/igrybkov/leosync/src/lingualeo/api"
	"log"
)

var leoClient api.Client

func getClient() api.Client {
	if leoClient == (api.Client{}) {
		var errs []error = nil

		config := configuration.GetConfig()

		errs, leoClient = api.NewClient(config.LinguaLeo.Email, config.LinguaLeo.Password)
		if errs != nil {
			log.Fatalf("%v \n", errs)
		}
	}
	return leoClient
}

func GetTranslations(word string) api.Word {
	errs, translations := getClient().GetTranslations(word)
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	return translations
}

func AddWordWithTranslation(word string, translation string) []error {
	errs, _ := getClient().AddWord(word, translation)
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	return errs
}

func AddWord(word string) {
	translations := GetTranslations(word)
	if len(translations.Translations) == 0 {
		log.Fatalln("Translation not found for word \"" + word + "\"")
	}
	translation := translations.Translations[0].Value
	errs := AddWordWithTranslation(word, translation)
	if errs != nil {
		log.Fatalf("Cannot add word: %v", errs)
	}
}
