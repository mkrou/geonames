package geonames

import (
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

func (p Parser) getArchive(archive string, handler func(columns []string) error, skipHeaders bool, fieldsPerRecord int) error {
	r, err := p(archive)
	if err != nil {
		return err
	}

	return stream.StreamArchive(r, defaultFilename(archive), func(columns []string) error {
		return handler(columns)
	}, skipHeaders, fieldsPerRecord)
}

func (p Parser) getFile(archive string, handler func(columns []string) error, skipHeaders bool, fieldsPerRecord int) error {
	r, err := p(archive)
	if err != nil {
		return err
	}

	return stream.StreamFile(r, func(columns []string) error {
		return handler(columns)
	}, skipHeaders, fieldsPerRecord)
}

func (p Parser) GetGeonames(archive GeoNameFile, handler func(*models.Geoname) error) error {
	return p.getArchive(string(archive), func(columns []string) error {
		model, err := models.ParseGeoname(columns)
		if err != nil {
			return err
		}

		return handler(model)
	}, false, models.GeonameFields)
}

func (p Parser) GetAlternames(archive AltNameFile, handler func(*models.Altername) error) error {
	return p.getArchive(string(archive), func(columns []string) error {
		model, err := models.ParseAltername(columns)
		if err != nil {
			return err
		}

		return handler(model)
	}, false, models.AlternameFields)
}

func (p Parser) GetLanguages(handler func(language *models.Language) error) error {
	return p.getFile(string(LangCodes), func(columns []string) error {
		model, err := models.ParseLanguage(columns)
		if err != nil {
			return err
		}

		return handler(model)
	}, true, models.LanguageFields)
}

func (p Parser) GetTimeZones(handler func(language *models.TimeZone) error) error {
	return p.getFile(string(TimeZones), func(columns []string) error {
		model, err := models.ParseTimeZone(columns)
		if err != nil {
			return err
		}

		return handler(model)
	}, true, models.TimeZoneFields)
}
