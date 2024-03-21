package buffers

import "github.com/dmitriibb/go2/common/model"

var ReadyOrderItems = make(chan *model.ReadyOrderItem, 100)
