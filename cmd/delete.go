package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func init() {
	RootCmd.AddCommand(delCmd)
}

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete items",
	Long:  "delete item from list",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		homeDir := usr.HomeDir
		dirCheck := strings.Join([]string{homeDir, "/.todo"}, "")
		configFile := strings.Join([]string{dirCheck, "/config"}, "")
		tmpConfig := strings.Join([]string{dirCheck, "/tmpconfig"}, "")
		bakConfig := strings.Join([]string{dirCheck, "/old.bak"}, "")

		//if CheckExists(configFile) == false {
		//	CreateConfig(dirCheck)
		//}

		output, err := ReadConfig(configFile)
		if err != nil {
			fmt.Println("Could not read config file. Please run the add command to create a list.")
			os.Exit(1)
		}

		if len(os.Args) <= 2 {
			fmt.Println("Please specify a line number to delete")
			os.Exit(1)
		}

			lineDel := os.Args[2]
			del, err := strconv.Atoi(lineDel)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			d := del - 1

			output = append(output[:d], output[d+1:]...)
			os.Create(tmpConfig)

			var j, r int
			for r = range output {
				for j = 1; j <= r; j++ {
				}
				WriteConfig(output, tmpConfig)
			}
			os.Rename(configFile, bakConfig)
			os.Rename(tmpConfig, configFile)
			os.Remove(bakConfig)
	},
}