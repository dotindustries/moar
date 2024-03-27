#!/bin/bash
set -eu

echo "Starting dev env"
docker-compose -f ./docker/docker-compose-dev.yml up -d

source integration/wait.sh

source integration/test.sh

echo "Stopping dev environment"
docker-compose -f ./docker/docker-compose-dev.yml down