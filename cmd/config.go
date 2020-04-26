//Package cmd ...
package cmd

import (
	dlog "dp/utils/log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "A brief description of your command",
	Aliases: []string{"c"},
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			dlog.Err(err.Error())
		}
		editCmd := exec.Command("vim", home+"/.config/config.yaml")
		editCmd.Stdout = os.Stdout
		editCmd.Stdin = os.Stdin
		editCmd.Stderr = os.Stderr
		editCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
