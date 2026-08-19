package ui

import (
	"yell/internal/ui/components"
	"yell/internal/ui/themes"
	"yell/internal/ui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Build(title, message, icon, appTheme *string) {
	// Initialise the application on the entry thread
	myApp := app.NewWithID("yell.app")

	// Schedule all UI building and window actions on the main thread via fyne.Do
	go fyne.Do(func() {
		ytheme := themes.GetTheme(themes.ParseTheme(*appTheme))
		myApp.Settings().SetTheme(ytheme)

		var myWindow fyne.Window
		if drv, ok := myApp.Driver().(desktop.Driver); ok {
			myWindow = drv.CreateSplashWindow()
		} else {
			myWindow = myApp.NewWindow(*title)
		}

		// Emoji Icon
		emojiObject := utils.GetEmbeddedEmojiImage(*icon, 42)
		iconContainer := container.NewCenter(emojiObject)

		// Elements specify semantic ColorNames instead of explicit RGB colors
		titleLabel := widget.NewRichText(
			&widget.TextSegment{
				Text: *title,
				Style: widget.RichTextStyle{
					TextStyle: fyne.TextStyle{Bold: true},
					SizeName:  theme.SizeNameHeadingText,
					ColorName: theme.ColorNameForeground,
				},
			},
		)

		messageLabel := widget.NewRichText(
			&widget.TextSegment{
				Text: *message,
				Style: widget.RichTextStyle{
					ColorName: themes.ColorNameSubtleText,
				},
			},
		)

		// Group text with tight 2px vertical spacing
		textGroup := container.New(&utils.TightVBoxLayout{Spacing: 2},
			titleLabel,
			messageLabel,
		)

		// Close Button with Circular Hover
		closeBtn := components.NewCircularCloseButton(func() {
			myApp.Quit() // Standard Fyne callback: safe to call directly here
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
		myWindow.Show() // Use Show() instead of ShowAndRun() when inside fyne.Do
	})

	// Block and run the driver lifecycle loop on the main execution thread
	myApp.Run()
}
