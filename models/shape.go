package models

type Shape struct {
	GeonameId int    `csv:"geoNameId" valid:"required"`
	GeoJson   string `csv:"geoJSON" valid:"required"`
}
