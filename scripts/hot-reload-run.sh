#!/bin/bash

BINARY=service.bin
DIRECTORY=cmd/service

if [[ ! -x "$(command -v inotifywait)" ]]; then
    echo >&2 "inotify-tools is required"
    exit 1
fi

function start_runner {
    # Building
    cd ${DIRECTORY}
    echo -e "[Building]"
    go build  -o ${BINARY}
    BUILD_PID=$!
    trap "rm -f ./${BINARY}" SIGINT SIGTERM EXIT
    trap "kill -2 $(jobs -p) &> /dev/null" SIGINT SIGTERM EXIT
    wait ${BUILD_PID} &> /dev/null
    cd - > /dev/null

    # Executing
    echo -e "[Starting]"
    exec ./${DIRECTORY}/${BINARY}
}

function stop_runner {
    if kill -0 ${PID} 2> /dev/null; then # If pid is still running
        echo -e "[Stopping]"
        kill -2 ${PID} 2> /dev/null
        wait ${PID} 2> /dev/null
        rm -f ./${DIRECTORY}/${BINARY}
    fi
}

function main {
    HOTRELOAD=true

    function stop_hotreload {
        stop_runner ${PID}
        HOTRELOAD=false
    }
    trap "stop_hotreload" SIGINT SIGTERM EXIT

    while ${HOTRELOAD}; do
        start_runner &
        pid=$!

        inotifywait -e modify -e move -e create -e delete -r --include "\.(go)$" . > /dev/null 2> /dev/null

        stop_runner ${PID}
    done
}

main