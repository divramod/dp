/*Package cmd bla*/
package cmd

import (
	// dfind "dp/utils/find"
	// dgit "dp/utils/git"
	// "fmt"

	dpprint "dp/utils/print"
	"fmt"

	"github.com/spf13/cobra"
)

// oneCmd represents the one command
var oneCmd = &cobra.Command{
	Use:   "one",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dpprint.PrintHeader1("dp test one")
		// dpfs.CreateIfNotExistsForFile("here/is/something/README.md", 0777)
		// dirs, _ := dfind.FindDirs(".")
		// for i := range dirs {
		// dir := dirs[i]
		// fmt.Println("dir:", dir)
		// }

		// --- [check] ssh
		// pubKey := dgit.SSHKeyGet()
		// fmt.Println("pubKey:", pubKey)
		fmt.Println("hello")
	},
}

func init() {
	testCmd.AddCommand(oneCmd)
}
