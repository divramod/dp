package dptime

import (
	"time"
)

// GetNow returns the current date
func GetNow() string {
	t := time.Now()
	return t.Format("2006.01.02 15:04:05")
}
