//Package dos ...
package dos

import (
	"bytes"
	"os/exec"
)

// ShellToUse
var ShellToUse string = "sh"

// Exec ...
// * https://stackoverflow.com/questions/6182369/exec-a-shell-command-in-go/7786922#7786922
func Exec(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
