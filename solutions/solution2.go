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
)

func Solution2(filePath string, out io.Writer) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stations := make(map[string]*Station)

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		// Split the line into station and data
		name, data, foundSemi := strings.Cut(line, ";")
		assert.AssertWithErrorAndContext(foundSemi,
			"no semi-colon found in line: %s",
			nil,
			assert.ErrorContext{Name: "Line", Value: line})

		// Parse the data to a float
		value, err := fastfloat.Parse(data)
		assert.AssertWithError(err == nil, "error parsing data to float", err)

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

	return nil
}
