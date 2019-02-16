package models

type AdminSubdivision struct {
	Code      string `csv:"concatenated codes" valid:"required"`
	Name      string `csv:"name" valid:"required"`
	AsciiName string `csv:"asciiname" valid:"required"`
	GeonameId int    `csv:"geonameId" valid:"required"`
}
