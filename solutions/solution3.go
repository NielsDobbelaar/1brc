package solutions

import (
	"bufio"
	"fmt"
	"github.com/NielsDobbelaar/goUtils/assert"
	"github.com/valyala/fastjson/fastfloat"
	"io"
	"os"
	// "runtime"
	"sort"
	"strings"
	"sync"
)

func Solution3(filePath string, out io.Writer) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stations := make(map[string]*Station)

	var mu sync.Mutex
	var wg sync.WaitGroup

	lines := make(chan string, 10000)

	// define number of goroutines
	numRoutines := 6

	// Read the file
	for range numRoutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range lines {
				// Split the line into station and data
				name, data, foundSemi := strings.Cut(line, ";")
				assert.AssertWithErrorAndContext(foundSemi,
					"no semi-colon found in line: %s",
					nil,
					assert.ErrorContext{Name: "Line", Value: line})

				// Parse the data to a float
				value, err := fastfloat.Parse(data)
				assert.AssertWithError(err == nil, "error parsing data to float", err)

				mu.Lock()
				// Check if the station is already in the map and update the values
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
				mu.Unlock()
			}
		}()
	}

	// Read the file lines and send to workers
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	// Wait for all workers to finish
	wg.Wait()

	// Get the station names
	stationNames := make([]string, len(stations))
	index := 0
	for name := range stations {
		stationNames[index] = name
		index++
	}

	// Sort the station names
	sort.Strings(stationNames)

	// print output
	fmt.Fprint(out, "{")
	for i, name := range stationNames {
		if i > 0 {
			fmt.Fprint(out, ", ")
		}
		station := stations[name]
		fmt.Fprintf(out, "%s=%.1f/%.1f/%.1f", name, station.min, station.sum/float64(station.count), station.max)
	}
	fmt.Fprintln(out, "}")

	// station, ok := stations["Port Vila"]
	// assert.AssertWithError(ok, "Port Vila not found in stations", errors.New("Port Vila not found in stations"))
	// fmt.Fprintf(out, "station = %v - ", station)

	return nil
}
