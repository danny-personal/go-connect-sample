# go-connect-sample
Getting started

## Usage
```bash
go run ./cmd/server/main.go
```

## Install
```bash
go mod init github.com/danny-personal/go-connect-sample
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
[ -n "$(go env GOBIN)" ] && export PATH="$(go env GOBIN):${PATH}"
[ -n "$(go env GOPATH)" ] && export PATH="$(go env GOPATH)/bin:${PATH}"
mkdir -p greet/v1
touch greet/v1/greet.proto
touch buf.gen.yaml
buf generate
go mod tidy
mkdir -p cmd/server
touch cmd/server/main.go
go get golang.org/x/net/http2
go get golang.org/x/net/http2
go get github.com/bufbuild/connect-go
go get go-connect-sample/greet/v1
go install go-connect-sample/greet/v1
go mod edit -replace github.com/danny-personal/go-connect-sample/gen/greet/v1/greetv1connect=../gen/greet/v1/greetv1connect
go mod edit -replace github.com/danny-personal/go-connect-sample/gen/greet/v1=../gen/greet/v1
go mod tidy
```
