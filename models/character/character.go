package character

import (
	"github.com/dzwiedz90/go-public-game-demo/models"
	"github.com/dzwiedz90/go-public-game-demo/models/skills"
)

type Character struct {
	Name                    string
	Race                    string
	Class                   string
	Character               string
	Level                   int
	ExperiencePoints        int
	ProficiencyBonus        int
	PassiveWisdomPerception int
	ArmorClass              int
	Initiative              int
	Speed                   float32
	CurrentHitPoint         int
	MaxHitPoints            int
	TemporaryHitPoints      int
	HitPointDie             int
	SavingThrows            skills.SavingThrows
	Inventory               []models.Item
	InventoryCapacity       int
	InventoryWeight         float32
	AttackAndMagic          skills.AttackAndMagic
	FeaturesAndTraits       skills.FeaturesAndTraits
	Proficiencies           skills.Proficiencies
	Spells                  skills.Spells
}
