package dpread

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "github.com/manifoldco/promptui"
)

// UserInput returns the user input for a give question.
func UserInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question)
	userInput, _ := reader.ReadString('\n')
	return userInput
}

// YesNo asks a user for confirmation.
func YesNo(question string) bool {
	prompt := promptui.Select{
		Label: question + " [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}
