package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	dbPath := "/resources/items/items.sqlite"

	s := strings.Split(currentDir, "/")

	var output string

	for i := 0; i < len(s)-3; i++ {
		output += "/"
		output += s[i]
	}

	output += dbPath

	output = output[1:]

	fmt.Println(output)

	db, err := sql.Open("sqlite3", output)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	dropTableQuery := `DROP TABLE IF EXISTS weapons;`

	_, err = db.Exec(dropTableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS weapons (
		id TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL UNIQUE,
		damage TEXT NOT NULL,
		damage_type TEXT NOT NULL,
		is_melee NUMERIC NOT NULL,
		is_ranged NUMERIC NOT NULL,
		is_two_handed NUMERIC NOT NULL,
		is_finesse NUMERIC NOT NULL,
		PRIMARY KEY(id)
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	weaponsData := []struct {
		ID          string
		Name        string
		Damage      string
		DamageType  string
		IsMelee     int
		IsRanged    int
		IsTwoHanded int
		IsFinesse   int
	}{
		{"bulawa", "Bulawa", "1d6", "obuchowe", 1, 0, 0, 0},
		{"drag", "Drag", "1k6", "obuchowe", 1, 0, 0, 0},
		{"lekki_mlot", "Lekki mlot", "1k4", "obuchowe", 1, 0, 0, 0},
		{"maczuga", "Maczuga", "1k8", "obuchowe", 1, 0, 0, 0},
		{"toporek", "Toporek", "1k6", "ciete", 1, 0, 0, 0},
		{"sztylet", "sztylet", "1k4", "klute", 1, 0, 0, 1},
		{"krotki_luk", "Krotki luk", "1k6", "klute", 0, 1, 1, 0},
		{"dlugi_luk", "Dlugi luk", "1k8", "klute", 0, 1, 1, 0},
		{"kusza_lekka", "Kusza lekka", "1k8", "klute", 0, 1, 1, 0},
		{"kiscien", "Kiscien", "1k8", "obuchowe", 1, 0, 0, 0},
		{"miecz_dlugi", "Miecz dlugi", "1k8", "ciete", 1, 0, 0, 0},
		{"halabarda", "Halabarda", "1k10", "ciete", 1, 0, 1, 0},
		{"mlot_bojowy", "Mlot bojowy", "1k8", "obuchowe", 1, 0, 0, 0},
		{"mlot_dwureczny", "Mlot dwureczny", "2k6", "", 1, 0, 1, 0},
		{"topor_dwureczny", "Topr dwureczny", "1k12", "ciete", 1, 0, 1, 0},
		{"rapier", "Rapier", "1k8", "klute", 1, 0, 0, 1},
		{"sejmitar", "Sejmitar", "1k6", "ciete", 1, 0, 0, 1},
		{"miecz_krotki", "Miecz krotki", "1k6", "ciete", 1, 0, 0, 1},
	}

	for _, weapon := range weaponsData {
		_, err := db.Exec("INSERT INTO weapons (id, name, damage, damage_type, is_melee, is_ranged, is_two_handed, is_finesse) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			weapon.ID, weapon.Name, weapon.Damage, weapon.DamageType, weapon.IsMelee, weapon.IsRanged, weapon.IsTwoHanded, weapon.IsFinesse)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Wstawiono dane do tabeli weapons.")
}
