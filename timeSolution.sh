#!/bin/bash

# Start timing
start_time=$(date +%s)

# Run the solution script and compare the output with correct.txt using diff
diff_output=$(diff --color=always <(./runSolution.sh) correct.txt)

# End timing
end_time=$(date +%s)
elapsed_time=$((end_time - start_time))

# Check if diff_output is empty
if [ -z "$diff_output" ]; then
  echo "Correct. Time taken: ${elapsed_time} seconds."
else
  echo "Differences found:"
  echo "$diff_output"
  echo "Time taken: ${elapsed_time} seconds."
fi

