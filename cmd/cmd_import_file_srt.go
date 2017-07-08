package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/icza/srtgears"
	"github.com/spf13/cobra"
)

var importFileSrtCmd = &cobra.Command{
	Use:   "srt",
	Short: "Import from SRT file",
	Run: func(cmd *cobra.Command, args []string) {

		subsPack, err := srtgears.ReadSrtFile(filePath)

		subsPack.RemoveControl()
		subsPack.RemoveHI()
		subsPack.RemoveHTML()

		allSubsSet := []string{}

		for _, subs := range subsPack.Subs {
			allSubsSet = append(allSubsSet, strings.Join(subs.Lines, " "))
		}

		//allText := strings.Join(allSubsSet, " ")

		//sentenceTokenizer, err := english.NewSentenceTokenizer(nil)
		//if err != nil {
		//	panic(err)
		//}

		//sentencesSet := sentenceTokenizer.Tokenize(allText)
		//for _, s := range sentencesSet {
		//fmt.Println(strings.TrimSpace(s.Text))
		//}

		leo := getLeoClient()
		wl, err := leo.GetList()
		if err != nil {
			log.Fatalln(err)
		}

		for _, s := range wl {
			fmt.Println(strings.TrimSpace(s.Word))
		}

		//wordTokenizer := english.NewWordTokenizer(sentences.NewPunctStrings())
		//wordsSet := wordTokenizer.Tokenize(allText, false)
		//for _, s := range wordsSet {
		//	fmt.Println(strings.TrimSpace(s.Tok))
		//}

		log.Fatalln("lol")

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		r := csv.NewReader(file)
		r.Comma = ';'
		var word string

		//conf := settings.GetSettings()
		//leoClient, err := lingualeo.New(conf.LinguaLeo)

		for {
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			func() {
				err = func() (err error) {
					//err, result := leoClient.AddWord(
					//	lingualeo.AddWordRequest{
					//		Word:        word,
					//		Translation: translation,
					//		Context:     context,
					//	},
					//)
					//
					//if err != nil {
					//	return err
					//}
					//
					//log.Println("Imported: " + word + " = " + translation)
					//
					//imgURL := strings.TrimSpace(record[2])
					//if len(imgURL) > 5 {
					//	lingualeo.DownloadPicture(imgURL, strconv.Itoa(result.TranslateID))
					//	log.Println("+picture: " + imgURL)
					//	time.Sleep(1 * time.Second) //anti-ban delay :)
					//}
					return err
				}()
				if err != nil {
					log.Println("Can't add the word \"" + word + "\": " + err.Error())
				}
			}()
		}

	},
}

func init() {
	importCmd.AddCommand(importFileSrtCmd)
	importFileSrtCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")

}
