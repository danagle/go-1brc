package v01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type SummaryFloat64 struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func Run(dataFilePath string) {
	dataFile, err := os.Open(dataFilePath)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()

	measurements := make(map[string]*SummaryFloat64)

	fileScanner := bufio.NewScanner(dataFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		rawString := fileScanner.Text()
		stationName, temperatureStr, found := strings.Cut(rawString, ";")
		if !found {
			continue
		}
		temperature, err := strconv.ParseFloat(temperatureStr, 32)
		if err != nil {
			panic(err)
		}

		measurement := measurements[stationName]
		if measurement == nil {
			measurements[stationName] = &SummaryFloat64{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temperature)
			measurement.Max = max(measurement.Max, temperature)
			measurement.Sum += temperature
			measurement.Count += 1
		}
	}

	printResultFloat64(measurements)
}

func printResultFloat64(data map[string]*SummaryFloat64) {
	locations := make([]string, 0, len(data))
	for location := range data {
		locations = append(locations, location)
	}

	slices.Sort(locations)

	fmt.Printf("{")
	for idx, location := range locations {
		measurement := data[location]
		mean := measurement.Sum / float64(measurement.Count)
		fmt.Printf("%s=%.1f/%.1f/%.1f",
			location, measurement.Min, mean, measurement.Max)
		if idx < len(locations)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Println("}")
}
