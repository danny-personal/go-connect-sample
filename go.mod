module github.com/danny-personal/go-connect-sample

go 1.20

require (
	github.com/bufbuild/connect-go v1.7.0
	golang.org/x/net v0.9.0
	google.golang.org/protobuf v1.30.0
)

require golang.org/x/text v0.9.0 // indirect

replace github.com/danny-personal/go-connect-sample/gen/greet/v1/greetv1connect => ../gen/greet/v1/greetv1connect
