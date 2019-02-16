package models

type Country struct {
	Iso2Code           string  `csv:"ISO" valid:"required"`
	Iso3Code           string  `csv:"ISO3" valid:"required"`
	IsoNumeric         string  `csv:"ISO-Numeric" valid:"required"`
	Fips               string  `csv:"fips"`
	Name               string  `csv:"Country" valid:"required"`
	Capital            string  `csv:"Capital"`
	Area               float64 `csv:"Area(in sq km)"`
	Population         int     `csv:"Population"`
	Continent          string  `csv:"Continent" valid:"required"`
	Tld                string  `csv:"tld"`
	CurrencyCode       string  `csv:"CurrencyCode"`
	CurrencyName       string  `csv:"CurrencyName"`
	Phone              string  `csv:"Phone"`
	PostalCodeFormat   string  `csv:"Postal Code Format"`
	PostalCodeRegex    string  `csv:"Postal Code Regex"`
	Languages          string  `csv:"Languages"`
	GeonameID          int     `csv:"geonameid" valid:"required"`
	Neighbours         string  `csv:"neighbours"`
	EquivalentFipsCode string  `csv:"EquivalentFipsCode"`
}
