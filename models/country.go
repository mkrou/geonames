package models

type Country struct {
	Iso2Code           string  `csv:"ISO"`
	Iso3Code           string  `csv:"ISO3"`
	IsoNumeric         string  `csv:"ISO-Numeric"`
	Fips               string  `csv:"fips"`
	Name               string  `csv:"Country"`
	Capital            string  `csv:"Capital"`
	Area               float64 `csv:"Area(in sq km)"`
	Population         int     `csv:"Population"`
	Continent          string  `csv:"Continent"`
	Tld                string  `csv:"tld"`
	CurrencyCode       string  `csv:"CurrencyCode"`
	CurrencyName       string  `csv:"CurrencyName"`
	Phone              string  `csv:"Phone"`
	PostalCodeFormat   string  `csv:"Postal Code Format"`
	PostalCodeRegex    string  `csv:"Postal Code Regex"`
	Languages          string  `csv:"Languages"`
	GeonameID          int     `csv:"geonameid"`
	Neighbours         string  `csv:"neighbours"`
	EquivalentFipsCode string  `csv:"EquivalentFipsCode"`
}
