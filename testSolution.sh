#!/bin/bash

echo "Running solution test"

# Compile the Go program
go build -ldflags="-s -w" -o main main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
  # Run the compiled program
  diff_output=$(diff --color=always <(./main) correct.txt)

  if [ -z "$diff_output" ]; then
    echo "Passed"
  else
    echo "Failed"
    echo "$diff_output"
  fi
else
  echo "Build failed."
fi


