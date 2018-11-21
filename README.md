# gonpe

gonpe is video management server written by golang.

# Setup

installation

```
# language, library
brew install go (ver 1.11+)
brew install dep (ver 0.5+)
brew install protobuf (ver 3+)
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/nametake/protoc-gen-gohttp

# Google Cloud SDK
curl https://sdk.cloud.google.com | bash
```

# Install dependency

`make dep`

# Serve

`make run`

# Gen proto

`make gen`

# Clean

`make clean`
