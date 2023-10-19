package class

import "github.com/dzwiedz90/go-public-game-demo/models/skills"

var (
	ClassBarbarian = "Barbarzynca"
	ClassBard      = "Bard"
	ClassCleric    = "Kleryk"
	ClassDruid     = "Druid"
	ClassFighter   = "Wojownik"
	ClassMonk      = "Mnich"
	ClassPaladin   = "Paladyn"
	ClassRanger    = "Lowca"
	ClassRogue     = "Lotr"
	ClassSorcerer  = "Zaklinacz"
	ClassWarlock   = "Czarownik"
	ClassWizard    = "Mag"

	Class = []string{
		ClassBarbarian,
		ClassBard,
		ClassCleric,
		ClassDruid,
		ClassFighter,
		ClassMonk,
		ClassPaladin,
		ClassRanger,
		ClassRogue,
		ClassSorcerer,
		ClassWarlock,
		ClassWizard,
	}

	ClassBarbarianDescription = "Barbarzynca - psychol w futrach. Zazwyczaj z dalekiej polnocy. Nie umie sie zachowac w towarzystwie i zaczepia panienki."
	ClassBardDescription      = "Bard - wyje niemilosernie tak, ze uszy puchna. Czesto konczy zaklebnowany i zwiazany w rogu pomieszczenia lub pod drzewem, zalezy gdzie akurat odbywa sie impreza."
	ClassClericDescription    = "Kleryk - czlowiek boga, jednego z calego stada darmozjadow. Zyje za cudze pieniadze i ma dlugie raczki do robty."
	ClassDruidDescription     = "Druid - sra w lesie i rucha dziki. Smierdzi na kilometr, gorzej nawet niz barbarzynca."
	ClassFighterDescription   = "Wojownik - typowy Seba co wszystkich pyta czy m,aja problem i by sie tylko napierdalal."
	ClassMonkDescription      = "Mnich - jesli myslales ze Wojownik to typowy Seba to ten tutaj to ultra Seba. Lysy napiedala sie piesciami. Przynajmniej nie pije."
	ClassPaladinDescription   = "Paladyn - swiatabliwy maz co niby ma celibat a tak naprawde to fujara mu zaraz odpadnie od wenery"
	ClassRangerDescription    = "Lowca - tez sra w lesie ale rucha jelenie zamiast dzikow. Strzela z luku ale ma slaba celnosc i zawsze trafia w stope, niekoniecznie swoja."
	ClassRogueDescription     = "Lotr - jesli masz sakiewke z hajsem to juz jej nie masz. Ciesz sie chociaz ze w dupe nie zapial ukradkiem od tylca"
	ClassSorcererDescription  = "Zaklinacz - kurwa kto?"
	ClassWarlockDescription   = "Czarownik - ni to mag, ni to warrior. Ani jednego ani drugiego nie robi dobrze"
	ClassWizardDescription    = "Mag - zlodupiec co ci sfajczy dupsko kula ognia. Jednostrzalowiec dlatego trzyma sie z tylu."

	ClassDescription = []string{
		ClassBarbarianDescription,
		ClassBardDescription,
		ClassClericDescription,
		ClassDruidDescription,
		ClassFighterDescription,
		ClassMonkDescription,
		ClassPaladinDescription,
		ClassRangerDescription,
		ClassRogueDescription,
		ClassSorcererDescription,
		ClassWarlockDescription,
		ClassWizardDescription,
	}

	ClassBarbarianAbilities = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassBardAbilities      = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassClericAbilities    = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassDruidAbilities     = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassFighterAbilities   = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassMonkAbilities      = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassPaladinAbilities   = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassRangerAbilities    = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassRogueAbilities     = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassSorcererAbilities  = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassWarlockAbilities   = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"
	ClassWizardAbilities    = "Cechy, bieglości /n Skile jakies tam costam costam /n Inne bullshity i rzeczy"

	ClassAbilities = []string{
		ClassBarbarianAbilities,
		ClassBardAbilities,
		ClassClericAbilities,
		ClassDruidAbilities,
		ClassFighterAbilities,
		ClassMonkAbilities,
		ClassPaladinAbilities,
		ClassRangerAbilities,
		ClassRogueAbilities,
		ClassSorcererAbilities,
		ClassWarlockAbilities,
		ClassWizardAbilities,
	}

	ClassHitPointsDie = map[string]int{
		ClassBarbarian: 12,
		ClassBard:      8,
		ClassCleric:    8,
		ClassDruid:     8,
		ClassFighter:   10,
		ClassMonk:      8,
		ClassPaladin:   10,
		ClassRanger:    10,
		ClassRogue:     8,
		ClassSorcerer:  6,
		ClassWarlock:   8,
		ClassWizard:    6,
	}

	ClassSavingThrows = map[string]skills.SavingThrows{
		ClassBarbarian: {ThrowsWithProficiency: []skills.Ability{skills.Strength, skills.Constitution}},
		ClassBard:      {ThrowsWithProficiency: []skills.Ability{skills.Dexteriety, skills.Charisma}},
		ClassCleric:    {ThrowsWithProficiency: []skills.Ability{skills.Wisdom, skills.Charisma}},
		ClassDruid:     {ThrowsWithProficiency: []skills.Ability{skills.Ingelligence, skills.Wisdom}},
		ClassFighter:   {ThrowsWithProficiency: []skills.Ability{skills.Strength, skills.Constitution}},
		ClassMonk:      {ThrowsWithProficiency: []skills.Ability{skills.Strength, skills.Dexteriety}},
		ClassPaladin:   {ThrowsWithProficiency: []skills.Ability{skills.Wisdom, skills.Charisma}},
		ClassRanger:    {ThrowsWithProficiency: []skills.Ability{skills.Strength, skills.Dexteriety}},
		ClassRogue:     {ThrowsWithProficiency: []skills.Ability{skills.Dexteriety, skills.Ingelligence}},
		ClassSorcerer:  {ThrowsWithProficiency: []skills.Ability{skills.Constitution, skills.Charisma}},
		ClassWarlock:   {ThrowsWithProficiency: []skills.Ability{skills.Wisdom, skills.Charisma}},
		ClassWizard:    {ThrowsWithProficiency: []skills.Ability{skills.Ingelligence, skills.Wisdom}},
	}
)
