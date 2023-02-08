package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	// ParseDuration can only identify units end with "ns", "us" (or "µs"), "ms", "s", "m", "h".
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	// parse unit d
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		// ndr: parse like 1d2h
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}
	// for pure number, unit ns
	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
