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
    -v go-1.12:/go \
    -w /app \
    golang:1.12.9-alpine /bin/sh -c \
    '
        set -xe;
        apk add git;
        go build -o /terraform/terraform-provider-algolia;
        chown ${UID}:${GID} /terraform/terraform-provider-algolia;
        chmod +x /terraform/terraform-provider-algolia
    '
