package dpprint

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// PrintHeader1 prints a string as header of level one for the cli.
func PrintHeader1(headerTxt string) {
	termWidth, termHeight, err := terminal.GetSize(0)
	headerLen := 80
	if err != nil {
		fmt.Println("err", err)
	} else {
		_ = termHeight
		if termWidth < 80 {
			headerLen = termWidth
		}
	}
	equalSigns := ""
	for i := 0; i < headerLen; i++ {
		equalSigns = equalSigns + "="
	}
	headerTxtLen := len(headerTxt)
	headerRestLen := (headerLen - headerTxtLen) / 2
	for i := 0; i < headerRestLen-1; i++ {
		headerTxt = " " + headerTxt + " "
	}
	allLen := headerRestLen*2 + headerTxtLen
	if allLen < termWidth {
		headerTxt = headerTxt + " "
	}
	headerTxt = "|" + headerTxt + "|"
	fmt.Println(equalSigns)
	fmt.Println(headerTxt)
	fmt.Println(equalSigns)
	fmt.Println("---", time.Now().Format("2006-01-02 15:04:05.000000"), "---")
}
