package lingualeo

import (
	"github.com/igrybkov/leosync/config"
	"github.com/igrybkov/leosync/src/lingualeo/response"
	"log"
)

// LinguaLeo represents api of LinguaLeo
type LinguaLeo struct {
	leoClient apiClient
}

func (l *LinguaLeo) getClient() apiClient {
	if l.leoClient == (apiClient{}) {
		var errs []error

		conf := config.GetConfig()

		errs, leoClient := newClient(conf.LinguaLeo.Email, conf.LinguaLeo.Password)
		if errs != nil {
			log.Fatalf("%v \n", errs)
		}
		l.leoClient = leoClient
	}
	return l.leoClient
}

// GetTranslations return a translation for the word
func (l *LinguaLeo) GetTranslations(word string) response.Word {
	errs, translations := l.getClient().GetTranslations(word)
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	return translations
}

// AddWordWithTranslation add a word with already defined translation
func (l *LinguaLeo) AddWordWithTranslation(word string, translation string) []error {
	errs, _ := l.getClient().AddWord(word, translation)
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	return errs
}

// AddWord adds a word with first proposed translation
func (l *LinguaLeo) AddWord(word string) {
	translations := l.GetTranslations(word)
	if len(translations.Translations) == 0 {
		log.Fatalln("Translation not found for word \"" + word + "\"")
	}
	translation := translations.Translations[0].Value
	errs := l.AddWordWithTranslation(word, translation)
	if errs != nil {
		log.Fatalf("Cannot add word: %v", errs)
	}
}

// AddWordWithTranslationAndContext adds a word with a context
func (l *LinguaLeo) AddWordWithTranslationAndContext(word string, translation string, context string) ([]error, response.Word) {
	errs, result := l.getClient().AddWordWithContext(word, translation, context)
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	return errs, result
}

// DownloadPicture posts a picture to the translation
func (l *LinguaLeo) DownloadPicture(url string, translateID string) {
	errs := l.getClient().DownloadPicture(url, translateID)
	if errs != nil {
		log.Fatalf("Cannot set picture: %v", errs)
	}
}
