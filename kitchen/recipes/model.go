package recipes

// RecipeStage Used to store recipes and for tracing dish cooking progress
type RecipeStage struct {
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Ingredients   []string          `json:"ingredients"`
	TimeToWaitSec int64             `json:"timeToWaitSec"`
	TimeStarted   int64             `json:"timeStarted"`
	Status        RecipeStageStatus `json:"status"`
	Comment       string            `json:"comment"`
	SubStages     []*RecipeStage    `json:"subStages"`
}

type RecipeStageStatus string

const (
	RecipeStageStatusEmpty      RecipeStageStatus = ""
	RecipeStageStatusInProgress RecipeStageStatus = "InProgress"
	RecipeStageStatusFinished   RecipeStageStatus = "Finished"
	RecipeStageStatusError      RecipeStageStatus = "Error"
)
