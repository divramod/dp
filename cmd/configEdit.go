//Package cmd ...
package cmd

import (
	dlog "dp/utils/log"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// configEditCmd represents the configEdit command
var configEditCmd = &cobra.Command{
	Use:   "configEdit",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configEdit called")
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
	configCmd.AddCommand(configEditCmd)
}
