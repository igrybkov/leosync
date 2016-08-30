package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"io"
	"strings"

	leo "github.com/igrybkov/leosync/src/lingualeo"
	"github.com/spf13/cobra"
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
		leo.AddWordWithTranslation(word, translation)
		log.Println("Imported: " + word + " = " + translation)
		}
	
	},
}

func init() {
	importCmd.AddCommand(importFileCsvCmd)
	importFileCsvCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")

}
