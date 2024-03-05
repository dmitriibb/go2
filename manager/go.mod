module dmbb.com/go2/manager

go 1.21

replace dmbb.com/go2/common => ../common
replace dmbb.com/go2/kitchen => ../kitchen

require (
	dmbb.com/go2/common v0.0.0
	dmbb.com/go2/kitchen v0.0.0
	google.golang.org/grpc v1.62.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240221002015-b0ce06bbee7c // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)
