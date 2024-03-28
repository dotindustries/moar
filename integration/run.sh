#!/bin/bash
set -eu

echo "Starting dev env"
docker compose -f ./docker/docker-compose-dev.yml up --detach --wait

docker ps -a

source integration/test.sh

echo "Stopping dev environment"
docker compose -f ./docker/docker-compose-dev.yml down
