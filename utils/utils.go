package utils

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/constants"
)

// RemoveElement takes a slice of strings, removes element from the given argument and returns original slice without this element
func RemoveElement(slice []string, element string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

func WrapText(text string) []string {
	words := strings.Fields(text)
	var lines []string
	currentLine := ""

	for _, word := range words {
		testLine := currentLine
		if testLine != "" {
			testLine += " "
		}
		testLine += word

		textSize := rl.MeasureText(testLine, constants.FontSize)
		if word == "/n" {
			lines = append(lines, currentLine)
			currentLine = ""
		} else if textSize > 1000 {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			currentLine = testLine
		}
	}

	lines = append(lines, currentLine)
	return lines
}
