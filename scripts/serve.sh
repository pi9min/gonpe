#!/bin/bash

# main function
main () {
    port=${1}

    echo "############################### serve... ###############################"
    appyml_path="${PROJ_ROOT}/appengine/app.yaml"
    storage_dir="/tmp/ponpe_devappserver_storage"

    dev_appserver.py --port=${port} --storage_path=${storage_dir} ${appyml_path}
    echo
}

SCRIPT_DIR="$(cd "$(dirname "${0}")" && echo "${PWD}")"
PROJ_ROOT="${SCRIPT_DIR}/.."
main "$@"
