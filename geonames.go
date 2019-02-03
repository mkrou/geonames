package geonames

import (
	"geonames/models"
	"io"
	"net/http"
	"strings"
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

	return Stream(r, defaultFilename(string(archive)), func(columns []string) error {
		model, err := models.ParseGeoname(columns)
		if err != nil {
			return err
		}

		return handler(model)
	})
}

func (p Parser) GetAlternames(handler func(*models.Altername) error) error {
	r, err := p(string(AlternateNames))
	if err != nil {
		return err
	}

	return Stream(r, defaultFilename(string(AlternateNames)), func(columns []string) error {
		model, err := models.ParseAltername(columns)
		if err != nil {
			return err
		}

		return handler(model)
	})
}

func defaultFilename(archive string) string {
	return strings.Replace(string(archive), ".zip", ".txt", 1)
}
