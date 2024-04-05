package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/cmd/version"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kobra",
	Short: "An go demo with gorm gin gen etc.",
	Long:  "kobra is an example used by Cui ChangHe to practice gorm, gin, and gen.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Welcome to use kobra!")
	},
}

func init() {
	rootCmd.AddCommand(version.VersionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
