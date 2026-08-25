package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"yell/internal/assets"
	"yell/internal/ui"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/spf13/cobra"
)

var (
	title   string
	message string
	icon    string
	ftheme  string
	sound   string
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

func playSound(sound string) error {
	f, err := assets.OpenAudio(sound)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Decode the WAV file
	// The Decode function returns a stream that yields PCM data
	decodedWav, err := wav.DecodeWithSampleRate(44100, f)
	if err != nil {
		panic(err)
	}

	// Create an Oto context
	// You need to specify the sample rate (must match your WAV file)
	ctx, ready, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	})
	if err != nil {
		panic(err)
	}
	<-ready // Wait for the context to be ready

	player := ctx.NewPlayer(decodedWav)
	player.Play()

	// Keep the program running while the sound plays
	for player.IsPlaying() {
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func main() {
	defaultMsg := "Shawty lemme holla at you"

	// Root command (default alert execution)
	var rootCmd = &cobra.Command{
		Use:   "yell",
		Short: "Display a customized alert popup",
		Run: func(cmd *cobra.Command, args []string) {
			if piped := readPipedInput(); piped != "" {
				if message == defaultMsg {
					message = piped
				}
			}

			if sound != "" {
				go playSound(sound)
			}

			ui.Build(&title, &message, &icon, &ftheme)
		},
	}

	// Parent subcommand: list
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List available configurations (themes, sounds)",
	}

	// Child subcommand: list themes
	var themesCmd = &cobra.Command{
		Use:   "themes",
		Short: "List all available UI themes",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Available themes:\n- light\n- dark\n- warm\n-solarized-light\n-solarized-dark")
		},
	}

	// Child subcommand: list sounds
	var soundsCmd = &cobra.Command{
		Use:   "sounds",
		Short: "List all available alert sounds",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Available sounds:\n\n")
			sounds, err := assets.GetSoundNames()
			if err != nil {
				panic(err)
			}
			for _, sound := range sounds {
				fmt.Printf("- %s\n", sound)
			}
		},
	}

	// Define root command flags
	rootCmd.Flags().StringVarP(&title, "title", "t", "Somebody yelling", "Title of the alert popup")
	rootCmd.Flags().StringVarP(&message, "message", "m", "Shawty lemme holla at you", "Message text to display")
	rootCmd.Flags().StringVarP(&icon, "icon", "i", "🗣️", "Icon or emoji on the left")
	rootCmd.Flags().StringVarP(&ftheme, "theme", "e", "dark", "Theme mode: light or dark")
	rootCmd.Flags().StringVarP(&sound, "sound", "s", "", "Background Sound")

	// Nest child subcommands under listCmd
	listCmd.AddCommand(themesCmd, soundsCmd)

	// Attach listCmd to rootCmd
	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
