package dlog

import "github.com/fatih/color"

// Inf logs a colored info string
func Inf(msg string) {
	cInf := color.New(color.FgBlue)
	cInf.Println("[INF]", msg)
}

// Err logs a colored error string
func Err(msg string) {
	cErr := color.New(color.FgRed)
	cErr.Println("[Err]", msg)
}

// Suc logs a colored success string
func Suc(msg string) {
	cErr := color.New(color.FgCyan)
	cErr.Println("[SUC]", msg)
}

// Deb logs a colored debug string
func Deb(msg string) {
	cDeb := color.New(color.FgMagenta)
	cDeb.Println("[DEB]", msg)
}

// War logs a colored debug string
func War(msg string) {
	cWar:= color.New(color.FgYellow)
	cWar.Println("[WAR]", msg)
}
