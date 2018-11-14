package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/user"
	"strings"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list items",
	Long:  "numbered list of items",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		homeDir := usr.HomeDir
		dirCheck := strings.Join([]string{homeDir, "/.todo"}, "")
		configFile := strings.Join([]string{dirCheck, "/config"}, "")

		if CheckExists(configFile) == false {
			CreateConfig(dirCheck)
		}

		output, err := ReadConfig(configFile)
		if err != nil {
			fmt.Println("Could not read config file.")
			return
		}

		var j, r int
		for r = range output {
			for j = 1; j <= r; j++ {
			}
			fmt.Println(j, output[r])
		}
	},
}

