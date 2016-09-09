package cmd

import (
	//"os"

	"bufio"
	"fmt"
	"github.com/igrybkov/leosync/src/configuration"
	home "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		config := configuration.GetConfig()

		file := cfgFile

		if strings.TrimSpace(file) == "" {
			file = viper.ConfigFileUsed()
			if strings.TrimSpace(file) == "" {
				log.Fatalf("Configuration file not found")
			}
		}
		file, err = home.Expand(file)
		if err != nil {
			log.Fatalln(err.Error())
		}

		reader := bufio.NewReader(os.Stdin)
		var input string

		getInputValue := func(fieldName string, currentValue string) string {
			fmt.Print("Enter " + fieldName + " (leave empty to use current value: '" + currentValue + "'): ")
			input, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			input = strings.TrimSpace(input)
			result := currentValue
			if input != "" {
				result = input
			}
			return result
		}

		config.LinguaLeo.Email = getInputValue("LinguaLeo email", config.LinguaLeo.Email)
		config.LinguaLeo.Password = getInputValue("LinguaLeo password", config.LinguaLeo.Password)

		b, err := yaml.Marshal(config)
		if err != nil {
			log.Fatalln(err.Error())
		}

		f, err := os.Create(file)
		if err != nil {
			log.Fatalln(err.Error())
		}

		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		_, err = f.WriteString(string(b))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
