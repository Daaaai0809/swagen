#!/bin/bash

# Generator test
# Get Root Dir
ROOT_DIR=$(pwd | sed 's/\/scripts//g')
MESSAGE_GENERATE_DIR=$ROOT_DIR/generate/messages

echo "Running tests for message generator"
echo ROOT_DIR=$ROOT_DIR
go test $MESSAGE_GENERATE_DIR -v
