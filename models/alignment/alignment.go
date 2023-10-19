package alignment

var (
	LawfulGood     = "Prawodzadny dobry"
	NeutralGood    = "Neutralny dobry"
	ChaoticGood    = "Chaotyczny dobry"
	LawfulNeutral  = "Prawodzadny neutralny"
	TrueNeutral    = "Neutralny"
	ChaoticNeutral = "Chaotyczny zly"
	LawfulEvil     = "Prawodzadny zly"
	NeutralEvil    = "Neutralny zly"
	ChaoticEvil    = "Chaotyczny zly"

	Alignement = []string{
		LawfulGood,
		NeutralGood,
		ChaoticGood,
		LawfulNeutral,
		TrueNeutral,
		ChaoticNeutral,
		LawfulEvil,
		NeutralEvil,
		ChaoticEvil,
	}

	LawfulGoodDescirption     = "Praworzadny dobry charakter zazwyczaj dziala ze wspolczuciem, ale zawsze z honorem i poczuciem obowiazku. Jednakze zgodne z prawem dobre postacie czesto beda zalowac podjecia jakichkolwiek dzialan, ktorych obawiaja sie, ze narusza ich kodeks, nawet jesli uznaja takie dzialanie za dobre. Do takich postaci zaliczaja sie zlote smoki, prawi rycerze, paladyni i wiekszosc krasnoludow."
	NeutralGoodDescirption    = "Neutralny dobry charakter zazwyczaj dziala altruistycznie, bez wzgledu na zgodne z prawem zasady, takie jak zasady czy tradycja, lub przeciwko nim. Neutralny, dobry charakter nie ma problemow ze wspolpraca z prawowitymi urzednikami, ale nie czuje sie wobec nich zobowiazany. W przypadku, gdy postepowanie wlasciwe wymaga naginania lub lamania zasad, nie doswiadczaja tego samego wewnetrznego konfliktu, jaki odczuwalby praworzadny, dobry charakter. Przykladami tego ustawienia sa: wielu niebian, niektore olbrzymy chmurowe i wiekszosc gnomow."
	ChaoticGoodDescirption    = "Chaotyczny dobry charakter robi wszystko, co konieczne, aby wprowadzic zmiany na lepsze, gardzi organizacjami biurokratycznymi, ktore staja na drodze do poprawy spolecznej i przywiazuje duza wage do wolnosci osobistej, nie tylko dla siebie, ale takze dla innych. Chaotyczni dobrzy bohaterowie zwykle zamierzaja postepowac wlasciwie, ale ich metody sa na ogol zdezorganizowane i czesto niezsynchronizowane z reszta spoleczenstwa. Przykladami tego charakteru sa miedziane smoki, wiele elfow i jednorozce."
	LawfulNeutralDescirption  = "Praworzadna neutralna postac zazwyczaj mocno wierzy w zgodne z prawem koncepcje, takie jak honor, porzadek, zasady i tradycja, ale czesto kieruje sie osobistym kodeksem oprocz lub nawet zamiast tego ustanowionego przez zyczliwa wladze. Przykladami takiego charakteru sa zolnierz, ktory zawsze wykonuje rozkazy, sedzia lub egzekutor, ktory bezlitosnie przestrzega litery prawa, zdyscyplinowany mnich i niektorzy czarodzieje."
	TrueNeutralDescirption    = "Postac neutralna (nazywana takze „prawdziwie neutralna”) jest neutralna w obu osiach i zwykle nie jest silnie zwiazana z zadnym charakterem lub aktywnie poszukuje ich rownowagi. Druidzi czesto podazaja za tym dazeniem do zachowania rownowagi i zgodnie z zasadami Advanced Dungeons & Dragons musieli miec ten charakter. W przykladzie podanym w Drugiej edycji Podrecznika Gracza, typowy druid moze walczyc z banda grasujacych gnolli, tylko po to, aby zmienic strone, aby ocalic klan gnolli przed calkowita eksterminacja. Przykladami takiego charakteru sa jaszczuroludzie, wiekszosc druidow i wielu ludzi."
	ChaoticNeutralDescirption = "Chaotyczna neutralna postac to indywidualista, ktory podaza za wlasnym sercem i generalnie uchyla sie od zasad i tradycji. Chociaz chaotycznie neutralne postacie promuja idealy wolnosci, na pierwszym miejscu jest ich wlasna wolnosc; dobro i zlo sa na drugim miejscu w stosunku do ich potrzeby wolnosci. Przykladami tego charakteru sa wielu barbarzyncow i lotrzykow oraz niektorzy bardowie."
	LawfulEvilDescirption     = "Praworzadny zly charakter postrzega dobrze uporzadkowany system jako niezbedny do zaspokojenia swoich osobistych pragnien i potrzeb, wykorzystujac te systemy do zwiekszania swojej wladzy i wplywow. Przykladami tego dostosowania sa tyrani, diably, skorumpowani urzednicy, bezkrytyczni najemnicy, ktorzy maja scisly kodeks postepowania, niebieskie smoki i hobgobliny."
	NeutralEvilDescirption    = "Neutralna zla postac jest zazwyczaj samolubna i nie ma zadnych skrupulow, jesli chodzi o zwrocenie sie przeciwko chwilowym sojusznikom i zazwyczaj zawiera sojusze przede wszystkim po to, by realizowac wlasne cele. Neutralna zla postac nie ma skrupulow, by skrzywdzic innych, aby osiagnac to, czego chca, ale nie zrobi tez nic, co w jej mocy, aby wywolac rzez lub chaos, gdy nie widzi dla siebie bezposrednich korzysci. Inna sluszna interpretacja neutralnego zla przedstawia zlo jako ideal, czyniac zlo dla samego zla i probujac szerzyc jego wplywy. Przykladami pierwszego typu sa zabojca, ktory nie zwraca uwagi na formalne prawa, ale nie zabija niepotrzebnie, pomocnik spiskujacy za plecami przelozonego lub najemnik, ktory chetnie zmienia strone, jesli otrzyma lepsza oferte. Przykladem drugiego typu moze byc zamaskowany zabojca, ktory uderza jedynie po to, aby wywolac strach i nieufnosc w spoleczenstwie. Przykladami tego charakteru sa wiele drowow, niektorzy chmurowi olbrzymy i yugolothy."
	ChaoticEvilDescirption    = "Chaotyczny zly charakter zwykle nie szanuje zasad, zycia innych ludzi ani niczego poza wlasnymi pragnieniami, ktore sa zazwyczaj samolubne i okrutne. Przywiazuja duza wage do wolnosci osobistej, ale nie maja wiekszego szacunku dla zycia i wolnosci innych ludzi. Chaotyczne zle postacie nie radza sobie dobrze w grupach, poniewaz nie lubia, gdy wydaje im sie rozkazy i zwykle nie zachowuja sie samodzielnie, chyba ze nie ma alternatywy. Przykladami tego charakteru sa wyzsze formy nieumarlych (takie jak licze), brutalni zabojcy, ktorzy uderzaja dla przyjemnosci, a nie dla zysku, demony, czerwone smoki i orki."

	AlignementDescription = []string{
		LawfulGoodDescirption,
		NeutralGoodDescirption,
		ChaoticGoodDescirption,
		LawfulNeutralDescirption,
		TrueNeutralDescirption,
		ChaoticNeutralDescirption,
		LawfulEvilDescirption,
		NeutralEvilDescirption,
		ChaoticEvilDescirption,
	}
)
