//Package cmd ...
package cmd

import (
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// uiTestCmd represents the uiTest command
var uiTestCmd = &cobra.Command{
	Use:   "uiTest",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runUiTest()
	},
}

func runUiTest() {
		box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
		if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
			panic(err)
		}
}

func init() {
	uiCmd.AddCommand(uiTestCmd)
}
