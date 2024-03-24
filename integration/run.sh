#!/bin/bash
set -eu

echo "Starting dev env"
docker-compose -f ./docker/docker-compose-dev.yml up -d
echo "Starting moar registry"
# docker run --name moar --rm --detach --publish 8000:8000 dotindustries/moar-registry:latest
S3_ACCESS_KEY_ID=minio S3_SECRET_ACCESS_KEY=minio123 go run cli/main.go -c docker/.dev.yaml up -d > /dev/null 2>&1 &

source integration/wait.sh

source integration/test.sh

echo "Stopping moar instance"
kill %1
docker-compose -f ./docker/docker-compose-dev.yml down