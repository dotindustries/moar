#!/bin/bash
set -eu

wait_for_url () {
    echo "Testing $1"
    max_in_s=$2
    delay_in_s=1
    total_in_s=0
    while [ $total_in_s -le "$max_in_s" ]
    do
        if (echo -e "GET $1\n{}\nHTTP 200" | hurl) then
            return 0
        fi
        total_in_s=$(( total_in_s +  delay_in_s))
        echo "Wait ${total_in_s}s"
        sleep $delay_in_s
    done
    return 1
}

echo "Starting moar instance to be ready"
wait_for_url 'http://app:8000/' 60