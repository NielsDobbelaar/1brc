#!/bin/bash

# Compile the Go program
go build -ldflags="-s -w" -o main main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
  # Run the compiled program
  ./main
else
  echo "Build failed."
fi

