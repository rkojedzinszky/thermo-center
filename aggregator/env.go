package aggregator

import (
	"os"
	"strconv"
)

func getenv(key string, def string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}

	return def
}

func getenvInt(key string, def int) int {
	if value, found := os.LookupEnv(key); found {
		value, err := strconv.Atoi(value)
		if err != nil {
			return value
		}
	}

	return def
}

func getenvFloat64(key string, def float64) float64 {
	if value, found := os.LookupEnv(key); found {
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return value
		}
	}

	return def
}
