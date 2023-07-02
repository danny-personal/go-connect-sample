module github.com/danny-personal/go-connect-sample

go 1.20

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.30.0-20230609233005-3757a25ff0b9.1
	github.com/bufbuild/connect-go v1.7.0
	github.com/bufbuild/protovalidate-go v0.1.1
	golang.org/x/net v0.9.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230512164433-5d1fd1a340c9 // indirect
	github.com/google/cel-go v0.16.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)

replace github.com/danny-personal/go-connect-sample/gen/greet/v1/greetv1connect => ../gen/greet/v1/greetv1connect

//replace github.com/danny-personal/go-connect-sample/gen/greet/v1 => ../gen/greet/v1
