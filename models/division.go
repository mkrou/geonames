package models

type AdminDivision struct {
	Code      string `csv:"code"`
	Name      string `csv:"name"`
	AsciiName string `csv:"ascii name"`
	Geonameid int    `csv:"geonameId"`
}
