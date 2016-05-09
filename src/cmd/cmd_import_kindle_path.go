package cmd

import (
	"log"

	"github.com/igrybkov/leosync/src/kindle"
	"github.com/igrybkov/leosync/src/lingualeo"
	"github.com/spf13/cobra"
	"os"
	path "path/filepath"
	"strings"
)

var kindleMountPoint string

var kindlePathCmd = &cobra.Command{
	Use:   "kindle:path",
	Short: "Import from Kindle's vocabulary",
	Run: func(cmd *cobra.Command, args []string) {

		const VOCABULARY_SUBPATH string = "system/vocabulary/vocab.db"

		kindleMountPoint = strings.TrimSpace(kindleMountPoint)

		if kindleMountPoint == "" {
			log.Fatalln("Path to Kindle should not be empty")
		}

		database := path.Join(kindleMountPoint, VOCABULARY_SUBPATH)
		if _, err := os.Stat(database); os.IsNotExist(err) {
			log.Fatalln("Cannot find vocabulary in " + database)
		}

		words := kindle.GetWords(database)
		for _, word := range words {
			lingualeo.AddWord(word.Word)
			log.Println("Imported: " + word.Word)
		}
	},
}

func init() {
	importCmd.AddCommand(kindlePathCmd)
	kindlePathCmd.Flags().StringVarP(&kindleMountPoint, "path", "p", "", "Path to Kindle")
}
