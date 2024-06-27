#!/bin/bash

# Start timing
echo "Starting solution timing test."
echo "Builing..."
go build -ldflags="-s -w" -o main main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
  # Run the compiled program
  echo "Build successful. Running solution timing test."
  start_time=$(date +%s)
  ./main > /dev/null
  end_time=$(date +%s)
else
  echo "Build failed."
fi

# End timing
elapsed_time=$((end_time - start_time))

echo "Finished running solution timing test."
echo "Time taken: ${elapsed_time} seconds."

