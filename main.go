package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"yell/internal/ui"
)

func readPipedInput() string {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return ""
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := io.ReadAll(os.Stdin)
		if err == nil {
			return strings.TrimSpace(string(bytes))
		}
	}
	return ""
}

func main() {
	defaultMsg := "Shawty lemme holla at you"
	title := flag.String("title", "Somebody yelling", "Title of the alert popup")
	message := flag.String("message", defaultMsg, "Message text to display")
	icon := flag.String("icon", "🗣️", "Icon or emoji on the left")
	flag.Parse()

	if piped := readPipedInput(); piped != "" {
		if *message == defaultMsg {
			*message = piped
		}
	}

	ui.Build(title, message, icon)
}
