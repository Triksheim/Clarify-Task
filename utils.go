package main

import (
	"bufio"
	"os"
	"sort"
)

func LoadLinesFromFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func FahrenheitToCelsius(f float64) float64 {
	c := (f - 32) * 5 / 9
	return c
}

func PsiToBar(psi float64) float64 {
	bar := psi * 0.0689476
	return bar
}

func GpmToLitrePerSecond(gpm float64) float64 {
	litrePerSecond := gpm * 0.0630902
	return litrePerSecond
}

func GetSortedKeys[T any](mapData map[string][]T) []string {
	keys := make([]string, 0, len(mapData))
	for k := range mapData {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
