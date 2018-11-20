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
MAKE_DIR:= $(dir $(lastword $(MAKEFILE_LIST)))
ROOT := $(realpath $(MAKE_DIR)/..)

######## Vars
DEV_STORAGE_PATH := "/tmp/ponpe_devappserver_storage"
APPYAML_PATH := "$(ROOT)/appengine/app.yaml"
PROTO_PATH := "$(ROOT)/proto"

######## Rules
dep:
	dep ensure -v

clean:
	rm -f $(PROTO_PATH)/**/*{pb,http}.go

gen:
	protoc --proto_path=$(PROTO_PATH) --go_out=plugins=grpc:$(PROTO_PATH) --gohttp_out=$(PROTO_PATH) $(PROTO_PATH)/**/*.proto

run: gen
	dev_appserver.py --port=4040 $(APPYAML_PATH)
