package workers

import (
	"github.com/dmitriibb/go-common/logging"
	commonInitializer "github.com/dmitriibb/go-common/utils/initializer"
)

var logger = logging.NewLogger("WorkersManager")
var inittializer = commonInitializer.New(logger)
var workerNames = []string{"Dima", "John", "Karl", "Kate"}

func Init() {
	inittializer.Init(func() error {

		for _, wName := range workerNames {
			worker := NewWorker(wName)
			worker.Start()
		}

		return nil
	})
}
