package helper

import "fmt"

type Console struct {
}

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func (console Console) Log(messageType string, message string) {
	if messageType == "Debug" {
		fmt.Printf(DebugColor, message)
		fmt.Println("")
	} else if messageType == "Info" {
		fmt.Printf(InfoColor, message)
		fmt.Println("")
	} else if messageType == "Notice" {
		fmt.Printf(NoticeColor, message)
		fmt.Println("")
	} else if messageType == "Error" {
		fmt.Printf(WarningColor, message)
		fmt.Println("")
	} else if messageType == "Warning" {
		fmt.Printf(WarningColor, message)
		fmt.Println("")
	} else {

		fmt.Println(message)
	}
}
