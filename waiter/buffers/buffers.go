package buffers

import "dmbb.com/go2/common/model"

var ReadyOrderItems = make(chan *model.ReadyOrderItem, 100)
