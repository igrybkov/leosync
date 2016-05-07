package cmd

import (
	"github.com/igrybkov/leosync/src/kindle"
	"github.com/igrybkov/leosync/src/lingualeo"
	"github.com/spf13/cobra"
	"log"
)

var kindleDatabaseCmd = &cobra.Command{
	Use:   "kindle:database",
	Short: "Import from path to Kindle root",
	Run: func(cmd *cobra.Command, args []string) {
		words := kindle.GetWords(filePath)
		for _, word := range words {
			lingualeo.AddWord(word.Word)
			log.Println("Imported: " + word.Word)
		}
	},
}

func init() {
	importCmd.AddCommand(kindleDatabaseCmd)
	kindleDatabaseCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
}
