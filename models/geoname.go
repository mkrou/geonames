package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
geonameid         : integer id of record in geonames database
name              : name of geographical point (utf8) varchar(200)
asciiname         : name of geographical point in plain ascii characters, varchar(200)
alternatenames    : alternatenames, comma separated, ascii names automatically transliterated, convenience attribute from alternatename table, varchar(10000)
latitude          : latitude in decimal degrees (wgs84)
longitude         : longitude in decimal degrees (wgs84)
feature class     : see http://www.geonames.org/export/codes.html, char(1)
feature code      : see http://www.geonames.org/export/codes.html, varchar(10)
country code      : ISO-3166 2-letter country code, 2 characters
cc2               : alternate country codes, comma separated, ISO-3166 2-letter country code, 200 characters
admin1 code       : fipscode (subject to change to iso code), see exceptions below, see file admin1Codes.txt for display names of this code; varchar(20)
admin2 code       : code for the second administrative division, a county in the US, see file admin2Codes.txt; varchar(80)
admin3 code       : code for third level administrative division, varchar(20)
admin4 code       : code for fourth level administrative division, varchar(20)
population        : bigint (8 byte int)
elevation         : in meters, integer
dem               : digital elevation model, srtm3 or gtopo30, average elevation of 3''x3'' (ca 90mx90m) or 30''x30'' (ca 900mx900m) area in meters, integer. srtm processed by cgiar/ciat.
timezone          : the iana timezone id (see file timeZone.txt) varchar(40)
modification date : date of last modification in yyyy-MM-dd format
*/

type Geoname struct {
	Id                    int
	Name                  string
	AsciiName             string
	AlternateNames        []string
	Latitude              float64
	Longitude             float64
	Class                 string
	Code                  string
	CountryCode           string
	AlternateCountryCodes []string
	Admin1Code            string
	Admin2Code            string
	Admin3Code            string
	Admin4Code            string
	Population            int
	Elevation             int
	DigitalElevationModel int
	Timezone              string
	ModificationDate      time.Time
}

func ParseGeoname(parts []string) (*Geoname, error) {
	if len(parts) != 19 {
		return nil, fmt.Errorf("Line contains wrong number of columns: %d", len(parts))
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	aNames := strings.Split(parts[3], ",")

	latitude, err := strconv.ParseFloat(string(parts[4]), 64)
	if err != nil {
		return nil, err
	}

	longitude, err := strconv.ParseFloat(string(parts[5]), 64)
	if err != nil {
		return nil, err
	}

	acc := strings.Split(parts[9], ",")

	population := 0
	if parts[14] != "" {
		population, err = strconv.Atoi(parts[14])
		if err != nil {
			return nil, err
		}
	}

	elevation := 0
	if parts[15] != "" {
		elevation, err = strconv.Atoi(parts[15])
		if err != nil {
			return nil, err
		}
	}

	dem, err := strconv.Atoi(parts[16])
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("2006-01-02", parts[18])
	if err != nil {
		return nil, err
	}

	return &Geoname{
		Id:                    id,
		Name:                  parts[1],
		AsciiName:             parts[2],
		AlternateNames:        aNames,
		Latitude:              latitude,
		Longitude:             longitude,
		Class:                 parts[6],
		Code:                  parts[7],
		CountryCode:           parts[8],
		AlternateCountryCodes: acc,
		Admin1Code:            parts[10],
		Admin2Code:            parts[11],
		Admin3Code:            parts[12],
		Admin4Code:            parts[13],
		Population:            population,
		Elevation:             elevation,
		DigitalElevationModel: dem,
		Timezone:              parts[17],
		ModificationDate:      date,
	}, nil
}
