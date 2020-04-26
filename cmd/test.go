/*Package cmd ...*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `Long`,
	Run: func(cmd *cobra.Command, args []string) {
		runTest()
	},
}

func runTest() {
	fmt.Println("cfgFile:", cfgFile) 
}

func init() {
	rootCmd.AddCommand(testCmd)
}
