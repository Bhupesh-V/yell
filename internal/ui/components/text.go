package components

import (
	"strings"
	"yell/internal/ui/themes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Cap on-screen size for markdown images so a large source image can't blow
// up the notification window; images are scaled down to fit while keeping
// their aspect ratio (Fyne's default ImageSegment renders at native pixel
// size, which is unbounded).
const (
	maxMarkdownImageWidth  float32 = 320
	maxMarkdownImageHeight float32 = 200
)

// leftAlignBox pins its child to its own MinSize at the top-left corner,
// ignoring any extra space the parent hands it. Used to keep markdown
// images at a fixed, capped size and left-anchored even when placed inside
// a layout (e.g. VBox) that would otherwise stretch them to fill the row.
type leftAlignBox struct{}

func (leftAlignBox) Layout(objects []fyne.CanvasObject, _ fyne.Size) {
	for _, o := range objects {
		o.Resize(o.MinSize())
		o.Move(fyne.NewPos(0, 0))
	}
}

func (leftAlignBox) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var w, h float32
	for _, o := range objects {
		m := o.MinSize()
		if m.Width > w {
			w = m.Width
		}
		if m.Height > h {
			h = m.Height
		}
	}
	return fyne.NewSize(w, h)
}

// fitMarkdownImageSize returns the largest size that fits within
// maxMarkdownImageWidth x maxMarkdownImageHeight while preserving aspect.
// Using this (instead of just capping to a fixed box and relying on
// canvas.Image's ImageFillContain to center within it) avoids letterboxing:
// ImageFillContain centers the picture inside whatever Size() it's given,
// so a box whose aspect ratio doesn't match the image leaves the visible
// picture inset from the box's edges - which then also throws off the
// left-alignment with the message text above it.
func fitMarkdownImageSize(aspect float32) fyne.Size {
	w, h := maxMarkdownImageWidth, maxMarkdownImageHeight
	if aspect <= 0 {
		return fyne.NewSize(w, h)
	}
	if aspect > w/h {
		h = w / aspect
	} else {
		w = h * aspect
	}
	return fyne.NewSize(w, h)
}

type SelectableRichText struct {
	widget.RichText
}

// NewMarkdownMessage parses content as markdown and returns a widget that
// renders the text (selectable/copyable, see SelectableRichText) followed
// by any images, each capped to maxMarkdownImageWidth x
// maxMarkdownImageHeight and left-anchored under the first letter of the
// text (Fyne's RichText renders inline images at native pixel size and
// mis-positions them when they follow text on the same markdown line, so
// images are pulled out of the text flow and laid out separately here).
func NewMarkdownMessage(content string) fyne.CanvasObject {
	base := widget.NewRichTextFromMarkdown(content)

	textSegments := make([]widget.RichTextSegment, 0, len(base.Segments))
	var images []*widget.ImageSegment
	for _, seg := range base.Segments {
		if is, ok := seg.(*widget.ImageSegment); ok {
			images = append(images, is)
			continue
		}
		if ts, ok := seg.(*widget.TextSegment); ok {
			ts.Style.ColorName = themes.ColorNameSubtleText
		}
		textSegments = append(textSegments, seg)
	}

	rt := &SelectableRichText{}
	rt.Segments = textSegments
	rt.Wrapping = fyne.TextWrapWord
	// fyne.ScrollDirection's zero value is ScrollBoth (it's the first iota),
	// so building via a bare struct literal instead of widget.NewRichText()
	// leaves this widget scrollable by default, which confines it to a tiny
	// viewport instead of wrapping/growing - must set explicitly.
	rt.Scroll = fyne.ScrollNone
	rt.ExtendBaseWidget(rt)

	if len(images) == 0 {
		return rt
	}

	objects := []fyne.CanvasObject{rt}
	for _, is := range images {
		img := canvas.NewImageFromURI(is.Source)
		img.FillMode = canvas.ImageFillContain
		// Aspect() forces the image to load (Fyne loads URIs synchronously
		// here) so we know its real proportions before sizing it.
		img.SetMinSize(fitMarkdownImageSize(img.Aspect()))
		objects = append(objects, container.New(leftAlignBox{}, img))
	}
	return container.NewVBox(objects...)
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
