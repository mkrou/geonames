package geonames

type GeoNameFile string
type AltNameFile string

const (
	Cities500      GeoNameFile = "cities500.zip"
	Cities1000     GeoNameFile = "cities1000.zip"
	Cities5000     GeoNameFile = "cities5000.zip"
	Cities15000    GeoNameFile = "cities15000.zip"
	AllCountries   GeoNameFile = "allCountries.zip"
	AlternateNames AltNameFile = "alternateNamesV2.zip"
)
