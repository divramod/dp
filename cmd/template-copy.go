/*Package cmd bla*/
package cmd

import (
	dfind "dp/utils/find"
	dpfs "dp/utils/fs"
	dgit "dp/utils/git"
	dlog "dp/utils/log"
	dpprint "dp/utils/print"
	dprompt "dp/utils/prompt"
	"errors"
	"os"

	// "errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// VARS
// ----------------------------------------------------------------------------
var copyCmd = &cobra.Command{
	Use:     "copy",
	Aliases: []string{"c", "cp"},
	Short:   "A brief description of your command",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run")
		fmt.Println("overwritePath:", overwritePath)
		runTemplateCopy()
	},
}

// TYPES
// ----------------------------------------------------------------------------

// Path template
type Path struct {
	Name  string
	Type  string
	Value string
}

// Template test
type Template struct {
	Name        string
	Description string
	Pathes      []Path
}

// FUNCTIONS
// ----------------------------------------------------------------------------
var gitBranch string
var namePath string
var nameTemplate string
var overwritePath bool
var pathDestination string
var pathSource string

// --- init()
func init() {
	fmt.Println("init")
	copyCmd.Flags().StringVarP(&nameTemplate, "name-template", "t", "", "Pass a template name.")
	copyCmd.Flags().StringVarP(&namePath, "name-path", "p", "", "Pass a path name.")
	copyCmd.Flags().StringVarP(&pathDestination, "path-des", "d", "", "Pass a destination path.")
	copyCmd.Flags().StringVarP(&pathSource, "path-src", "s", "", "Pass the name of the path to copy from.")
	copyCmd.Flags().StringVarP(&gitBranch, "git-branch", "b", "", "Pass the name of the branch you want to copy from.")
	copyCmd.Flags().BoolVarP(&overwritePath, "overwrite-if-existant", "o", false, "Overwrite path if already existant.")
	templateCmd.AddCommand(copyCmd)
}

func runTemplateCopy() {

	// --- [log] header
	dpprint.PrintHeader1("Copy template")

	// --- [set] template
	var templates []Template
	err := viper.UnmarshalKey("templates", &templates)
	if err != nil {
		panic("Unable to unmarshal templates")
	}

	// --- [set] choose template
	var template Template
	fmt.Println("nameTemplate:", nameTemplate)
	if nameTemplate != "" {
		found := false
		var templateNo int
		for i := range templates {
			if templates[i].Name == nameTemplate {
				found = true
				templateNo = i
			}
		}
		if found == false {
			dlog.Err("Template" + nameTemplate + "not found!")
			return
		}
		dlog.Inf("Chose template from flag!")
		template = templates[templateNo]
	} else {
		templateNo, err := promptChooseTemplate(templates)
		if err != nil {
			dlog.Err(err.Error())
			return
		}
		template = templates[templateNo]
	}

	// --- [set] pathes
	var path Path

	if pathSource != "" {
		for _, searchPath := range template.Pathes {
			if searchPath.Name == pathSource {
				path = searchPath
			}
		}
		if path.Name == "" {
			dlog.Err("The given path '" + pathSource + "' could not be found!")
			return
		}
	} else {
		if len(template.Pathes) == 0 {
			dlog.Err("No local path or git path defined.")
			return
		} else if len(template.Pathes) == 1 {
			path = template.Pathes[0]
		} else {
			dlog.Inf("Choose path")
			pathNo, err := promptChooseBranch(template.Pathes)
			if err != nil {
				dlog.Err(err.Error())
				return
			}
			path = template.Pathes[pathNo]
		}
	}

	// --- [get] template
	if path.Type != "git" && path.Type != "local" {
		dlog.Err("Unknown path type '" + path.Type + "'! Choose one of [git|local]!")
		return
	} else if path.Type == "local" {
		err = copyLocal(path)
		if err != nil {
			dlog.Err(err.Error())
			return
		}
	} else if path.Type == "git" {
		err = copyGit(path)
		if err != nil {
			dlog.Err(err.Error())
			return
		}
	}

	// --- finish
	dlog.Suc("created folder from template")
}

