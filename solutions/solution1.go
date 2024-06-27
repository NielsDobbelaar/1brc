package solutions

import (
	"bufio"
	"fmt"
	"github.com/NielsDobbelaar/goUtils/assert"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Station struct {
	min, max, sum float64
	count         int64
}

func Solution1(filePath string, out io.Writer) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stations := make(map[string]Station)
	// Read the file
	scanner := bufio.NewScanner(file)
	// numLines := 0
	for scanner.Scan() {
		// numLines++
		// if numLines%10000000 == 0 {
		// 	fmt.Println(numLines)
		// }

		line := scanner.Text()

		// Split the line into station and data
		name, data, foundSemi := strings.Cut(line, ";")
		assert.AssertWithErrorAndContext(foundSemi,
			"no semi-colon found in line: %s",
			nil,
			assert.ErrorContext{Name: "Line", Value: line})

		// Parse the data to a float
		value, err := strconv.ParseFloat(data, 64)
		assert.AssertWithError(err == nil, "error parsing data to float", err)

		// Check if the station is already in the map and update the values
		station, ok := stations[name]
		if !ok {
			station.min = value
			station.max = value
			station.sum = value
			station.count = 1
		} else {
			station.max = max(value, station.max)
			station.min = min(value, station.min)
			station.sum += value
			station.count++
		}
		stations[name] = station

	}

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
