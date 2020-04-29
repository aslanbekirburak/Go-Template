#!/usr/bin/env bash
set -e

DOCKER_REGISTRY='104.248.249.210:5000'
IMAGE_NAME='iugo-layout/backend'
IMAGE_VERSION='0.0.1'

if [ "$1" != "" ]
then
  IMAGE_VERSION=$1
fi

REGISTRY_IMAGE="$DOCKER_REGISTRY/$IMAGE_NAME:$IMAGE_VERSION"
GOOS=linux GOARCH=amd64 go build ./apps/Dashboard
docker build -t ${REGISTRY_IMAGE} ./

docker push ${REGISTRY_IMAGE}