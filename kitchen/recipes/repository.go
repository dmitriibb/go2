package recipes

import (
	"context"
	"dmbb.com/go2/common/db/mongo"
	"dmbb.com/go2/common/logging"
	"dmbb.com/go2/common/utils"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kitchen/constants"
)

const (
	recipes_collection = "recipes"
)

var logger = logging.NewLogger("RecipesRepository")

// store in json
var inMemoryData = make(map[string]string)
var dbName = ""

var initialized = false

func Init() {
	if initialized {
		return
	}
	dbName = utils.GetEnvProperty(constants.MongoDbNameEnv)
	initData()
	initialized = true
	logger.Debug("initialized")
}

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

func initData() {
	mongo.TestConnection()
	saveInMongo(recipeBurger())
	saveInMongo(recipeCola())
	saveInMongo(recipeCoffee())
	saveInMongo(recipeWater())
	saveInMongo(recipeBread())
	saveInMongo(recipePasta())
}

func saveInMemory(recipe RecipeStage) {
	v, _ := json.Marshal(recipe)
	inMemoryData[recipe.Name] = string(v)
}

func saveInMongo(recipe RecipeStage) {
	client := mongo.GetClient()
	defer client.Disconnect(context.TODO())
	collection := client.Database(dbName).Collection(recipes_collection)

	filter := bson.D{{"name", recipe.Name}}
	update := bson.D{{"$set", recipe}}
	result, err := collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	//result, err := collection.InsertOne(context.TODO(), recipe)
	if err != nil {
		logger.Error("can't save recipe in DB because %v", err.Error())
		return
	}
	logger.Info("saved recipe '%v' in DB. Result: %v", recipe.Name, result)

}

func recipeBurger() RecipeStage {
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
	return burger
}
func recipeCola() RecipeStage {
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
	return cola
}
func recipeCoffee() RecipeStage {
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
	return coffee
}
func recipeWater() RecipeStage {
	water := RecipeStage{}
	water.Name = "water"
	water.Ingredients = []string{"water"}
	return water
}
func recipeBread() RecipeStage {
	water := RecipeStage{}
	water.Name = "bread"
	water.Ingredients = []string{"bread"}
	return water
}
func recipePasta() RecipeStage {
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
	return pasta
}
