package sqldriver

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func Connect() (*sql.DB, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil, err
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
		return nil, err
	}

	return db, nil
}
