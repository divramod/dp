package dprompt

import "github.com/manifoldco/promptui"

// Arr lets the user choose from a array of strings
func Arr(question string, choices []string) (string, error) {
	prompt := promptui.Select{
		Label: question,
		Items: choices,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

// YesNo ask for yes no
func YesNo(question string) (bool, error) {
    prompt := promptui.Select{
		Label: question + " [Yes/No]:",
        Items: []string{"Yes", "No"},
    }
    _, result, err := prompt.Run()
    if err != nil {
		return false, err
    }
    return result == "Yes", nil
}
