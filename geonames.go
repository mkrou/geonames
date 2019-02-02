package geonames

import (
	"geonames/models"
	"io"
	"net/http"
)

const (
	Url = "https://download.geonames.org/export/dump/"
)

type Parser func(file string) (io.ReadCloser, error)

func NewParser() Parser {
	return Parser(func(file string) (io.ReadCloser, error) {
		resp, err := http.Get(Url + file)
		if err != nil {
			return nil, err
		}

		return resp.Body, nil
	})
}

func (p Parser) GetGeonames(archive GeoNameFile, handler func(*models.Geoname) error) error {
	r, err := p(string(archive))
	if err != nil {
		return err
	}

	return Stream(r, func(_ string, columns []string) error {
		model, err := models.ParseGeoname(columns)
		if err != nil {
			return err
		}

		return handler(model)
	})
}
