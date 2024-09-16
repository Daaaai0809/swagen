#!/bin/bash

ROOT_DIR=$(pwd | sed 's/\/scripts//g')

# Check if yamlfmt is installed
if ! [ -x "$(command -v yamlfmt)" ]; then
  echo 'Error: yamlfmt is not installed.' >&2
  # Install yamlfmt
  read -p "Do you want to install yamlfmt? [y/n]: " answer
  if [ "$answer" = "y" ]; then
    go install github.com/google/yamlfmt/cmd/yamlfmt@latest
  else
    exit 1
  fi
fi

# Run yamlfmt
echo "Running format test yaml files"
yamlfmt -conf .yamlfmt $ROOT_DIR
