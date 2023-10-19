package skills

type Ability string

var (
	Strength     Ability = "Sila"
	Dexteriety   Ability = "Zrecznosc"
	Constitution Ability = "Kondycja"
	Ingelligence Ability = "Inteligencja"
	Wisdom       Ability = "Madrosc"
	Charisma     Ability = "Charyzma"
)

// Abilities Cechy
type Abilities struct {
	Strength     int
	Dexteriety   int
	Constitution int
	Ingelligence int
	Wisdom       int
	Charisma     int
}

// Skills Umiejętności
type Skills struct {
	// Akrobatyka
	Actobatics int
	// Atletyka
	Athletics int
	// Historia
	History int
	// Intuicja
	Insight int
	// Medycyna
	Medicine int
	// Opieka nad zwierzętami
	AnimalHandling int
	// Oszustwo
	Deception int
	// Percepcja
	Perception int
	// Perswazja
	Persuasion int
	// Przyroda
	Nature int
	// Religia
	Religion int
	// Skradanie się
	Stealth int
	// Sztuka przetrwania
	Survival int
	// Śledztwo
	Investigation int
	// Wiedza tajemna
	Arcana int
	// Wystepy
	Performance int
	// Zastraszanie
	Intimidation int
	// Zwinne dłonie
	SleightOfHand int
}

// SavingThrows Rzuty obronne
type SavingThrows struct {
	ThrowsWithProficiency []Ability
	Strength              int
	Dexteriety            int
	Constitution          int
	Intelligence          int
	Wisdom                int
	Charisma              int
}

// AttackAndMagic Ataki i magia
type AttackAndMagic struct {
	EquippedItemId int
	// Premia do ataku
	HitDC      int
	Damage     int
	DamageType string
}

// FeaturesAndTraits Korzyści i zdolności
type FeaturesAndTraits struct {
	Features []Feature
	Traits   []Trait
}

// Proficiencies Biegłości i języki
type Proficiencies struct {
	Languages   []string
	Proficiency []Proficiency
}

type Spells struct {
	BasicAbility                            string
	SpellSaveDC                             int
	SpellAttackModifier                     int
	Cantrip                                 []Spell
	CantripNumberOfSlots                    int
	FirststLevel                            []Spell
	FirstLevelNumberOfSlots                 int
	FirstLevelNumberOfAvaialbleSlots        int
	SecondLevel                             []Spell
	SecondLevelLevelNumberOfSlots           int
	SecondLevelLevelNumberOfAvaialbleSlots  int
	ThirdLevel                              []Spell
	ThirdLevelLevelNumberOfSlots            int
	ThirdLevelLevelNumberOfAvaialbleSlots   int
	FourthLevel                             []Spell
	FourthLevelLevelNumberOfSlots           int
	FourthLevelLevelNumberOfAvaialbleSlots  int
	FifthLevel                              []Spell
	FifthLevelLevelNumberOfSlots            int
	FifthLevelLevelNumberOfAvaialbleSlots   int
	SixthLevel                              []Spell
	SixthLevelLevelNumberOfSlots            int
	SixthLevelLevelNumberOfAvaialbleSlots   int
	SeventhLevel                            []Spell
	SeventhLevelLevelNumberOfSlots          int
	SeventhLevelLevelNumberOfAvaialbleSlots int
	EightLevel                              []Spell
	EightLevelLevelNumberOfSlots            int
	EightLevelLevelNumberOfAvaialbleSlots   int
	NinthLevel                              []Spell
	NinthLevelLevelNumberOfSlots            int
	NinthLevelLevelNumberOfAvaialbleSlots   int
}
