package newgame

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/dzwiedz90/go-public-game-demo/constants"
	"github.com/dzwiedz90/go-public-game-demo/engine/core"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
	"github.com/dzwiedz90/go-public-game-demo/models/alignment"
	"github.com/dzwiedz90/go-public-game-demo/models/class"
	"github.com/dzwiedz90/go-public-game-demo/models/race"
)

type letter struct {
	letter     string
	isSelected bool
}

var (
	// key = false

	nameStage      = true
	raceStage      = false
	classStage     = false
	characterStage = false
	stagesDone     = false

	musicStream *rl.Music

	selectedRow    = 0
	selectedColumn = 0

	name string

	currentRace string
	raceCounter = 0

	currentClass string
	classCounter = 0

	currrentCharacter string
	characterCounter  = 0

	NewGameState = states.GameState{
		StateName: "NewGameState",
		OnEnter: func() {
			core.Character.Level = 1
			core.Character.ExperiencePoints = 0
			core.Character.ProficiencyBonus = 2
			core.Character.Initiative = 0
			core.Character.TemporaryHitPoints = 0
		},
		OnUpdate: func() {
			rl.ClearBackground(rl.NewColor(23, 20, 22, 1))

			if nameStage {
				nameStageInput()
			} else if raceStage {
				raceStageInput()
			} else if classStage {
				classStageInput()
			} else if characterStage {
				characterStageInput()
			}

			rl.UpdateMusicStream(*musicStream)

			drawMenu()
		},
		OnExit: func() {

		},
	}
)

func drawMenu() {
	rl.DrawText("Tworzenie postaci", 25, 25, constants.FontSize, rl.DarkGray)

	if nameStage {
		rl.DrawText("Imie: "+name, 50, 100, constants.FontSize, rl.White)
		drawAlphabet()
	} else if raceStage {
		rl.DrawText("Imie: "+core.Character.Name, 50, 100, constants.FontSize, rl.DarkGray)
		rl.DrawText("Rasa: "+currentRace, 50, 200, constants.FontSize, rl.White)
	} else if classStage {
		rl.DrawText("Imie: "+core.Character.Name, 50, 100, constants.FontSize, rl.DarkGray)
		rl.DrawText("Rasa: "+currentRace, 50, 200, constants.FontSize, rl.DarkGray)
		rl.DrawText("Klasa: "+currentClass, 50, 300, constants.FontSize, rl.White)
	} else if characterStage {
		rl.DrawText("Imie: "+core.Character.Name, 50, 100, constants.FontSize, rl.DarkGray)
		rl.DrawText("Rasa: "+currentRace, 50, 200, constants.FontSize, rl.DarkGray)
		rl.DrawText("Klasa: "+currentClass, 50, 300, constants.FontSize, rl.DarkGray)
		rl.DrawText("Charakter: ", 50, 400, constants.FontSize, rl.White)
		rl.DrawText(currrentCharacter, 50, 450, constants.FontSize, rl.White)
	} else if stagesDone {
		rl.DrawText("Imie: "+core.Character.Name, 50, 100, constants.FontSize, rl.DarkGray)
		rl.DrawText("Rasa: "+core.Character.Race, 50, 200, constants.FontSize, rl.DarkGray)
		rl.DrawText("Klasa: "+core.Character.Class, 50, 300, constants.FontSize, rl.DarkGray)
		rl.DrawText("Charakter: "+core.Character.Character, 50, 400, constants.FontSize, rl.DarkGray)

		rl.DrawText(core.Character.Name, 50, 500, constants.FontSize, rl.White)
		rl.DrawText(core.Character.Race, 50, 525, constants.FontSize, rl.White)
		rl.DrawText(core.Character.Class, 50, 550, constants.FontSize, rl.White)
		rl.DrawText(core.Character.Character, 50, 575, constants.FontSize, rl.White)

		rl.DrawText("Wciśnij enter, aby kontynuowac", 50, 600, constants.FontSize, rl.White)
		if rl.IsKeyPressed(rl.KeyEnter) {
			rl.DrawText("Dziala", 50, 650, constants.FontSize, rl.White)
			// Next state
			states.CurrentGameState = NewGameAbilitiesState
			NewGameAbilitiesState.OnEnter()
		}
	}
}

