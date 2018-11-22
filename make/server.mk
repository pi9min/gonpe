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
SERVER_ROOT := $(MAKE_DIR)

######## Vars
DEV_STORAGE_PATH := "/tmp/ponpe_devappserver_storage"
DEV_APPYAML_PATH := "$(ROOT)/appengine/app.local.yaml"
APPYAML_PATH := "$(ROOT)/appengine/app.yaml"
CRONYAML_PATH := "$(ROOT)/appengine/cron.yaml"
DISPATCHYAML_PATH := "$(ROOT)/appengine/dispatch.yaml"
PROTO_PATH := "$(ROOT)/server/proto"

######## Rules
dep:
	cd $(SERVER_ROOT); go mod download

clean:
	rm -f $(PROTO_PATH)/**/*{pb,http}.go

gen:
	protoc --proto_path=$(PROTO_PATH) --go_out=plugins=grpc:$(PROTO_PATH) --gohttp_out=$(PROTO_PATH) $(PROTO_PATH)/**/*.proto

test: gen
	cd $(SERVER_ROOT); go test ./...
	cd $(SERVER_ROOT); go vet ./...

run: gen
	dev_appserver.py --port=4040 $(DEV_APPYAML_PATH)

deploy: gen
	gcloud app deploy $(APPYAML_PATH)

deploy_cron:
	gcloud app deploy $(CRONYAML_PATH)

deploy_dispatch:
	gcloud app deploy $(DISPATCHYAML_PATH)
