package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"yell/internal/ui"
	"yell/internal/ui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

	myApp := app.NewWithID("com.yell.popup")

	fyne.Do(func() {})

	myApp.Settings().SetTheme(&ui.ToastTheme{})

	var myWindow fyne.Window
	if drv, ok := myApp.Driver().(desktop.Driver); ok {
		myWindow = drv.CreateSplashWindow()
	} else {
		myWindow = myApp.NewWindow(*title)
	}

	// Emoji Icon
	emojiObject := ui.GetEmbeddedEmojiImage(*icon, 42)
	iconContainer := container.NewCenter(emojiObject)

	// Selectable Title (20pt Bold via RichText Heading) & Message
	titleLabel := widget.NewRichText(
		&widget.TextSegment{
			Text: *title,
			Style: widget.RichTextStyle{
				TextStyle: fyne.TextStyle{Bold: true},
				SizeName:  theme.SizeNameHeadingText,
			},
		},
	)

	messageLabel := widget.NewLabel(*message)
	messageLabel.Selectable = true

	// Group text with tight 2px vertical spacing
	textGroup := container.New(&ui.TightVBoxLayout{Spacing: 2},
		titleLabel,
		messageLabel,
	)

	// Close Button with Circular Hover
	closeBtn := ui.NewCircularCloseButton(func() {
		myApp.Quit()
	})

	// Layout
	leftSide := container.NewHBox(
		iconContainer,
		utils.Spacer(12, 0),
	)

	rightSide := container.NewHBox(
		utils.Spacer(20, 0),
		container.NewCenter(closeBtn),
	)

	toastLayout := container.NewBorder(
		nil, nil,
		leftSide,
		rightSide,
		container.NewCenter(textGroup),
	)

	paddedContainer := container.NewBorder(
		utils.Spacer(0, 8), utils.Spacer(0, 8),
		utils.Spacer(14, 0), utils.Spacer(14, 0),
		toastLayout,
	)

	myWindow.SetContent(paddedContainer)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
