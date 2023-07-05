package main

import (
	"github.com/kurzepajedrzej/recipes_finder/cmd"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	cmd.Execute()
}

