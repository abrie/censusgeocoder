package utils

import (
	"log"
	"strconv"
	"strings"
)

func mustConvertToFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalf("Fatal: '%s' is not a valid number.", str)
	}
	return num
}

func ParseCoordinates(coordinates *string) (float64, float64) {
	strs := strings.Split(*coordinates, ",")
	if len(strs) != 2 {
		log.Fatalf("Failed to parse '%s' as a lat long pair.", *coordinates)
	}

	x := mustConvertToFloat64(strs[0])
	y := mustConvertToFloat64(strs[1])

	return x, y
}
