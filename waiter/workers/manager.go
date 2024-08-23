package workers

import (
	"context"
	"github.com/dmitriibb/go-common/logging"
	commonInitializer "github.com/dmitriibb/go-common/utils/initializer"
)

var logger = logging.NewLogger("WorkersManager")
var inittializer = commonInitializer.New(logger)
var workerNames = []string{"Dima", "John", "Karl", "Kate"}

func Init(rootContext context.Context) {
	inittializer.Init(func() error {

		for _, wName := range workerNames {
			worker := NewWorker(wName)
			ctx, _ := context.WithCancel(rootContext)
			worker.Start(ctx)
		}

		return nil
	})
}
