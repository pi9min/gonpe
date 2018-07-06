#!/bin/bash

# main function
main () {
    port=${1}
    services="${@:3:($#-2)}"
    if [ $# -lt 1 ]; then
        echo "usage: ./serve.sh [port] [default etc...]"
        return 1
    fi

    echo "############################### serve... ###############################"
    default_dir="${PROJ_ROOT}/entrypoint"
    mod_dir="${PROJ_ROOT}/entrypoint"

    serve_target+=" ${mod_dir}/app.yaml"
    storage_dir="/tmp/ponpe_devappserver_storage"
    dev_appserver.py --port=${port} --storage_path=${storage_dir} ${serve_target}
    echo
}

SCRIPT_DIR="$(cd "$(dirname "${0}")" && echo "${PWD}")"
PROJ_ROOT="${SCRIPT_DIR}/.."
main "$@"
