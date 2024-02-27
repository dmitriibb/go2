package workers

import "dmbb.com/go2/common/logging"

var loggerRepo = logging.NewLogger("WorkersRepository")

func saveOrderItemWrapper(wrapper *OrderItemWrapper) {
	loggerRepo.Debug("save - %v", wrapper)
}
