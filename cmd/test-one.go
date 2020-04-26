/*Package cmd bla*/
package cmd

import (
	// dfind "dp/utils/find"
	// dgit "dp/utils/git"
	// "fmt"

	dlog "dp/utils/log"
	dpprint "dp/utils/print"
	"fmt"
	"os"

	// dprompt "dp/utils/prompt"
	dos "dp/utils"
	// "fmt"

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
		// fmt.Println("hello")
		// result, _ := dprompt.YesNo("Do you want to?")
		// fmt.Println("result:", result)

		os.Setenv("FOO", "1")
		err, stdOut, stdErr := dos.Exec("hcloud server list; echo $FOO")
		if err != nil {
			dlog.Err(err.Error())
		}
		fmt.Print("stdErr:", stdErr)
		fmt.Print("stdOut:", stdOut)

	},
}

func init() {
	testCmd.AddCommand(oneCmd)
}
