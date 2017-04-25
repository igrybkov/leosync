package cmd

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	leo "github.com/igrybkov/leosync/src/lingualeo"
	"github.com/spf13/cobra"
)

var importFileSrtCmd = &cobra.Command{
	Use:   "csv",
	Short: "Import from CSV file ('word;translation' per line)",
	Run: func(cmd *cobra.Command, args []string) {
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
		var translation string
		var context string

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			word = strings.TrimSpace(record[0])
			translation = strings.TrimSpace(record[1])
			context = strings.TrimSpace(record[4])

			_, result := leo.AddWordWithTranslationAndContext(word, translation, context)
			log.Println("Imported: " + word + " = " + translation)

			if context != "" {
				log.Println("+context: " + context)
			}

			imgURL := strings.TrimSpace(record[2])
			if len(imgURL) > 5 {
				leo.DownloadPicture(imgURL, strconv.Itoa(result.TranslateID))
				log.Println("+picture: " + imgURL)
				time.Sleep(1 * time.Second) //anti-ban delay :)
			}

		}

	},
}

func init() {
	importCmd.AddCommand(importFileSrtCmd)
	importFileSrtCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")

}
