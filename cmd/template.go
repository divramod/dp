/*Package cmd bla */
package cmd

import (
	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Aliases: []string{"t", "tpl"},
	Short: "",
	Long: ``,
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
