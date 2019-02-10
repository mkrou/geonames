package models

type UserTag struct {
	GeonameId int    `csv:"geonameId"`
	Name      string `csv:"tag"`
}
