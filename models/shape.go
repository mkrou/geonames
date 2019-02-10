package models

type Shape struct {
	GeonameId int    `csv:"geoNameId"`
	GeoJson   string `csv:"geoJSON"`
}
