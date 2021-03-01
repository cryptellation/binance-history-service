#!/bin/bash

SCRIPTS_PATH=$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )
PROJECT_PATH=$(realpath --relative-to=${PWD} ${SCRIPTS_PATH}/..)
API_PATH=${PROJECT_PATH}/api
PKG_PATH=${PROJECT_PATH}/pkg

function check-swagger {
	echo "Check OpenAPI generation tool..."
	which swagger || (go get -u github.com/go-swagger/go-swagger/cmd/swagger)
}

function generate-swagger {
	local SOURCE_DIR=${PKG_PATH}/server
	echo "Generate OpenAPI specifications from ${SOURCE_DIR} to ${API_PATH}..."
	mkdir -p ${API_PATH}
	swagger generate spec -o ${API_PATH}/swagger.json -w ${SOURCE_DIR} --scan-models
}

function serve-swagger {
	echo "Serve OpenAPI specifications at ${API_PATH}..."
	check-swagger
	swagger serve -F=swagger ${API_PATH}/swagger.json
}

function help {
	echo "$0: OpenAPI specifications utility"
	echo ""
	echo "Positional arguments:"
	echo "  check:        Check if the OpenAPI tool is preset. If not, install it."
	echo "  generate:     Generate the OpenAPI specifications from source code."
	echo "                Path is ${SWAGGER_FILE_PATH}."
	echo "  serve:        Serve a web server with OpenAPI specifications located"
	echo "                in ${SWAGGER_FILE_PATH}."
	echo "  help:         Display help."
}

while (( "$#" )); do
	case "$1" in
		check) CMD+="check-swagger ";;
		generate) CMD+="generate-swagger ";;
		serve) CMD+="serve-swagger ";;
		help) help;;
		*) CMD+="$1 ";;
	esac
	shift
done

eval $CMD