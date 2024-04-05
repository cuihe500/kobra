package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kobra",
	Long:  "All software has versions. This is kobra's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("kobra Version:\t", "v0.0.1_2024_4_5")
	},
}
