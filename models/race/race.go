package race

var (
	RaceDragonBorn        = "Drakon"
	RaceDwarf             = "Krasnolud"
	RaceMountainDwarf     = "Krasnolud gorski"
	RaceHillDwarf         = "Krasnolud wzgorzowy"
	RaceElf               = "Elf"
	RaceForestElf         = "Elf lesny"
	RaceHighElf           = "Elf wysoki"
	RaceDrowElf           = "Elf mroczny (Drow)"
	RaceGnome             = "Gnom"
	RaceForestGnome       = "Gnom leśny"
	RaceRockGnome         = "Gnom skalny"
	RaceHalfElf           = "Pólelf"
	RaceHalfling          = "Niziolek"
	RaceToughHalfling     = "Niziolek hardy"
	RaceLightfootHalfling = "Niziolek lekkostopy"
	RaceHalfOrc           = "Pólork"
	RaceHuman             = "Czlowiek"
	RaceTiefling          = "Diabelstwo"

	Race = []string{
		RaceDragonBorn,
		RaceDwarf,
		RaceElf,
		RaceGnome,
		RaceForestGnome,
		RaceRockGnome,
		RaceHalfElf,
		RaceHalfling,
		RaceToughHalfling,
		RaceLightfootHalfling,
		RaceHalfOrc,
		RaceHuman,
		RaceTiefling,
	}

	RaceDragonBornDescription        = "Drakon - ma luski i smierdzi siarka. Zre duzo miecha i potem ma refluks. Jego cechami charakterystycznymi sa dlugi ogon, luski i dlugie duze oczy."
	RaceDwarfDescription             = "Krasnolud - ma dluga brode i jeszcze wieksza knage. Potrafi wychlac tyle piwa ile sam wazy i dalej stac na nogach."
	RaceMountainDwarfDescription     = "Krasnolud - ma dluga brode i jeszcze wieksza knage. Potrafi wychlac tyle piwa ile sam wazy i dalej stac na nogach."
	RaceHillDwarfDescription         = "Krasnolud - ma dluga brode i jeszcze wieksza knage. Potrafi wychlac tyle piwa ile sam wazy i dalej stac na nogach."
	RaceElfDescription               = "Elf - lubi chasac po drzewach, biegac wsrod kwiatow i brac w dupe. Jego spiew ma wlasciwosci ogluszajace a glebokie spojrzenie w jego oczy spedali nawet najbardziej zatwardzialego hetero."
	RaceForestElfDescription         = "Elf lesny - jeszcze bardziej lubi chasac po drzewach, biegac wsrod kwiatow i brac w dupe. Jego spiew ma wlasciwosci ogluszajace a glebokie spojrzenie w jego oczy spedali nawet najbardziej zatwardzialego hetero."
	RaceHighElfDescription           = "Elf wysoki - moze wcale nie byc tak wysoki. Siedzi tylko na koscistej dupie i czyta ksiazki. Czesto para sie magia. Jego ulubione zaklecie to Penisum erectum.Jego glebokie spojrzenie w jego oczy spedali nawet najbardziej zatwardzialego hetero."
	RaceDrowElfDescription           = "Elf mroczny (Drow) - nie lubi chasac po drzewach, biegac wsrod kwaitow ale lubi brac w dupe. Lubi kaszanke i torcik na klate"
	RaceGnomeDescription             = "Gnom - jest malym kurduplem, ktory lubuje sie w jedzeniu paczkow z makiem i szczaniu do beczek z piwem w ludzkich gospodach."
	RaceForestGnomeDescription       = "Gnom - jest malym kurduplem, ktory lubuje sie w jedzeniu paczkow z makiem i szczaniu do beczek z piwem w ludzkich gospodach."
	RaceRockGnomeDescription         = "Gnom - jest malym kurduplem, ktory lubuje sie w jedzeniu paczkow z makiem i szczaniu do beczek z piwem w ludzkich gospodach."
	RaceHalfElfDescription           = "Pólelf - jak elf tylko lubi jeszcze z dupy do ust. Gra tez na lutni."
	RaceHalflingDescription          = "Niziolek - jeszcze mniejszy kurdupel. Mieszka w norach, zre, chleje i pasie sadlo. Lepiej nie powiezac mu zadnej bizuterii bo polezie do wulkanu i wrzuci do lawy."
	RaceToughHalflingDescription     = "Niziolek - jeszcze mniejszy kurdupel. Mieszka w norach, zre, chleje i pasie sadlo. Lepiej nie powiezac mu zadnej bizuterii bo polezie do wulkanu i wrzuci do lawy."
	RaceLightfootHalflingDescription = "Niziolek - jeszcze mniejszy kurdupel. Mieszka w norach, zre, chleje i pasie sadlo. Lepiej nie powiezac mu zadnej bizuterii bo polezie do wulkanu i wrzuci do lawy."
	RaceHalfOrcDescription           = "Pólork - zielony i smierdzacy. Gwalci wszystko co ma blond wlosy i nie spierdala na drzewo."
	RaceHumanDescription             = "Czlowiek - czlek jak to czlek, rozjebie wszystko czego sie dotknie i wszedzie spusci szambo."
	RaceTieflingDescription          = "Diabelstwo - fur Deutschland jak rzekl Donaldu Tusku. Jak sam Tusk pochodzi z czelusci piekeilnych i tylko mysli o tym jak zaciagnac male dzieci do piwnicy."

	RacesDescription = []string{
		RaceDragonBornDescription,
		RaceDwarfDescription,
		RaceMountainDwarfDescription,
		RaceHillDwarfDescription,
		RaceElfDescription,
		RaceForestElfDescription,
		RaceHighElfDescription,
		RaceDrowElfDescription,
		RaceGnomeDescription,
		RaceForestGnomeDescription,
		RaceRockGnomeDescription,
		RaceHalfElfDescription,
		RaceHalflingDescription,
		RaceToughHalflingDescription,
		RaceLightfootHalflingDescription,
		RaceHalfOrcDescription,
		RaceHumanDescription,
		RaceTieflingDescription,
	}

	RaceDragonBornRaceAbilities        = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceDwarfRaceAbilities             = "Zwiekszenie wartosci cechy: Wartosc twojej Kondycji wzrasta o 2. /n Szybkosc: Twoja podstawowa szybkosc to 9 metrow. /n Jezyki: potrafisz mowic we wspolnym oraz dodatkowym wybranym jezyku."
	RaceMountainDwarfRaceAbilities     = "Zwiekszenie wartosci cechy: Wartosc twojej Kondycji wzrasta o 2. /n Szybkosc: Twoja podstawowa szybkosc to 9 metrow. /n Jezyki: potrafisz mowic we wspolnym oraz dodatkowym wybranym jezyku."
	RaceHillDwarfRaceAbilities         = "Zwiekszenie wartosci cechy: Wartosc twojej Kondycji wzrasta o 2. /n Szybkosc: Twoja podstawowa szybkosc to 9 metrow. /n Jezyki: potrafisz mowic we wspolnym oraz dodatkowym wybranym jezyku."
	RaceElfRaceAbilities               = "Zwiekszenie wartosci cechy: Wartosc twojej Zrecznosci wzrasta o 2. /n Szybkosc: Twoja podstawowa szybkosc to 9 metrow. /n Jezyki: potrafisz mowic we wspolnym oraz dodatkowym wybranym jezyku."
	RaceForestElfRaceAbilities         = "Zwiekszenie wartosci cechy: Wartosc twojej Zrecznosci wzrasta o 2. /n Szybkosc: Twoja podstawowa szybkosc to 9 metrow. /n Jezyki: potrafisz mowic we wspolnym oraz dodatkowym wybranym jezyku."
	RaceHighElfRaceAbilities           = "Zwiekszenie wartosci cechy: Wartosc twojej Zrecznosci wzrasta o 2. /n Szybkosc: Twoja podstawowa szybkosc to 9 metrow. /n Jezyki: potrafisz mowic we wspolnym oraz dodatkowym wybranym jezyku."
	RaceDrowElfRaceAbilities           = "Zwiekszenie wartosci cechy: Wartosc twojej Zrecznosci wzrasta o 2. /n Szybkosc: Twoja podstawowa szybkosc to 9 metrow. /n Jezyki: potrafisz mowic we wspolnym oraz dodatkowym wybranym jezyku."
	RaceGnomeRaceAbilities             = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceForestGnomeRaceAbilities       = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceRockGnomeRaceAbilities         = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceHalfElfRaceAbilities           = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceHalflingRaceAbilities          = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceToughHalflingRaceAbilities     = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceLightfootHalflingRaceAbilities = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceHalfOrcRaceAbilities           = "Jakies Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"
	RaceHumanRaceAbilities             = "Zwiekszenie wartosci cechy: Wartosc wszystkich wtoich cech wzrasta o 1. /n Wiek: Ludzie osiagaja doroslosc okolo 18 lat i rzadko zyja dluzej niz 100 lat. /n Charakter: Ludzie moga miec dowolny charakter. /n Szybkosc: Twoja podstawowa szybkosc to 9 metrow. /n Jezyki: potrafisz mowic we wspolnym oraz dodatkowym wybranym jezyku."
	RaceTieflingRaceAbilities          = "Zwiekszenie wartosci: Wartosc Kondycji. /n Szybkosc: Twoja podstawowa"

	RacesAbilities = []string{
		RaceDragonBornRaceAbilities,
		RaceDwarfRaceAbilities,
		RaceMountainDwarfRaceAbilities,
		RaceHillDwarfRaceAbilities,
		RaceElfRaceAbilities,
		RaceForestElfRaceAbilities,
		RaceHighElfRaceAbilities,
		RaceDrowElfRaceAbilities,
		RaceGnomeRaceAbilities,
		RaceForestGnomeRaceAbilities,
		RaceRockGnomeRaceAbilities,
		RaceHalfElfRaceAbilities,
		RaceHalflingRaceAbilities,
		RaceToughHalflingRaceAbilities,
		RaceLightfootHalflingRaceAbilities,
		RaceHalfOrcRaceAbilities,
		RaceHumanRaceAbilities,
		RaceTieflingRaceAbilities,
	}

	RaceSpeed = map[string]float32{
		RaceDragonBorn:        9,
		RaceDwarf:             7.5,
		RaceMountainDwarf:     7.5,
		RaceHillDwarf:         7.5,
		RaceElf:               9,
		RaceForestElf:         10.5,
		RaceHighElf:           9,
		RaceDrowElf:           9,
		RaceGnome:             7.5,
		RaceForestGnome:       7.5,
		RaceRockGnome:         7.5,
		RaceHalfElf:           9,
		RaceHalfling:          7.5,
		RaceToughHalfling:     7.5,
		RaceLightfootHalfling: 7.5,
		RaceHalfOrc:           9,
		RaceHuman:             9,
		RaceTiefling:          9,
	}
)
