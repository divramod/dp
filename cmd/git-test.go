/*Package cmd ... */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	dgit "dp/utils/git"
)

// gitTestCmd represents the gitTest command
var gitTestCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitTest called")
		runGitTest()
	},
}

func init() {
	gitCmd.AddCommand(gitTestCmd)
}

func runGitTest() {
	// tags, err := dgit.TagsGet()
	// if err == nil {
		// for i, s := range tags {
			// fmt.Println(i, s)
		// }
	// }
	branches, err := dgit.BranchesGet("https://github.com/divramod/dp-tpl-ansible.git")
	if err == nil {
		for i, s := range branches {
			fmt.Println(i, s)
		}
	}
}
