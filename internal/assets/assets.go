package assets

import "embed"

// Embed assets/ including emojis and font.ttf

//go:embed emojis/*
var Emojis embed.FS

//go:embed font/*
var Font embed.FS

// Load Inter font from assets/font.ttf
var FontBytes, _ = Font.ReadFile("font.ttf")
