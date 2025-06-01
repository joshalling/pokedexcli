package commands

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"regexp"
	"strings"

	"github.com/qeesung/image2ascii/convert"
)

func Print(c *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokémon name or ID")
	}
	id := args[0]

	pokemon, exists := c.Data.Pokedex[id]
	if !exists {
		return fmt.Errorf("pokemon (%s) not found in Pokedex", id)
	}

	if pokemon.Sprites.FrontDefault == "" {
		return fmt.Errorf("no sprite available for %s", pokemon.Name)
	}

	file, err := c.Client.GetImage(pokemon.Sprites.FrontDefault)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	defer file.Close()

	// Diagnostic: check file size
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("could not stat temp image file: %w", err)
	}
	if info.Size() == 0 {
		return fmt.Errorf("downloaded image file is empty")
	}

	_, _, imgErr := image.Decode(file)
	if imgErr != nil {
		return fmt.Errorf("could not decode image: %w", imgErr)
	}
	file.Seek(0, 0)

	options := convert.DefaultOptions
	options.FixedWidth = 0
	options.FixedHeight = 80
	converter := convert.NewImageConverter()
	ascii := converter.ImageFile2ASCIIString(file.Name(), &options)
	ansi := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

	lines := strings.Split(ascii, "\n")
	start := 0
	end := len(lines)
	for start < end {
		clean := ansi.ReplaceAllString(lines[start], "")
		if strings.TrimSpace(clean) != "" {
			break
		}
		start++
	}
	for end > start {
		clean := ansi.ReplaceAllString(lines[end-1], "")
		if strings.TrimSpace(clean) != "" {
			break
		}
		end--
	}
	trimmed := lines[start:end]
	fmt.Printf("\n%s\n", strings.Join(trimmed, "\n"))

	return nil
}
