package workers

import (
	"dmbb.com/go2/common/logging"
	commonInitializer "dmbb.com/go2/common/utils/initializer"
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
