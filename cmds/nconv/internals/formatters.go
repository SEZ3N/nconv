package internals

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/savioxavier/termlink"
)

func DisplayFormattedErrorMessage(msg string) {
	defer color.Unset()
	formattedErrMessage := color.RedString(msg)
	color.Set(color.Bold)
	fmt.Println(formattedErrMessage)
}

func DisplayFomattedMessageFromErr(err error) {
	defer color.Unset()
	errCol := color.New(color.FgRed, color.Bold)
	errCol.Print(err.Error())
}

func DisplayMessageWithColor(msg string, col color.Attribute, bold bool) {
	defer color.Unset()
	var msgColor color.Color
	msgColor.Add(col)
	if bold {
		msgColor.Add(color.Bold)
	}
	msgColor.Print(msg)
}

func getFormmatedClickableLink(link string) string {
	linkWithProtocol := "file://" + link
	clickable := termlink.Link(link, linkWithProtocol)
	defer color.Unset()
	color.Set(color.Underline)
	fmttedClickable := color.CyanString(clickable)
	return fmttedClickable
}

func DisplayMessageOnSuccess(linkDesc, link string) {
	fmttedLnk := getFormmatedClickableLink(link)
	legend := color.New(color.Bold, color.FgGreen).SprintFunc()
	defer color.Unset()
	fmt.Printf("%s: %s", legend(linkDesc), fmttedLnk)
}
