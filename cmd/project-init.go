package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	dpprint "dp/utils/print"
	dpread "dp/utils/read"
	dpfs "dp/utils/fs"
	dptime "dp/utils/time"
)

var projectName string
var projectPath string
var projectType string
var createAll bool
var createTmuxpConf bool
var initGitRepo bool
var skipConfirmProjectCreation bool
var resetExistingProject bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initializes a project.",
	Aliases: []string{"i", "ini"},
	Long:    `The process of initializing a project should save you time for repeated tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	projectCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "Pass a project name.")
	initCmd.Flags().StringVarP(&projectPath, "path", "p", "", "Pass a project path.")
	initCmd.Flags().StringVarP(&projectType, "type", "y", "", "Pass a project type.")
	initCmd.Flags().BoolVarP(&createAll, "create-all", "a", false, "Define if all base things should be created (default: false).")
	initCmd.Flags().BoolVarP(&createTmuxpConf, "create-tmuxp-conf", "t", false, "Define if a tmuxp configuration file should be created (default: false).")
	initCmd.Flags().BoolVarP(&initGitRepo, "create-git-repo", "g", false, "Define if a git repository should be created (default: false).")
	initCmd.Flags().BoolVarP(&skipConfirmProjectCreation, "skip-confirm-project-creation", "s", false, "Avoid asking for confirmation for project creation.")
	initCmd.Flags().BoolVarP(&resetExistingProject, "reset-existing-project", "r", false, "Overwrite choosen project directory, if already existant.")
}

func run() (bool, string) {
	// --- log
	dpprint.PrintHeader1("[dp project init] Create a project")
	// --- define project name
	if projectName == "" {
		projectName = dpread.UserInput("Please enter a project name: ")
	}
	// --- define project path
	if projectPath == "" {
		projectPath = dpread.UserInput("Please enter a project path: ")
	}
	projectPath = filepath.Join(".", ".tmp", projectPath)
	// --- set create all
	if createAll {
		createTmuxpConf = true
		initGitRepo = true
	}
	// --- project create
	projectCreatedSuccess, projectCreatedMsg := projectCreate()
	fmt.Println(projectCreatedMsg) 
	if !projectCreatedSuccess {
		return false, projectCreatedMsg
	}
	// --- choose project type

	// --- choose elements 

	// --- return 
	return true, "Project successfully created!"
}

func projectSelectType() {

}

func projectCreate() (bool, string) {
	fmt.Println("- project name:", projectName)
	fmt.Println("- project path:", projectPath)
	fmt.Println("- create tmuxp conf:", createTmuxpConf)
	fmt.Println("- init git repo:", initGitRepo)
	fmt.Println("time:", dptime.GetNow()) 
	createProject := false
	if skipConfirmProjectCreation {
		createProject = true
	} else {
		createProject = dpread.YesNo("Confirm project creation!")
	}
	if !createProject {
		return false, "Project creation aborted by the user!"
	}
	projectPathExists, err :=  dpfs.ExistsOld(projectPath)
	if err != nil {
		return false, "Project creation aborted! Problem checking path existance!"
	} 
	if projectPathExists {
		if !resetExistingProject {
			createProject = dpread.YesNo("Project path '"+projectPath+"' already existant! Reset it?")
			if !createProject {
				return false, "Project path '"+projectPath+"' already existant! Project creation aborted by you!"
			}
			os.Remove(projectPath)
		} else {
			os.Remove(projectPath)
		}
	}
	os.Mkdir(projectPath, os.ModeDir)
	return true, "Project '"+projectName+"' created at path '"+projectPath+"'!"
}
