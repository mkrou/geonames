package geonames

import (
	"github.com/jszwec/csvutil"
	"github.com/mkrou/geonames/models"
	"github.com/mkrou/geonames/stream"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

const Url = "https://download.geonames.org/export/dump/"

type GeoNameFile string
type AltNameFile string
type LangCodeFile string
type TimeZoneFile string

//List of dump archives
const (
	Cities500      GeoNameFile  = "cities500.zip"
	Cities1000     GeoNameFile  = "cities1000.zip"
	Cities5000     GeoNameFile  = "cities5000.zip"
	Cities15000    GeoNameFile  = "cities15000.zip"
	AllCountries   GeoNameFile  = "allCountries.zip"
	NoCountry      GeoNameFile  = "no-country.zip"
	AlternateNames AltNameFile  = "alternateNamesV2.zip"
	LangCodes      LangCodeFile = "iso-languagecodes.txt"
	TimeZones      TimeZoneFile = "timeZones.txt"
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

func defaultFilename(archive string) string {
	return strings.Replace(filepath.Base(archive), ".zip", ".txt", 1)
}

func (p Parser) getArchive(archive string, handler func(f func(v interface{}) error) error, header ...string) error {
	r, err := p(archive)
	if err != nil {
		return err
	}

	return stream.StreamArchive(r, defaultFilename(archive), handler, header)
}

func (p Parser) getFile(file string, handler func(f func(v interface{}) error) error, header ...string) error {
	r, err := p(file)
	if err != nil {
		return err
	}

	return stream.StreamFile(r, handler, header)
}

func (p Parser) GetGeonames(archive GeoNameFile, handler func(*models.Geoname) error) error {
	headers, err := csvutil.Header(models.Geoname{}, "csv")
	if err != nil {
		return err
	}

	return p.getArchive(string(archive), func(parse func(v interface{}) error) error {
		model := &models.Geoname{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetAlternames(archive AltNameFile, handler func(*models.Altername) error) error {
	headers, err := csvutil.Header(models.Altername{}, "csv")
	if err != nil {
		return err
	}

	return p.getArchive(string(archive), func(parse func(v interface{}) error) error {
		model := &models.Altername{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetLanguages(handler func(language *models.Language) error) error {
	return p.getFile(string(LangCodes), func(parse func(v interface{}) error) error {
		model := &models.Language{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	})
}

func (p Parser) GetTimeZones(handler func(language *models.TimeZone) error) error {
	return p.getFile(string(TimeZones), func(parse func(v interface{}) error) error {
		model := &models.TimeZone{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	})
}
