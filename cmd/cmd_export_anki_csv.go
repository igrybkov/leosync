package cmd

//
//import (
//	"encoding/csv"
//	"log"
//	"os"
//
//	"github.com/spf13/cobra"
//)
//
//var exportAnkiCsvCmd = &cobra.Command{
//	Use:   "anki:csv",
//	Short: "Export data from LinguaLeo to CSV file that may be imported to Anki",
//	Run: func(cmd *cobra.Command, args []string) {
//
//		var data = [][]string{{"Line1", "Hello Readers of:"}, {"Line2", "golangcode.com"}}
//
//		file, err := os.Create(filePath)
//		checkError("Cannot create file", err)
//		defer file.Close()
//
//		writer := csv.NewWriter(file)
//		writer.Comma = ';'
//
//		for _, value := range data {
//			err := writer.Write(value)
//			checkError("Cannot write to file", err)
//		}
//
//		defer writer.Flush()
//	},
//}
//
//func checkError(message string, err error) {
//	if err != nil {
//		log.Fatal(message, err)
//	}
//}
//
//func init() {
//	exportCmd.AddCommand(exportAnkiCsvCmd)
//	exportAnkiCsvCmd.Flags().StringVarP(&filePath, "file", "f", "", "CSV file path")
//}
