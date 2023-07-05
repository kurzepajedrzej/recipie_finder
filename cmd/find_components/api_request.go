package find_components

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)


func ApiRequest(ingredients string, count int) SearchRecipesResponse {
	apiKey := os.Getenv("API_KEY")
	url_template := "https://api.spoonacular.com/recipes/complexSearch?apiKey=%s&includeIngredients=%s&addRecipeNutrition=true&number=%d"
	url := fmt.Sprintf(url_template, apiKey, ingredients, count)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	
	defer resp.Body.Close()

	var result SearchRecipesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal("Error decoding response:", err)
	}

	return result
}
