// Code generated by aocgen; DO NOT EDIT.
package year2023

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct{}

type Converter struct {
	Source int
	Dest   int
	Range  int
}

func (c Converter) getDest(value int) (int, bool) {
	if value >= c.Source && value < c.Source+c.Range {
		// TODO: How can we get this tranformation?
		// Value minus source to get diff and add that to dest?

		return c.Dest + (value - c.Source), true
	}
	return 0, false
}

type Converters struct {
	converters []Converter
}

func GetDestForConverters(source int, converters []Converter) (int, bool) {
	for _, c := range converters {
		if v, ok := c.getDest(source); ok {
			return v, true
		}
	}
	return 0, false
}

var converterOrder = []string{
	"seed-to-soil map:",
	"soil-to-fertilizer map:",
	"fertilizer-to-water map:",
	"water-to-light map:",
	"light-to-temperature map:",
	"temperature-to-humidity map:",
	"humidity-to-location map:",
}

func (p Day05) PartA(lines []string) any {
	seeds := strings.ReplaceAll(lines[0], "seeds: ", "")
	seedsInts := []int{}
	for _, seed := range strings.Fields(seeds) {
		seedInt, _ := strconv.Atoi(seed)
		seedsInts = append(seedsInts, seedInt)
	}
	return produceLocsForSeeds(lines, seedsInts)
}

func produceLocsForSeeds(lines []string, seeds []int) any {
	lines = lines[1:]
	converters := [][]Converter{}

	currentConverterCounter := -1
	for _, line := range lines {
		if line == "" {
			continue
		}

		if regexp.MustCompile("[A-Za-z-] map:").MatchString(line) {
			currentConverterCounter++
			continue
		}

		parts := strings.Fields(line)
		dest, _ := strconv.Atoi(parts[0])
		source, _ := strconv.Atoi(parts[1])
		rangeOfBoth, _ := strconv.Atoi(parts[2])
		if len(converters) != currentConverterCounter+1 {
			converters = append(converters, []Converter{})
		}
		converters[currentConverterCounter] = append(converters[currentConverterCounter], Converter{
			Source: source,
			Dest:   dest,
			Range:  rangeOfBoth,
		})

	}
	locations := []int{}
	for _, seed := range seeds {
		currentValue := seed
		for currentConvertersIndex, cs := range converters {
			if v, ok := GetDestForConverters(currentValue, cs); ok {
				currentValue = v
			}

			if currentConvertersIndex == 6 {
				locations = append(locations, currentValue)
			}
		}
	}

	return slices.Min(locations)
}

type SeedRange struct {
	Start int
	Range int
}

func (r SeedRange) ComputeAllValues()[]int {
	values := []int{}
	for i := r.Start;i< r.Start + r.Range;i++ {
		values = append(values, i)
	}
	return values
}


func parseSeedRanges(seeds string) []SeedRange {
	seedRanges := []SeedRange{}
	first := 0
	for _, s := range strings.Fields(seeds ) {
		if first == 0 {
			first, _ = strconv.Atoi(s)
			continue
		}

		second, _ := strconv.Atoi(s)
		seedRanges = append(seedRanges, SeedRange{
			Start: first,
			Range: second,
		})

		first = 0
	}
	return seedRanges
}
func (p Day05) PartB(lines []string) any {
	seeds := parseSeedRanges(strings.ReplaceAll(lines[0], "seeds: ", ""))

	allSeeds := []int {}
	for _, seedRange := range seeds {
		allSeeds = append(allSeeds,seedRange.ComputeAllValues()...)
		fmt.Println("Computed values")
	}

	return produceLocsForSeeds(lines, allSeeds)
}
