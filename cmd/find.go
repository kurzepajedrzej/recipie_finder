package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"github.com/kurzepajedrzej/recipes_finder/cmd/find_components"
	"strings"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)


var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Command used for finding recipies by ingredients", 
	Long:  `Command used for finding recipies by ingredients. 
	It takes two arguments comma separated ingredients,
	and second argument how many recipies you want to find

	EXAMPLE : go main.go find ingredient1,ingredient2,... count
	`,   
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Parsing string arguments
		items := strings.Split(args[0], ",")

		// Parsing integer argument
		count, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid integer argument:", args[1])
			return
		}
	
		//Load API_KEY
		godotenv.Load()
		db := find_components.DbConnect()
		ingredients := items

	
		var response = find_components.SearchRecipesResponse{}
		dbResponse := find_components.DbFindResponse(db, ingredients, count)
		if dbResponse != nil {
			fmt.Println("Found response in DB!")
			response = *dbResponse
		} else {
			response = find_components.ApiRequest(strings.Join(ingredients, ","), count)
			find_components.DbSaveResponse(db, ingredients, count, response)
		}
	
		var recipes = find_components.ParseResponseToRecipes(response, ingredients)
		sort.Sort(find_components.ByMissingIngredientsLength(recipes))
	
		for i := 0; i < len(recipes); i++ {
			find_components.FinalPrint(recipes[i])
		}

		
	},
}


func init() {
	rootCmd.AddCommand(findCmd)
}
