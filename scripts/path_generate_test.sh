#!/bin/bash

# Generator test
# Get Root Dir
ROOT_DIR=$(pwd | sed 's/\/scripts//g')
PATH_GENERATE_DIR=$ROOT_DIR/generate/methods

echo "Running tests for path generator"
echo ROOT_DIR=$ROOT_DIR
go test $PATH_GENERATE_DIR -v
