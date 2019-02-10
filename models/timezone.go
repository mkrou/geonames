package models

type TimeZone struct {
	Id          string  `csv:"TimeZoneId"`
	CountryCode string  `csv:"CountryCode"`
	GmtOffset   float64 `csv:"GMT offset 1. Jan 2019"`
	DstOffset   float64 `csv:"DST offset 1. Jul 2019"`
	RawOffset   float64 `csv:"RawOffset (independant of DST)"`
}
