package main

import (
	"1brc/solutions"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// use buffered output to avoid repeated system calls to stdout
	output := bufio.NewWriter(os.Stdout)

	err := solutions.Solution5("measurements.txt", output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	output.Flush()
}
