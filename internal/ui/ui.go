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
	myApp := app.NewWithID("yell.app")

	// SetOnStarted runs inside the active Fyne event loop once myApp.Run() starts
	myApp.Lifecycle().SetOnStarted(func() {
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

		textGroup := container.New(&utils.TightVBoxLayout{Spacing: 2},
			titleLabel,
			messageLabel,
		)

		closeBtn := components.NewCircularCloseButton(func() {
			myApp.Quit()
		})

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
		myWindow.Show()
	})

	// Starts the driver event loop on the main thread
	myApp.Run()
}
