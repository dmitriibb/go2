package recipes

import (
	"dmbb.com/go2/common/logging"
	"encoding/json"
	"fmt"
)

var logger = logging.NewLogger("RecipesRepository")

// store in json
var inMemoryData = make(map[string]string)

func GetRecipe(name string) (RecipeStage, error) {
	data, ok := inMemoryData[name]
	if !ok {
		return RecipeStage{}, fmt.Errorf("recipe for '%v' not found", name)
	}
	var recipe RecipeStage
	err := json.Unmarshal([]byte(data), &recipe)
	if err != nil {
		return RecipeStage{}, err
	}
	return recipe, nil
}

func InitData() {
	addBurger()
	addCola()
	addCoffee()
	addWater()
	addPasta()
}
func addBurger() {
	burger := RecipeStage{}
	burger.Name = "burger"
	burger.Ingredients = []string{}
	burger.SubStages = []*RecipeStage{
		{
			Name:        "cut vegetables",
			Ingredients: []string{"tomato", "lettuce", "onion"},
		},
		{
			Name:          "grill meet",
			Ingredients:   []string{"beef"},
			TimeToWaitSec: 20,
		},
		{
			Name:        "assemble burger",
			Ingredients: []string{"burger bun", "mayo", "cheese"},
		},
	}

	v, _ := json.Marshal(burger)
	inMemoryData[burger.Name] = string(v)
}
func addCola() {
	cola := RecipeStage{}
	cola.Name = "cola"
	cola.Ingredients = []string{"cola", "ice"}
	cola.SubStages = []*RecipeStage{
		{
			Name:        "open cola",
			Ingredients: []string{"cola"},
		},
		{
			Name:        "add ice",
			Ingredients: []string{"ice"},
		},
	}
	v, _ := json.Marshal(cola)
	inMemoryData[cola.Name] = string(v)
}
func addCoffee() {
	coffee := RecipeStage{}
	coffee.Name = "coffee"
	coffee.Ingredients = []string{"coffee", "milk"}
	coffee.SubStages = []*RecipeStage{
		{
			Name:        "brew coffee",
			Ingredients: []string{"coffee"},
		},
		{
			Name:        "add milk",
			Ingredients: []string{"milk"},
		},
	}
	v, _ := json.Marshal(coffee)
	inMemoryData[coffee.Name] = string(v)
}
func addWater() {
	water := RecipeStage{}
	water.Name = "water"
	water.Ingredients = []string{"water"}
	v, _ := json.Marshal(water)
	inMemoryData[water.Name] = string(v)
}
func addPasta() {
	pasta := RecipeStage{}
	pasta.Name = "pasta"
	pasta.Ingredients = []string{"egg", "flour", "onion", "chicken", "salt", "butter"}
	pasta.SubStages = []*RecipeStage{
		{
			Name:        "make pasta",
			Ingredients: []string{"egg", "flour"},
			SubStages: []*RecipeStage{
				{
					Name:        "mix eggs and flour",
					Ingredients: []string{"egg", "flour"},
				},
				{
					Name:          "boil water",
					TimeToWaitSec: 30,
				},
				{
					Name:          "cook pasta in water",
					TimeToWaitSec: 50,
				},
			},
		},
		{
			Name:        "fry chicken",
			Ingredients: []string{"chicken", "onion"},
		},
		{
			Name:        "fry chicken with pasta",
			Ingredients: []string{"salt", "butter"},
		},
	}
	v, _ := json.Marshal(pasta)
	inMemoryData[pasta.Name] = string(v)
}
