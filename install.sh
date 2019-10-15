#!/usr/bin/env bash

set -e

readonly PROJECT_DIRECTORY=$(realpath $(dirname $0))
readonly OS=linux
readonly ARCH=amd64
# https://www.terraform.io/docs/configuration/providers.html#third-party-plugins
readonly TERRAFORM_PLUGINS_DIRECTORY=${HOME}/.terraform.d/plugins/${OS}_${ARCH}

mkdir -p ${TERRAFORM_PLUGINS_DIRECTORY}

docker run \
    -e UID=$(id -u) \
    -e GID=$(id -g) \
    -e CGO_ENABLED=0 \
    -v ${PROJECT_DIRECTORY}:/app \
    -v ${HOME}/.terraform.d/plugins/linux_amd64:/terraform \
    -v go-1.13:/go \
    -w /app \
    golang:1.13.1-alpine /bin/sh -c \
    '
        set -xe;
        apk add git;
        TERRAFORM_PLUGIN_NAME=terraform-provider-algolia;
        if [[ ! -z $(git describe --exact-match --tags HEAD 2> /dev/null) ]]; then \
            TERRAFORM_PLUGIN_NAME=${TERRAFORM_PLUGIN_NAME}_$(git describe --exact-match --tags HEAD 2> /dev/null); \
        fi;
        go build -o /terraform/${TERRAFORM_PLUGIN_NAME};
        chown ${UID}:${GID} /terraform/${TERRAFORM_PLUGIN_NAME};
        chmod +x /terraform/${TERRAFORM_PLUGIN_NAME}
    '
