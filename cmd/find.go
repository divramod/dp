//Package cmd ...
package cmd

import (
	dfind "dp/utils/find"
	"fmt"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files, _ := dfind.FindFiles(".")
		for i := range files {
			file := files[i]
			fmt.Println("file:", file.PathGlobal)
			fmt.Println("file:", file.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
