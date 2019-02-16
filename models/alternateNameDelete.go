package models

type AlternateNameDelete struct {
	Id        int    `csv:"alternateNameId" valid:"required"`
	GeonameId int    `csv:"geonameId" valid:"required"`
	Name      string `csv:"name" valid:"required"`
	Comment   string `csv:"comment"`
}
