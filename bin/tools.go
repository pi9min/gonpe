// +build tools

package tools

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/nametake/protoc-gen-gohttp"
	_ "google.golang.org/grpc"
)
