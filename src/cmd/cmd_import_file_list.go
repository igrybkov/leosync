package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"

	leo "github.com/igrybkov/leosync/src/lingualeo"
	"github.com/spf13/cobra"
)

var importFileListCmd = &cobra.Command{
	Use:   "list",
	Short: "Import from file with list of words (one word per line)",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var word string

		for scanner.Scan() {
			word = strings.TrimSpace(scanner.Text())
			leo.AddWord(word)
			log.Println("Imported: " + word)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	importCmd.AddCommand(importFileListCmd)
	importFileListCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")

}
