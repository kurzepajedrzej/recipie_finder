package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "recipes_finder",
	Short: "Look for recipies in spoonacular API", 
	Long: `Program connects to remote database, look for recipies by igredients and ave searches 
to database.
`, 
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
