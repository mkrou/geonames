package models

type GeonameDelete struct {
	Id      int    `csv:"geonameId" valid:"required"`
	Name    string `csv:"name" valid:"required"`
	Comment string `csv:"comment"`
}
