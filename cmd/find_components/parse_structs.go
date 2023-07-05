package find_components

import (
	"fmt"
	"strings"
)

type SearchRecipesResponse struct {
	Offset       int32                         `json:"offset"`
	Number       int32                         `json:"number"`
	Results      []SearchRecipesResponseResult `json:"results"`
	TotalResults int32                         `json:"totalResults"`
}

type SearchRecipesResponseResult struct {
	Title     string    `json:"title"`
	Nutrition Nutrition `json:"nutrition"`
}

type Nutrition struct {
	Nutrients   []Nutrient   `json:"nutrients"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Nutrient struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type Ingredient struct {
	Name string `json:"name"`
}

type Recipe struct {
	Title              string
	PresentIngredients []string
	MissingIngredients []string
	Calories           float64
	Fat                float64
	Protein            float64
}

func ParseResponseToRecipes(response SearchRecipesResponse, requestedIngredients []string) []Recipe {
	count := len(response.Results)
	recipes := make([]Recipe, count)
	for i := 0; i < count; i += 1 {
		recipes[i] = ParseResultToRecipe(response.Results[i], requestedIngredients)
	}
	return recipes
}

func ParseResultToRecipe(result SearchRecipesResponseResult, requestedIngredients []string) Recipe {
	recipe := Recipe{
		Title:              result.Title,
		PresentIngredients: requestedIngredients,
		MissingIngredients: ParseIngredientsToMissingIngredients(result.Nutrition.Ingredients, requestedIngredients),
		Calories:           ParseNutrientsToAmount(result.Nutrition.Nutrients, "Calories"),
		Fat:                ParseNutrientsToAmount(result.Nutrition.Nutrients, "Fat"),
		Protein:            ParseNutrientsToAmount(result.Nutrition.Nutrients, "Protein"),
	}
	return recipe
}

func ParseIngredientsToMissingIngredients(ingredients []Ingredient, presentIngredients []string) []string {
	missing := []string{}
	for i := 0; i < len(ingredients); i++ {
		ingredientName := ingredients[i].Name
		if !Contains(presentIngredients, ingredientName) {
			missing = append(missing, ingredientName)
		}
	}
	return missing
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ParseNutrientsToAmount(nutrients []Nutrient, name string) float64 {
	missing := float64(0)
	for i := 0; i < len(nutrients); i++ {
		nutrient := nutrients[i]
		if nutrient.Name == name {
			missing = nutrient.Amount
			break
		}
	}
	return missing
}

func FinalPrint(recipe Recipe) {
	fmt.Println(recipe.Title)
	fmt.Println("\tPresent Ingredients:\n\t\t-", strings.Join(recipe.PresentIngredients, "\n\t\t- "))
	fmt.Println("\tMissing Ingredients:\n\t\t-", strings.Join(recipe.MissingIngredients, "\n\t\t- "))
	fmt.Println("\tCalories:", recipe.Calories, "kcal", "Fat:", recipe.Fat, "g", "Protein:", recipe.Protein, "g")
	fmt.Println()
}

type ByMissingIngredientsLength []Recipe

func (r ByMissingIngredientsLength) Len() int {
	return len(r)
}
func (r ByMissingIngredientsLength) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r ByMissingIngredientsLength) Less(i, j int) bool {
	return len(r[i].MissingIngredients) < len(r[j].MissingIngredients)
}




