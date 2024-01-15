package items

import (
	"database/sql"
	"fmt"

	"github.com/dzwiedz90/go-public-game-demo/db/items/sqldriver"
	"github.com/dzwiedz90/go-public-game-demo/models"
)

func GetItemById(id string) (interface{}, error) {
	db, err := sqldriver.Connect()
	if err != nil {
		fmt.Printf("cannot connect to the db: %v", err)
		return nil, err
	}

	itemType := id[:2]
	if itemType == "ar" {
		armor, err := getArmor(db, id)
		if err != nil {
			fmt.Printf("cannot get armor from the db: %v", err)
			return nil, err
		}

		return armor, nil
	} else if itemType == "wp" {
		weapon, err := getWeapon(db, id)
		if err != nil {
			fmt.Printf("cannot get armor from the db: %v", err)
			return nil, err
		}
		return weapon, nil
	}

	return nil, nil
}

func getWeapon(db *sql.DB, id string) (interface{}, error) {
	row := db.QueryRow("SELECT id, name, damage, damage_type, is_melee, is_ranged, is_two_handed, is_finesse FROM weapons WHERE id = ?", id)

	weapon := models.Weapon{}
	err := row.Scan(&weapon.ID, &weapon.Name, &weapon.Damage, &weapon.DamageType, &weapon.IsMelee, &weapon.IsRanged, &weapon.IsTwoHanded, &weapon.IsFinesse)
	if err != nil {
		fmt.Printf("failed to scan rows into Go Weapon structure: %v", err)
		return nil, err
	}

	return weapon, nil
}

func getArmor(db *sql.DB, id string) (interface{}, error) {
	row := db.QueryRow("SELECT id, name, type, ac, max_dex_mod, required_strength FROM armor WHERE id = ?", id)

	armor := models.Armor{}
	err := row.Scan(&armor.ID, &armor.Name, &armor.Type, &armor.AC, &armor.MaxDexMod, &armor.RequiredStrength)
	if err != nil {
		fmt.Printf("failed to scan rows into Go Armor structure: %v", err)
		return nil, err
	}

	return armor, nil
}
