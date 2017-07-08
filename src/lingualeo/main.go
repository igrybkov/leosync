package lingualeo

//import (
//	"log"
//
//	"github.com/igrybkov/leosync/src/lingualeo/api"
//	"github.com/igrybkov/leosync/src/settings"
//)
//
//func getClient() api.Client {
//	if leoClient == (api.Client{}) {
//		var errs []error
//
//		config := settings.GetConfig()
//
//		errs, leoClient = api.NewClient(config.LinguaLeo.Email, config.LinguaLeo.Password)
//		if errs != nil {
//			log.Fatalf("%v \n", errs)
//		}
//	}
//	return leoClient
//}
//
//// GetTranslations return a translation for the word
//func GetTranslations(word string) api.Word {
//	errs, translations := getClient().GetTranslations(word)
//	if errs != nil {
//		log.Fatalf("%v \n", errs)
//	}
//	return translations
//}
//
//// AddWordWithTranslation add a word with already defined translation
//func AddWordWithTranslation(word string, translation string) []error {
//	errs, _ := getClient().AddWord(word, translation)
//	if errs != nil {
//		log.Fatalf("%v \n", errs)
//	}
//	return errs
//}
//
//// AddWord adds a word with first proposed translation
//func AddWord(word string) {
//	translations := GetTranslations(word)
//	if len(translations.Translations) == 0 {
//		log.Fatalln("Translation not found for word \"" + word + "\"")
//	}
//	translation := translations.Translations[0].Value
//	errs := AddWordWithTranslation(word, translation)
//	if errs != nil {
//		log.Fatalf("Cannot add word: %v", errs)
//	}
//}
//
//// AddWordWithTranslationAndContext adds a word with a context
//func AddWordWithTranslationAndContext(word string, translation string, context string) ([]error, api.Word) {
//	errs, result := getClient().AddWordWithContext(word, translation, context)
//	if errs != nil {
//		log.Fatalf("%v \n", errs)
//	}
//	return errs, result
//}
//
//// DownloadPicture posts a picture to the translation
//func DownloadPicture(url string, translateID string) {
//	errs := getClient().DownloadPicture(url, translateID)
//	if errs != nil {
//		log.Fatalf("Cannot set picture: %v", errs)
//	}
//}
