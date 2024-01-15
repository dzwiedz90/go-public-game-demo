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

	dropTableQuery := `DROP TABLE IF EXISTS armor;`

	_, err = db.Exec(dropTableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	createTableQuery := `
	CREATE TABLE "armor" (
		"id"	TEXT NOT NULL UNIQUE,
		"name"	TEXT NOT NULL UNIQUE,
		"type"	TEXT NOT NULL,
		"ac"	INTEGER NOT NULL,
		"max_dex_mod"	INTEGER,
		"required_strength"	INTEGER,
		PRIMARY KEY("id")
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	armorData := []struct {
		ID               string
		Name             string
		Type             string
		AC               int
		MaxDexMod        int
		RequiredStrength int
	}{
		{"ar_kolczuga", "Kolczuga", "ciezki", 16, 0, 13},
		{"ar_skorznia", "Skorznia", "lekki", 11, 100, 0},
		{"ar_tarcza", "Tarcza", "tarcza", 2, 0, 0},
		{"ar_zbroja_luskowa", "Zbroja luskowa", "sredni", 14, 2, 0},
		{"ar_zbroja_skorzana", "Zbroja skorzana", "sredni", 12, 2, 0},
	}

	for _, armor := range armorData {
		_, err := db.Exec("INSERT INTO armor (id, name, type, ac, max_dex_mod, required_strength) VALUES (?, ?, ?, ?, ?, ?)",
			armor.ID, armor.Name, armor.Type, armor.AC, armor.MaxDexMod, armor.RequiredStrength)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Wstawiono dane do tabeli armors.")
}
