#!/bin/bash
set -eu

echo "Running Hurl tests"
hurl integration/*.hurl --test