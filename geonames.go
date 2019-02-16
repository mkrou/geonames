package geonames

import (
	"errors"
	"fmt"
	"github.com/jszwec/csvutil"
	"github.com/mkrou/geonames/models"
	"github.com/mkrou/geonames/stream"
	"io"
	"net/http"
)

const Url = "https://download.geonames.org/export/dump/"

//List of dump archives
const (
	Cities500                   models.GeoNameFile     = "cities500.zip"
	Cities1000                  models.GeoNameFile     = "cities1000.zip"
	Cities5000                  models.GeoNameFile     = "cities5000.zip"
	Cities15000                 models.GeoNameFile     = "cities15000.zip"
	AllCountries                models.GeoNameFile     = "allCountries.zip"
	NoCountry                   models.GeoNameFile     = "no-country.zip"
	AlternateNames              models.AltNameFile     = "alternateNamesV2.zip"
	LangCodes                   models.DumpFile        = "iso-languagecodes.txt"
	TimeZones                   models.DumpFile        = "timeZones.txt"
	Countries                   models.DumpFile        = "countryInfo.txt"
	FeatureCodeBg               models.FeatureCodeFile = "featureCodes_bg.txt"
	FeatureCodeEn               models.FeatureCodeFile = "featureCodes_en.txt"
	FeatureCodeNb               models.FeatureCodeFile = "featureCodes_nb.txt"
	FeatureCodeNn               models.FeatureCodeFile = "featureCodes_nn.txt"
	FeatureCodeNo               models.FeatureCodeFile = "featureCodes_no.txt"
	FeatureCodeRu               models.FeatureCodeFile = "featureCodes_ru.txt"
	FeatureCodeSv               models.FeatureCodeFile = "featureCodes_sv.txt"
	Hierarchy                   models.DumpFile        = "hierarchy.zip"
	Shapes                      models.DumpFile        = "shapes_all_low.zip"
	UserTags                    models.DumpFile        = "userTags.zip"
	AdminDivisions              models.DumpFile        = "admin1CodesASCII.txt"
	AdminSubDivisions           models.DumpFile        = "admin2Codes.txt"
	AdminCode5                  models.DumpFile        = "adminCode5.zip"
	AlternateNamesDeletes       models.DumpFile        = "alternateNamesDeletes-%s.txt"
	AlternateNamesModifications models.DumpFile        = "alternateNamesModifications-%s.txt"
	Deletes                     models.DumpFile        = "deletes-%s.txt"
	Modifications               models.DumpFile        = "modifications-%s.txt"
)

type Parser func(file string) (io.ReadCloser, error)

func NewParser() Parser {
	return Parser(func(file string) (io.ReadCloser, error) {
		url := Url + file
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		switch resp.StatusCode {
		case 200:
			return resp.Body, nil
		case 404:
			return nil, errors.New(fmt.Sprintf("Page %s does not exist", url))
		default:
			return nil, errors.New(fmt.Sprintf("Page %s returned unexpected code %d", url, resp.StatusCode))
		}
	})
}

func (p Parser) handle(dump models.DumpFile, isHeadersEmpty bool, handler interface{}) error {
	var err error
	var headers = []string{}

	model, err := getArgument(handler)
	if err != nil {
		return err
	}

	if isHeadersEmpty {
		headers, err = csvutil.Header(model, "csv")
		if err != nil {
			return err
		}
	}

	r, err := p(dump.String())
	if err != nil {
		return err
	}
	f := func(parse func(v interface{}) error) error {
		return fillArgument(handler, parse)
	}

	if dump.IsArchive() {
		err = stream.StreamArchive(r, dump.TextFilename(), f, headers)
	} else {
		err = stream.StreamFile(r, f, headers)
	}

	return err
}

func (p Parser) GetGeonames(archive models.GeoNameFile, handler func(*models.Geoname) error) error {
	return p.handle(models.DumpFile(archive), true, handler)
}

func (p Parser) GetAlternateNames(archive models.AltNameFile, handler func(*models.AlternateName) error) error {
	return p.handle(models.DumpFile(archive), true, handler)
}

func (p Parser) GetLanguages(handler func(*models.Language) error) error {
	return p.handle(LangCodes, false, handler)
}

func (p Parser) GetTimeZones(handler func(*models.TimeZone) error) error {
	return p.handle(TimeZones, false, handler)
}

func (p Parser) GetCountries(handler func(*models.Country) error) error {
	return p.handle(Countries, true, handler)
}

func (p Parser) GetFeatureCodes(file models.FeatureCodeFile, handler func(*models.FeatureCode) error) error {
	return p.handle(models.DumpFile(file), true, handler)
}

func (p Parser) GetHierarchy(handler func(*models.Hierarchy) error) error {
	return p.handle(Hierarchy, true, handler)
}

func (p Parser) GetShapes(handler func(*models.Shape) error) error {
	return p.handle(Shapes, false, handler)
}

func (p Parser) GetUserTags(handler func(*models.UserTag) error) error {
	return p.handle(UserTags, true, handler)
}

func (p Parser) GetAdminDivisions(handler func(*models.AdminDivision) error) error {
	return p.handle(AdminDivisions, true, handler)
}

func (p Parser) GetAdminSubdivisions(handler func(*models.AdminSubdivision) error) error {
	return p.handle(AdminSubDivisions, true, handler)
}

func (p Parser) GetAdminCodes5(handler func(*models.AdminCode5) error) error {
	return p.handle(AdminCode5, true, handler)
}

func (p Parser) GetAlternateNameDeletes(handler func(*models.AlternateNameDelete) error) error {
	return p.handle(AlternateNamesDeletes.WithLastDate(), true, handler)
}

func (p Parser) GetAlternateNameModifications(handler func(*models.AlternateNameModification) error) error {
	return p.handle(AlternateNamesModifications.WithLastDate(), true, handler)
}

func (p Parser) GetDeletes(handler func(*models.GeonameDelete) error) error {
	return p.handle(Deletes.WithLastDate(), true, handler)
}

func (p Parser) GetModifications(handler func(*models.Geoname) error) error {
	return p.handle(Modifications.WithLastDate(), true, handler)
}
