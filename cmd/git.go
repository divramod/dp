//Package cmd bla
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git called")
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
}
