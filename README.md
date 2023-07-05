# Recipes Finder

## How to Run the Program

### 1. Fill in the API_KEY variable in the .env file with your own API key from [spoonacula.com/food_api](https://spoonacula.com/food_api). (Note: It's not safe to store API keys on GitHub or other remote repositories)

```
API_KEY=[your_api_key]
```

### 2. Make sure you have installed the `go` package on your machine.

### 3. You can start the program in two ways. The first method uses two commands, and the second method uses only one command.

- Method 1 (Using commands in the root project directory):
```
go build .
./recipes_finder find ingredient1,ingredient2,... count
```

- Method 2:
```
go run main.go find ingredient1,ingredient2,... count
```

The first argument of the `find` command should be the ingredients separated by commas.

The second argument (`count`) of the `find` command is the number of recipes you want to find.

### 4. After running the `recipes_finder` app, the recipes that you see in the command line will be saved to the `recipes.db` file.








 
