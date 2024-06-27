package solutions

import (
	"bufio"
	"fmt"
	"github.com/NielsDobbelaar/goUtils/assert"
	"github.com/valyala/fastjson/fastfloat"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
)

func Solution4(filePath string, out io.Writer) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// define block size and number of blocks
	const numLines = 1000000000
	const blockSize = 4000000
	numBlocks := (numLines + blockSize - 1) / blockSize

	// define channels for the routines
	blocksChan := make(chan []string, numBlocks)
	resultsChan := make(chan map[string]*Station, numBlocks)

	// define waitgroup for the routines
	var wg sync.WaitGroup
	// Start worker goroutines
	for i := 0; i < numBlocks; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Create map for stations in this block
			stations := make(map[string]*Station)

			// Receive block of lines
			lines := <-blocksChan

			for _, line := range lines {

				// Split the line into station and data without using strings.Cut
				name, data, foundSemi := strings.Cut(line, ";")
				assert.AssertWithErrorAndContext(foundSemi,
					"no semi-colon found in line: %s",
					nil,
					assert.ErrorContext{Name: "Line", Value: line})

				// Parse the data to a float using fast float
				value, err := fastfloat.Parse(data)
				assert.AssertWithError(err == nil, "error parsing data to float", err)

				// Update the station's values
				station := stations[name]
				if station != nil {
					station.min = min(value, station.min)
					station.max = max(value, station.max)
					station.sum += value
					station.count++
				} else {
					stations[name] = &Station{
						min:   value,
						max:   value,
						sum:   value,
						count: 1,
					}
				}
			}

			// Send result back to main goroutine
			resultsChan <- stations
		}()
	}

	// Read the file in blocks and send to workers
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, blockSize)
	lineCount := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		lineCount++
		// If reached blockSize or end of file, send block to worker
		if lineCount%blockSize == 0 || lineCount == 1000000000 {
			blocksChan <- lines
			lines = make([]string, 0, blockSize)
		}
	}

	// Wait for all workers to finish
	close(blocksChan)
	wg.Wait()
	close(resultsChan)

	// merge results from all blocks
	finalStations := make(map[string]*Station)
	for result := range resultsChan {
		for name, station := range result {
			if _, ok := finalStations[name]; !ok {
				finalStations[name] = &Station{
					min:   station.min,
					max:   station.max,
					sum:   station.sum,
					count: station.count,
				}
			} else {
				finalStations[name].min = min(station.min, finalStations[name].min)
				finalStations[name].max = max(station.max, finalStations[name].max)
				finalStations[name].sum += station.sum
				finalStations[name].count += station.count
			}
		}
	}

	// Get the station names
	stationNames := make([]string, 0, len(finalStations))
	for name := range finalStations {
		stationNames = append(stationNames, name)
	}
	sort.Strings(stationNames)

	// Print output
	fmt.Fprint(out, "{")
	for i, name := range stationNames {
		if i > 0 {
			fmt.Fprint(out, ", ")
		}
		station := finalStations[name]
		fmt.Fprintf(out, "%s=%.1f/%.1f/%.1f", name, station.min, station.sum/float64(station.count), station.max)
	}
	fmt.Fprintln(out, "}")

	return nil
}
