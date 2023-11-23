package newgame

import (
	"fmt"
	"strconv"

	"github.com/dzwiedz90/go-public-game-demo/constants"
	"github.com/dzwiedz90/go-public-game-demo/engine/core"
	"github.com/dzwiedz90/go-public-game-demo/engine/core/states"
	"github.com/dzwiedz90/go-public-game-demo/models/class"
	"github.com/dzwiedz90/go-public-game-demo/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	skillsItems               []skillItem
	skillsSelected            []string
	skillsCounter             = 0
	availableSkillsSelections int
	stageDone                 = false

	NewGameSkillsState = states.GameState{
		StateName: "AbilitiesState",
		OnEnter: func() {
			mkAvailableSkillsSelections()
			skillsItems = mkSkillItems()
		},
		OnUpdate: func() {
			rl.ClearBackground(rl.NewColor(23, 20, 22, 1))

			rl.UpdateMusicStream(*musicStream)

			inputSkills()

			drawSkillsOutput()
		},
		OnExit: func() {

		},
	}
)

func inputSkills() {
	if rl.IsKeyPressed(rl.KeyDown) && !stageDone {
		if skillsCounter == len(skillsItems)-1 {
			skillsCounter = 0
		} else {
			skillsCounter += 1
		}
	} else if rl.IsKeyPressed(rl.KeyUp) && !stageDone {
		if skillsCounter == 0 {
			skillsCounter = len(skillsItems) - 1
		} else {
			skillsCounter -= 1
		}
	} else if rl.IsKeyPressed(rl.KeyEnter) {
		if stageDone {
			saveSkills()
			states.CurrentGameState = NewGameAncestryState
			NewGameAncestryState.OnEnter()
		} else {
			if skillsItems[skillsCounter].isSelected {
				skillsItems[skillsCounter].isSelected = false
				availableSkillsSelections += 1
				skillsSelected = utils.RemoveElement(skillsSelected, skillsItems[skillsCounter].text)
			} else if !skillsItems[skillsCounter].isSelected && availableSkillsSelections > 0 {
				skillsItems[skillsCounter].isSelected = true
				skillsSelected = append(skillsSelected, skillsItems[skillsCounter].text)
				availableSkillsSelections -= 1
				skillsSelected = append(skillsSelected, skillsItems[skillsCounter].text)

				if availableSkillsSelections == 0 {
					stageDone = true
				}
			}
		}
	} else if rl.IsKeyPressed(rl.KeyEscape) {
		skillsItems = mkSkillItems()
		skillsSelected = []string{}
		skillsCounter = 0
		mkAvailableSkillsSelections()
		stageDone = false
	}
}

func drawSkillsOutput() {
	rl.DrawText("Tworzenie postaci", 25, 25, constants.FontSize, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("Umiejetnosci, wybierz %v sposrod nastepujacych,", availableSkillsSelections), 500, 25, constants.FontSize, rl.DarkGray)
	rl.DrawText("w ktorych bedziesz miec bieglosc", 500, 75, constants.FontSize, rl.DarkGray)

	var y int32 = 200

	for i, skill := range skillsItems {
		if i <= 9 {
			if i == skillsCounter {
				rl.DrawText(skill.text, 25, y, constants.FontSize, rl.White)

				if skillsItems[i].isSelected {
					rl.DrawText("[X]", 500, y, constants.FontSize, rl.White)
				} else {
					rl.DrawText("[ ]", 500, y, constants.FontSize, rl.White)
				}
			} else {
				rl.DrawText(skill.text, 25, y, constants.FontSize, rl.DarkGray)

				if skillsItems[i].isSelected {
					rl.DrawText("[X]", 500, y, constants.FontSize, rl.DarkGray)
				} else {
					rl.DrawText("[ ]", 500, y, constants.FontSize, rl.DarkGray)
				}
			}

			y += 50
		} else {
			if i == 10 {
				y = 200
			}

			if i == skillsCounter {
				rl.DrawText(skill.text, 600, y, constants.FontSize, rl.White)

				if skillsItems[i].isSelected {
					rl.DrawText("[X]", 1075, y, constants.FontSize, rl.White)
				} else {
					rl.DrawText("[ ]", 1075, y, constants.FontSize, rl.White)
				}
			} else {
				rl.DrawText(skill.text, 600, y, constants.FontSize, rl.DarkGray)

				if skillsItems[i].isSelected {
					rl.DrawText("[X]", 1075, y, constants.FontSize, rl.DarkGray)
				} else {
					rl.DrawText("[ ]", 1075, y, constants.FontSize, rl.DarkGray)
				}
			}

			y += 50
		}
	}

	if stageDone {
		rl.DrawText("Umiejetnosci wybrane, nacisnij enter, aby kontynuowac lub ESC aby wybrac ponownie...", 25, 125, constants.FontSize, rl.White)
	}
}

type skillItem struct {
	text       string
	isSelected bool
}

func mkSkillItems() []skillItem {
	s := class.ClassSkills[core.Character.Class][1:]
	var skills []skillItem

	for _, sk := range s {
		skills = append(skills, skillItem{
			text:       sk,
			isSelected: false,
		})
	}

	return skills
}

func mkAvailableSkillsSelections() {
	c := core.Character.Class
	v := class.ClassSkills[c]
	val, _ := strconv.Atoi(v[0])

	availableSkillsSelections = val
}

func saveSkills() {
	for _, skill := range skillsItems {
		if skill.text == "Akrobatyka" {
			core.Character.Skills.AcrobaticsProficiency = true
		} else if skill.text == "Atletyka" {
			core.Character.Skills.AthleticsProficiency = true
		} else if skill.text == "Historia" {
			core.Character.Skills.HistoryProficiency = true
		} else if skill.text == "Intuicja" {
			core.Character.Skills.InsightProficiency = true
		} else if skill.text == "Medycyna" {
			core.Character.Skills.MedicineProficiency = true
		} else if skill.text == "Opieka nad zwierzetami" {
			core.Character.Skills.AnimalHandlingProficiency = true
		} else if skill.text == "Oszustwo" {
			core.Character.Skills.DeceptionProficiency = true
		} else if skill.text == "Percepcja" {
			core.Character.Skills.PerceptionProficiency = true
		} else if skill.text == "Perswazja" {
			core.Character.Skills.PersuasionProficiency = true
		} else if skill.text == "Przyroda" {
			core.Character.Skills.NatureProficiency = true
		} else if skill.text == "Religia" {
			core.Character.Skills.ReligionProficiency = true
		} else if skill.text == "Skradanie sie" {
			core.Character.Skills.StealthProficiency = true
		} else if skill.text == "Sztuka przetrwania" {
			core.Character.Skills.SurvivalProficiency = true
		} else if skill.text == "Sledztwo" {
			core.Character.Skills.InvestigationProficiency = true
		} else if skill.text == "Wiedza tajemna" {
			core.Character.Skills.ArcanaProficiency = true
		} else if skill.text == "Wystepy" {
			core.Character.Skills.PerformanceProficiency = true
		} else if skill.text == "Zastraszanie" {
			core.Character.Skills.IntimidationProficiency = true
		} else if skill.text == "Zwinne dlonie" {
			core.Character.Skills.SleightOfHandProficiency = true
		}
	}
}
