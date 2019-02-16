package models

type AdminDivision struct {
	Code      string `csv:"code" valid:"required"`
	Name      string `csv:"name"`
	AsciiName string `csv:"ascii name" valid:"required"`
	GeonameId int    `csv:"geonameId" valid:"required"`
}
