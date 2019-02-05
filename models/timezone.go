package models

import (
	"strconv"
)

//CountryCode
// TimeZoneId
// GMT offset 1. Jan 2019
// DST offset 1. Jul 2019
// rawOffset (independant of DST)

const TimeZoneFields = 5

type TimeZone struct {
	Id          string
	CountryCode string
	gmtOffset   float64
	dstOffset   float64
	rawOffset   float64
}

func ParseTimeZone(parts []string) (*TimeZone, error) {
	gmt, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return nil, err
	}

	dst, err := strconv.ParseFloat(parts[3], 64)
	if err != nil {
		return nil, err
	}

	raw, err := strconv.ParseFloat(parts[4], 64)
	if err != nil {
		return nil, err
	}

	return &TimeZone{
		Id:          parts[1],
		CountryCode: parts[0],
		gmtOffset:   gmt,
		dstOffset:   dst,
		rawOffset:   raw,
	}, nil
}
