package components

import (
	"strings"
	"yell/internal/ui/themes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type SelectableRichText struct {
	widget.RichText
}

func NewSelectableMarkdown(content string) *SelectableRichText {
	// Parse Markdown using Fyne's built-in parser
	base := widget.NewRichTextFromMarkdown(content)

	rt := &SelectableRichText{}
	rt.Segments = base.Segments
	rt.Wrapping = fyne.TextWrapWord
	// fyne.ScrollDirection's zero value is ScrollBoth (it's the first iota),
	// so building via a bare struct literal instead of widget.NewRichText()
	// leaves this widget scrollable by default, which confines it to a tiny
	// viewport instead of wrapping/growing - must set explicitly.
	rt.Scroll = fyne.ScrollNone
	rt.ExtendBaseWidget(rt)

	// Apply subtle text color to text segments while preserving monospace flags
	for _, seg := range rt.Segments {
		if ts, ok := seg.(*widget.TextSegment); ok {
			ts.Style.ColorName = themes.ColorNameSubtleText
			if ts.Style.TextStyle.Monospace {
				ts.Style.TextStyle.Monospace = true
			}
		}
	}

	return rt
}

// Enable right-click / secondary tap to copy all text content to clipboard
// TappedSecondary handles right-click / secondary tap to copy full message content
func (s *SelectableRichText) TappedSecondary(_ *fyne.PointEvent) {
	s.copyAllToClipboard()
}

// Tapped handles double-tap or primary tap on the label
func (s *SelectableRichText) Tapped(_ *fyne.PointEvent) {
	s.copyAllToClipboard()
}

func (s *SelectableRichText) copyAllToClipboard() {
	var fullText strings.Builder

	for _, seg := range s.Segments {
		// Extract regular text and code blocks
		if ts, ok := seg.(*widget.TextSegment); ok {
			fullText.WriteString(ts.Text)
		}

		// xtract hyperlink label text (Fixes missing hyperlinks from copy)
		if hs, ok := seg.(*widget.HyperlinkSegment); ok {
			fullText.WriteString(hs.Text)
		}
	}

	if len(fyne.CurrentApp().Driver().AllWindows()) > 0 {
		fyne.CurrentApp().Clipboard().SetContent(fullText.String())
	}
}
