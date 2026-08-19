package assets

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"strings"
)

//go:embed emojis/*
var Emojis embed.FS

//go:embed font/*
var Font embed.FS

//go:embed sound/*
var Sounds embed.FS

// Load Inter font from assets/font.ttf
var FontBytes, _ = Font.ReadFile("font.ttf")

func GetEmojiFilePath(emoji string) string {
	var fullHex []string
	var cleanHex []string

	for _, r := range emoji {
		hexVal := fmt.Sprintf("%x", r)
		fullHex = append(fullHex, hexVal)
		if r != 0xfe0f {
			cleanHex = append(cleanHex, hexVal)
		}
	}

	cleanPath := path.Join("emojis", strings.Join(cleanHex, "-")+".png")
	if _, err := Emojis.ReadFile(cleanPath); err == nil {
		return cleanPath
	}

	fullPath := path.Join("emojis", strings.Join(fullHex, "-")+".png")
	if _, err := Emojis.ReadFile(fullPath); err == nil {
		return fullPath
	}

	return ""
}

// OpenAudio opens and returns the embedded audio file
func OpenAudio(sound string) (fs.File, error) {
	return Sounds.Open(path.Join("sound", sound+".wav"))
}
