package items

import (
	"database/sql"

	"github.com/dzwiedz90/go-public-game-demo/models"
)

func connectItemsDb() (*sql.DB, error) {
	dbPath := "items.sqlite"

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetWeaponById(id string) (*models.Weapon, error) {
	db, err := connectItemsDb()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT id, name, damage, damage_type, is_melee, is_ranged, is_two_handed, is_finesse FROM weapons WHERE id = ?"
	row := db.QueryRow(query, id)

	var weapon models.Weapon
	err = row.Scan(
		&weapon.ID,
		&weapon.Name,
		&weapon.Damage,
		&weapon.DamageType,
		&weapon.IsMelee,
		&weapon.IsRanged,
		&weapon.IsTwoHanded,
		&weapon.IsFinesse)
	if err != nil {
		return nil, err
	}

	return &weapon, nil
}

func GetArmorById(id string) (*models.Armor, error) {
	db, err := connectItemsDb()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT id, name, type, ac, max_dex_mod, required_strength FROM armor WHERE id = ?"
	row := db.QueryRow(query, id)

	var armor models.Armor
	err = row.Scan(&armor.ID, &armor.Name, &armor.Type, &armor.AC, &armor.MaxDexMod, &armor.RequiredStrength)
	if err != nil {
		return nil, err
	}

	return &armor, nil
}
