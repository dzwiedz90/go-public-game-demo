package character

import (
	"github.com/dzwiedz90/go-public-game-demo/models/skills"
)

type Character struct {
	Name                    string
	Race                    string
	Sex                     string
	Class                   string
	Character               string
	Abilities               skills.Abilities
	Skills                  skills.Skills
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
	Gold                    int
	Inventory               []interface{}
	InventoryCapacity       int
	InventoryWeight         float32
	AttackAndMagic          skills.AttackAndMagic
	FeaturesAndTraits       skills.FeaturesAndTraits
	Proficiencies           skills.Proficiencies
	Spells                  skills.Spells
}