func copyGit(path Path) error {
	fmt.Println("path:", path)

	// --- [get] branches
	branches, err := dgit.BranchesGet(path.Value)
	if err != nil {
		return err
	}

	// --- [check] if branch was passed
	var branch string
	if gitBranch != "" {
		branch = gitBranch
		// --- TODO [check] if branch exists
	} else {
		// --- [check] if more than one branch exists
		if len(branches) == 0 {
			return errors.New("no branch existant")
		} else if len(branches) == 1 {
			branch = branches[0]
		} else {
			// --- [choose] branch
			promptBranches := promptui.Select{
				Label: "Select branch from " + path.Value,
				Items: branches,
			}
			_, result, err := promptBranches.Run()
			if err != nil {
				return err
			}
			branch = result
		}
	}

	pathChoosen, err := choosePath(path)
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	pathTmpTarget := home + "/.tmp/tmp_templates/" + pathChoosen
	err = dgit.BranchClone(map[string]string{
		"URL":       path.Value,
		"Branch":    branch,
		"Directory": pathTmpTarget,
	})
	if err != nil {
		return err
	}
	os.RemoveAll(pathTmpTarget + "/.git")


	whatToCopy, err := dprompt.Arr("Choose what to copy", []string{"All", "File", "Directory"})
	if whatToCopy == "File" {
		files, _ := dfind.FindFiles(pathTmpTarget)
		if len(files) > 0 {
			dpfs.CreateIfNotExists(pathChoosen, 0777)
			for i := range files {
				file := files[i]
				pathTarget := pathChoosen + "/" + file.PathLocal
				dpfs.CreateIfNotExistsForFile(pathTarget, 0777)
				dpfs.Copy(file.PathGlobal, pathTarget)
			}
		}
	} else if whatToCopy == "Directory" {
		dirs, _ := dfind.FindDirs(pathTmpTarget)
		if len(dirs) > 0 {
			dpfs.CreateIfNotExists(pathChoosen, 0777)
			for i := range dirs {
				dir := dirs[i]
				pathTarget := pathChoosen + "/" + dir.PathLocal
				fmt.Println("dir.PathGlobal:", dir.PathGlobal)
				fmt.Println("pathTarget:", pathTarget)
				dpfs.CreateIfNotExists(pathTarget, 0777)
				dpfs.CopyDirectory(dir.PathGlobal, pathTarget)
			}
		}
	} else if whatToCopy == "All" {
		os.Rename(pathTmpTarget, pathChoosen)
	}




	// --- [return]
	return nil
}

func copyLocal(path Path) error {

	// --- [check] path for existance
	exists, err := dpfs.ExistsOld(path.Value)
	if err != nil {
		return err
	} else if exists != true {
		msg := "path to copy '" + path.Value + "' not existant"
		return errors.New(msg)
	}

	// --- [copy]
	pathChoosen, err := choosePath(path)
	if err != nil {
		return err
	}

	// --- [check] path existant?

	// --- [copy]
	dlog.Deb("copy path " + pathChoosen)
	dpfs.CopyDirectory(path.Value, pathChoosen)

	// --- [return]
	return nil
}

func choosePath(path Path) (string, error) {

	// --- [set] path (absolute or relative?)
	dlog.Inf("Some info")

	// --- [check] path existant
	var pathToCheck string
	if pathDestination != "" {
		pathToCheck = pathDestination
	} else {
		prompt := promptui.Prompt{
			Label: "Path",
		}
		result, err := prompt.Run()
		if err != nil {
			return "", err
		}
		pathToCheck = result
		fmt.Println("pathToCheck:", pathToCheck)
	}

	// --- [check] if pathToCheck existant
	exists, err := dpfs.ExistsOld(pathToCheck)
	if err != nil {
		return pathToCheck, err
	} else if exists == true {
		msg := "Path '" + pathToCheck + "' already existant!"
		return pathToCheck, errors.New(msg)
	}

	// --- [return]
	return pathToCheck, nil
}

func promptChooseTemplate(templates []Template) (int, error) {
	templateSearcher := func(input string, index int) bool {
		templateS := templates[index]
		name := strings.Replace(strings.ToLower(templateS.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}
	templatesChoose := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .Description | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Description | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- Template ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }} {{ .Description }}
`,
	}
	// --- [choose] template
	prompt := promptui.Select{
		Label:     "Select Template",
		Items:     templates,
		Searcher:  templateSearcher,
		Templates: templatesChoose,
	}
	i, _, err := prompt.Run()
	if err != nil {
		return 100000, err
	}
	return i, nil
}

// wait for making this dynamic:
// - https://github.com/manifoldco/promptui/issues/150
func promptChooseBranch(pathes []Path) (int, error) {

	pathSearcher := func(input string, index int) bool {
		pathS := pathes[index]
		typeP := strings.Replace(strings.ToLower(pathS.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(typeP, input)
	}

	pathesChoose := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 [{{ .Type | yellow }}] {{ .Name | cyan }} ({{ .Value | red }})",
		Inactive: "[{{ .Type | yellow }}] {{ .Name | cyan }} ({{ .Value | red }})",
		Selected: "\U0001F336 [{{ .Type | yellow }}] {{ .Name | red | cyan }}",
		Details: `
--------- Path ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Type:" | faint }}	{{ .Type }}
{{ "Value:" | faint }}	{{ .Value }}
`,
	}
	prompt := promptui.Select{
		Label:     "Select Path",
		Items:     pathes,
		Searcher:  pathSearcher,
		Templates: pathesChoose,
	}
	i, _, err := prompt.Run()
	if err != nil {
		return 100000, err
	}
	return i, nil
}