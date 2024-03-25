package buffers

import "github.com/dmitriibb/go-common/restaurant-common/model"

var ReadyOrderItems = make(chan *model.ReadyOrderItem, 100)
