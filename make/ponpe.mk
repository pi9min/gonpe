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

# Set script path
MAKE_DIR:= $(dir $(lastword $(MAKEFILE_LIST)))
ROOT := $(realpath $(MAKE_DIR)/..)
export PATH := $(ROOT)/bin:$(ROOT)/scripts:$(ROOT)/scripts/lib:$(PATH)

######## Rules
DEFAULT_SERVICE := "default"

gen:
	(cd proto; ./protogen.sh)

clean:
	(cd proto; ./protoclean.sh)

dep:
	dep ensure -v

test:
	test.sh ./...

run:
	serve.sh 4040 $(DEFAULT_SERVICE)
