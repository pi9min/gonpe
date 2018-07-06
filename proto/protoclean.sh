#!/bin/sh

SCRIPT_DIR="$(cd "$(dirname "${0}")" && echo "${PWD}")"

rm -f ${SCRIPT_DIR}/*pb*.go
