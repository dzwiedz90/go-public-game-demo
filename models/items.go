package models

type Weapon struct {
	ID          string
	Name        string
	Damage      string
	DamageType  string
	IsMelee     int
	IsRanged    int
	IsTwoHanded int
	IsFinesse   int
}

type Armor struct {
	ID               string
	Name             string
	Type             string
	AC               int
	MaxDexMod        int
	RequiredStrength int
}
