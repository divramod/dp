package dpterm

import "fmt"

// Clear returns the user input for a give question.
// * https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
func Clear() {
	fmt.Println("\033[2J")
}
