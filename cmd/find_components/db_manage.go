package find_components

import (
	"database/sql"
	"encoding/json"
	"strings"
	"log"
	_ "github.com/mattn/go-sqlite3"
)


func DbConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "recipes.db")
	if err != nil {
		log.Fatal(err)
	}

	const create string = `
	CREATE TABLE IF NOT EXISTS responses (
	ingredients TEXT NOT NULL,
	count INTEGER NOT NULL,
	response TEXT NOT NULL
	);`

	_, err2 := db.Exec(create)
	if err2 != nil {
		log.Fatal(err)
	}

	return db
}

type ResponseRow struct {
	Ingredients string
	Count       int
	Response    string
}

func DbFindResponse(db *sql.DB, ingredients []string, count int) *SearchRecipesResponse {
	row := db.QueryRow("SELECT * FROM responses WHERE ingredients=?", strings.Join(ingredients, ", "))

	var err error
	var responseRow ResponseRow
	if err = row.Scan(&responseRow.Ingredients, &responseRow.Count, &responseRow.Response); err == sql.ErrNoRows {
		return nil
	}

	// not enough recipes
	if responseRow.Count < count {
		return nil
	}

	var result SearchRecipesResponse
	err = json.Unmarshal([]byte(responseRow.Response), &result)
	if err != nil {
		log.Fatal("Error decoding response:", err)
	}

	return &result
}

func DbSaveResponse(db *sql.DB, ingredients []string, count int, response SearchRecipesResponse) {
	_, err := db.Exec("DELETE FROM responses WHERE ingredients=?;", strings.Join(ingredients, ","))
	if err != nil {
		log.Fatal("Error deleting existing response:", err)
	}

	encodedResponse, err2 := json.Marshal(response)
	if err2 != nil {
		log.Fatal("Error encoding response:", err2)
	}

	_, err3 := db.Exec("INSERT INTO responses VALUES(?,?,?);", strings.Join(ingredients, ","), count, encodedResponse)
	if err3 != nil {
		log.Fatal("Error saving response to DB", err)
	}
}