func nameStageInput() {
	// TODO dodac ze jak jest na maksa z prawej lub lewej to idzie do nastepnej lini zerowa pozycja

	if rl.IsKeyPressed(rl.KeyDown) {
		if selectedRow == 4 {
			if selectedColumn == 3 {
				alphabet[selectedRow][selectedColumn].isSelected = false
				selectedColumn = 2
			}
			if selectedColumn == 4 {
				alphabet[selectedRow][selectedColumn].isSelected = false
				selectedColumn = 3
			}
		}
		if selectedRow < 5 {
			alphabet[selectedRow][selectedColumn].isSelected = false
			selectedRow += 1
			alphabet[selectedRow][selectedColumn].isSelected = true
		}
	} else if rl.IsKeyPressed(rl.KeyUp) {
		if selectedRow > 0 {
			if selectedRow == 5 && selectedColumn == 3 {
				alphabet[selectedRow][selectedColumn].isSelected = false
				selectedRow = 4
				selectedColumn = 4
				alphabet[selectedRow][selectedColumn].isSelected = true
			} else {
				alphabet[selectedRow][selectedColumn].isSelected = false
				selectedRow -= 1
				alphabet[selectedRow][selectedColumn].isSelected = true
			}
		}
	} else if rl.IsKeyPressed(rl.KeyRight) {
		if selectedRow == 5 && selectedColumn < 3 {
			alphabet[selectedRow][selectedColumn].isSelected = false
			selectedColumn += 1
			alphabet[selectedRow][selectedColumn].isSelected = true
		}
		if selectedRow < 5 && selectedColumn < 4 {
			alphabet[selectedRow][selectedColumn].isSelected = false
			selectedColumn += 1
			alphabet[selectedRow][selectedColumn].isSelected = true
		}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		if selectedColumn > 0 {
			alphabet[selectedRow][selectedColumn].isSelected = false
			selectedColumn -= 1
			alphabet[selectedRow][selectedColumn].isSelected = true
		}
	} else if key := rl.IsKeyPressed(rl.KeyEnter); key {
		l := alphabet[selectedRow][selectedColumn].letter
		if l == "DEL" {
			name = name[:len(name)-1]
		} else if l == "SPACE" {
			name += " "
		} else if l == "OK" {
			nameStage = false
			core.Character.Name = name
			raceStage = true
		} else {
			if len(name) == 0 {
				name += strings.ToUpper(l)
			} else if string(name[len(name)-1]) == " " {
				name += strings.ToUpper(l)
			} else {
				name += l
			}
		}
	}
}

func raceStageInput() {
	rl.DrawText("Press enter to choose", 500, 100, constants.FontSize, rl.DarkGray)

	lines := wrapText(race.RacesDescription[raceCounter])
	x := int32(500)
	y := int32(200)
	for _, line := range lines {
		rl.DrawText(line, x, y, constants.FontSize, rl.White)
		y += constants.FontSize + 5
	}

	y += 50
	rl.DrawText("Cechy rasy:", x, y, constants.FontSize, rl.White)
	raceAbilitiesLines := wrapText(race.RacesAbilities[raceCounter])
	y += 50
	for _, line := range raceAbilitiesLines {
		rl.DrawText(line, x, y, constants.FontSize, rl.White)
		y += constants.FontSize + 5
	}

	currentRace = race.Race[raceCounter]

	if rl.IsKeyPressed(rl.KeyUp) {
		if raceCounter > 0 {
			raceCounter -= 1
		} else {
			raceCounter = len(race.Race) - 1
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if raceCounter < len(race.Race)-1 {
			raceCounter += 1
		} else {
			raceCounter = 0
		}
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		core.Character.Race = currentRace
		core.Character.Speed = race.RaceSpeed[currentRace]
		raceStage = false
		classStage = true
	}
}

func classStageInput() {
	rl.DrawText("Press enter to choose", 500, 100, constants.FontSize, rl.DarkGray)

	lines := wrapText(class.ClassDescription[classCounter])
	x := int32(500)
	y := int32(200)
	for _, line := range lines {
		rl.DrawText(line, x, y, constants.FontSize, rl.White)
		y += constants.FontSize + 5
	}

	y += 50
	rl.DrawText("Cechy klasy:", x, y, constants.FontSize, rl.White)
	classAbilitiesLines := wrapText(class.ClassAbilities[classCounter])
	y += 50
	for _, line := range classAbilitiesLines {
		rl.DrawText(line, x, y, constants.FontSize, rl.White)
		y += constants.FontSize + 5
	}

	currentClass = class.Class[classCounter]

	if rl.IsKeyPressed(rl.KeyUp) {
		if classCounter > 0 {
			classCounter -= 1
		} else {
			classCounter = len(class.Class) - 1
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if classCounter < len(class.Class)-1 {
			classCounter += 1
		} else {
			classCounter = 0
		}
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		core.Character.Class = currentClass
		core.Character.HitPointDie = class.ClassHitPointsDie[currentClass]
		core.Character.SavingThrows = class.ClassSavingThrows[currentClass]
		classStage = false
		characterStage = true
	}
}

func characterStageInput() {
	rl.DrawText("Press enter to choose", 500, 100, constants.FontSize, rl.DarkGray)

	lines := wrapText(alignment.AlignementDescription[characterCounter])
	x := int32(500)
	y := int32(200)
	for _, line := range lines {
		rl.DrawText(line, x, y, constants.FontSize, rl.White)
		y += constants.FontSize + 5
	}

	currrentCharacter = alignment.Alignement[characterCounter]

	if rl.IsKeyPressed(rl.KeyUp) {
		if characterCounter > 0 {
			characterCounter -= 1
		} else {
			characterCounter = len(alignment.Alignement) - 1
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		if characterCounter < len(alignment.Alignement)-1 {
			characterCounter += 1
		} else {
			characterCounter = 0
		}
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		core.Character.Character = currrentCharacter
		characterStage = false
		stagesDone = true
	}
}

func drawAlphabet() {
	var x int32
	var y int32
	x = 700
	y = 100
	for _, row := range alphabet {
		for _, letter := range row {
			if letter.letter == "OK" {
				x += 75
			}
			if letter.isSelected {
				rl.DrawText(letter.letter, x, y, constants.FontSize, rl.White)
			} else {
				rl.DrawText(letter.letter, x, y, constants.FontSize, rl.DarkGray)
			}
			x += 75
		}
		x = 700
		y += 75
	}
}

func SetNewGameMusicStream(musicStreamIn *rl.Music) {
	musicStream = musicStreamIn
}

func wrapText(text string) []string {
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
