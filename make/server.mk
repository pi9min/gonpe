######## Settings
# Disable makefile default values and rules
# Unconditionally make all targets.
MAKEFLAGS=--no-builtin-rules \
          --no-builtin-variables \
          --always-make

# Fixes a bug in OSX Make with exporting PATH environment variables
# See: http://stackoverflow.com/questions/11745634/setting-path-variable-from-inside-makefile-does-not-work-for-make-3-81
OS := $(shell uname -s)
ifeq ($(OS),Darwin)
export SHELL := $(shell echo $$SHELL)
endif

# set root path
MAKE_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT := $(realpath $(MAKE_DIR)/..)
SERVER_ROOT := $(ROOT)/server
export PATH := $(ROOT)/bin:$(PATH)
export GOBIN := $(ROOT)/bin

######## Vars
DEV_STORAGE_PATH := "/tmp/ponpe_devappserver_storage"
DEV_APPYAML_PATH := "$(ROOT)/appengine/app.local.yaml"
APPYAML_PATH := "$(ROOT)/appengine/app.yaml"
CRONYAML_PATH := "$(ROOT)/appengine/cron.yaml"
DISPATCHYAML_PATH := "$(ROOT)/appengine/dispatch.yaml"
PROTO_PATH := "$(ROOT)/proto"
SERVER_PROTO_PATH := "$(ROOT)/server/proto"

######## Rules
dep:
	cd $(ROOT); go mod download

clean_bin:
	find bin ! -name tools.go -type f -exec rm -f {} \;

install_bin:
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/lyft/protoc-gen-validate@master
	go get github.com/nametake/protoc-gen-gohttp
	go get google.golang.org/grpc

clean:
	rm -f $(SERVER_PROTO_PATH)/**/*{pb,http,validate}.go

gen:
	protoc \
	-I $(PROTO_PATH) \
	-I $(GOPATH)/src \
	-I $(GOPATH)/src/github.com/lyft/protoc-gen-validate \
	--go_out=plugins=grpc:$(SERVER_PROTO_PATH) \
	--gohttp_out=$(SERVER_PROTO_PATH) \
	--validate_out=lang=go:$(SERVER_PROTO_PATH) \
	$(PROTO_PATH)/**/*.proto

test: gen
	cd $(SERVER_ROOT); go test ./...
	cd $(SERVER_ROOT); go vet ./...

run: gen
	dev_appserver.py --support_datastore_emulator=False --port=4040 $(DEV_APPYAML_PATH)

deploy: gen
	gcloud app deploy $(APPYAML_PATH)

deploy_cron:
	gcloud app deploy $(CRONYAML_PATH)

deploy_dispatch:
	gcloud app deploy $(DISPATCHYAML_PATH)
