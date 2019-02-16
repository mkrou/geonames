package models

type GeonameDelete struct {
	Id      int    `csv:"geonameId"`
	Name    string `csv:"name"`
	Comment string `csv:"comment"`
}
