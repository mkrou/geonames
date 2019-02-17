package models

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
	Id                    int     `csv:"geonameid" valid:"required"`
	Name                  string  `csv:"name" valid:"required"`
	AsciiName             string  `csv:"asciiname"`
	AlternateNames        string  `csv:"alternatenames"`
	Latitude              float64 `csv:"latitude"`
	Longitude             float64 `csv:"longitude"`
	Class                 string  `csv:"feature class"`
	Code                  string  `csv:"feature code"`
	CountryCode           string  `csv:"country code"`
	AlternateCountryCodes string  `csv:"cc2"`
	Admin1Code            string  `csv:"admin1 code"`
	Admin2Code            string  `csv:"admin2 code"`
	Admin3Code            string  `csv:"admin3 code"`
	Admin4Code            string  `csv:"admin4 code"`
	Population            int     `csv:"population"`
	Elevation             int     `csv:"elevation,omitempty"`
	DigitalElevationModel int     `csv:"dem,omitempty"`
	Timezone              string  `csv:"timezone"`
	ModificationDate      Time    `csv:"modification date" valid:"required"`
}
