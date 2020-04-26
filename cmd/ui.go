//Package cmd ...
package cmd

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// uiCmd represents the ui command
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		form := tview.NewForm().
			AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
			AddInputField("First name", "", 20, nil, nil).
			AddInputField("Last name", "", 20, nil, nil).
			AddCheckbox("Age 18+", false, nil).
			AddPasswordField("Password", "", 10, '*', nil).
			AddButton("Save", func() {
				fmt.Println("test")
			}).
			AddButton("Quit", func() {
				app.Stop()
			})
		form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
		flex := tview.NewFlex().
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
				// AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
				AddItem(form, 0, 3, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
		if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}

		// if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		// panic(err)
		// }

	},
}

func init() {
	rootCmd.AddCommand(uiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
