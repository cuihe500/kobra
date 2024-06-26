package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.eaip.top/gorm-gen-gin-learn-project/internal/common"
)

var VCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kobra",
	Long:  "All software has versions. This is kobra's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("kobra Version:\t", common.Version, "\nWritten for love and peace.", "\nBest wishes to Xu QianQian.\nI hope you have me in your dreams tonight.")
	},
}
