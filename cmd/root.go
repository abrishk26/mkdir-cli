package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// mkdirCmd represents the mkdir command
var mkdirCmd = &cobra.Command{
	Use: "mkdir",
	Short: "Creates a new directory.",
	Long: "This command allows you to create a new directory from the command line.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: you must specify a directory name")
			return
		}

		dirName := args[0]
		err := os.Mkdir(dirName, 0755)

		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		fmt.Println("Directory created:", dirName)
	},
}

func init() {
	mkdirCmd.Root().AddCommand(mkdirCmd)
}