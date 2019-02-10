package models

type AdminSubdivision struct {
	Code      string `csv:"concatenated codes"`
	Name      string `csv:"name"`
	AsciiName string `csv:"asciiname"`
	GeonameId int    `csv:"geonameId"`
}
