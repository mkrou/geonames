package models

type TimeZone struct {
	Id          string  `csv:"TimeZoneId" valid:"required"`
	CountryCode string  `csv:"CountryCode" valid:"required"`
	GmtOffset   float64 `csv:"GMT offset 1. Jan 2019"`
	DstOffset   float64 `csv:"DST offset 1. Jul 2019"`
	RawOffset   float64 `csv:"RawOffset (independant of DST)"`
}
