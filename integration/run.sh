#!/bin/bash
set -eu

wait_for_url () {
    echo "Testing $1"
    max_in_s=$2
    delay_in_s=1
    total_in_s=0
    while [ $total_in_s -le "$max_in_s" ]
    do
        echo "Wait ${total_in_s}s"
        if (echo -e "POST $1\n{}\nHTTP/* 200" | hurl > /dev/null 2>&1;) then
            return 0
        fi
        total_in_s=$(( total_in_s +  delay_in_s))
        sleep $delay_in_s
    done
    return 1
}

echo "Starting dev env"
docker-compose -f ./docker/docker-compose-dev.yml up -d
echo "Starting moar registry"
# docker run --name moar --rm --detach --publish 8000:8000 dotindustries/moar-registry:latest
S3_ACCESS_KEY_ID=minio S3_SECRET_ACCESS_KEY=minio123 go run cli/main.go -c docker/.dev.yaml up -d > /dev/null 2>&1 &

echo "Starting moar instance to be ready"
wait_for_url 'http://localhost:8000/moarpb.ModuleRegistry/GetModule' 60

echo "Running Hurl tests"
hurl integration/*.hurl --test

echo "Stopping moar instance"
kill %1
docker-compose -f ./docker/docker-compose-dev.yml down