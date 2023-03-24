package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "askcloudhelloworld",
	Short: "A basic Hello World example",
	Long: `
	A longer description on how this works
	example: askcloudhelloworld
	example: askcloudhelloworld -d
	`,
	Run: func(cmd *cobra.Command, args []string) {
		flagVal, err := cmd.Flags().GetBool("differentmessage")
		if err != nil {
			return
		}
		if flagVal {
			fmt.Println("I like these words better!")
			return
		}
		fmt.Println("Hello World v0.1.1!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("differentmessage", "d", false, "Toggle a different message")
}
