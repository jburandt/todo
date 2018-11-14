package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Track list items",
	Long:  "Used to quickly track list items",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Basic commands are: list, add, del")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
