package cmd

import (
	"encoding/csv"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	leo "github.com/igrybkov/leosync/src/lingualeo"
)

var importFileCsvCmd = &cobra.Command{
	Use:   "csv",
	Short: "Import from CSV file ('word;translation' per line)",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

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

			pic_url := strings.TrimSpace(record[2])
			if len(pic_url) > 5 {
				leo.DownloadPicture(pic_url, strconv.Itoa(result.TranslateId))
				log.Println("+picture: " + pic_url)
				time.Sleep(1 * time.Second) //anti-ban delay :)
			}

		}

	},
}

func init() {
	importCmd.AddCommand(importFileCsvCmd)
	importFileCsvCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")

}
