package geonames

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"github.com/mkrou/geonames/models"
	"github.com/mkrou/geonames/stream"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

const Url = "https://download.geonames.org/export/dump/"

type GeoNameFile string
type AltNameFile string
type FeatureCode string

//List of dump archives
const (
	Cities500                   GeoNameFile = "cities500.zip"
	Cities1000                  GeoNameFile = "cities1000.zip"
	Cities5000                  GeoNameFile = "cities5000.zip"
	Cities15000                 GeoNameFile = "cities15000.zip"
	AllCountries                GeoNameFile = "allCountries.zip"
	NoCountry                   GeoNameFile = "no-country.zip"
	AlternateNames              AltNameFile = "alternateNamesV2.zip"
	LangCodes                   string      = "iso-languagecodes.txt"
	TimeZones                   string      = "timeZones.txt"
	Countries                   string      = "countryInfo.txt"
	FeatureCodeBg               FeatureCode = "featureCodes_bg.txt"
	FeatureCodeEn               FeatureCode = "featureCodes_en.txt"
	FeatureCodeNb               FeatureCode = "featureCodes_nb.txt"
	FeatureCodeNn               FeatureCode = "featureCodes_nn.txt"
	FeatureCodeNo               FeatureCode = "featureCodes_no.txt"
	FeatureCodeRu               FeatureCode = "featureCodes_ru.txt"
	FeatureCodeSv               FeatureCode = "featureCodes_sv.txt"
	Hierarchy                   string      = "hierarchy.zip"
	Shapes                      string      = "shapes_all_low.zip"
	UserTags                    string      = "userTags.zip"
	AdminDivisions              string      = "admin1CodesASCII.txt"
	AdminSubDivisions           string      = "admin2Codes.txt"
	AdminCode5                  string      = "adminCode5.zip"
	AlternateNamesDeletes       string      = "alternateNamesDeletes-%s.txt"
	AlternateNamesModifications string      = "alternateNamesModifications-%s.txt"
	Deletes                     string      = "deletes-%s.txt"
)

func FormatLastDate(format string) string {
	loc, _ := time.LoadLocation("CET")
	t := time.Now().In(loc).AddDate(0, 0, -1)
	return fmt.Sprintf(format, t.Format("2006-01-02"))
}

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

func (p Parser) GetLanguages(handler func(*models.Language) error) error {
	return p.getFile(LangCodes, func(parse func(v interface{}) error) error {
		model := &models.Language{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	})
}

func (p Parser) GetTimeZones(handler func(*models.TimeZone) error) error {
	return p.getFile(TimeZones, func(parse func(v interface{}) error) error {
		model := &models.TimeZone{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	})
}

func (p Parser) GetCountries(handler func(*models.Country) error) error {
	headers, err := csvutil.Header(models.Country{}, "csv")
	if err != nil {
		return err
	}

	return p.getFile(Countries, func(parse func(v interface{}) error) error {
		model := &models.Country{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetFeatureCodes(file FeatureCode, handler func(*models.FeatureCode) error) error {
	headers, err := csvutil.Header(models.FeatureCode{}, "csv")
	if err != nil {
		return err
	}

	return p.getFile(string(file), func(parse func(v interface{}) error) error {
		model := &models.FeatureCode{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetHierarchy(handler func(*models.Hierarchy) error) error {
	headers, err := csvutil.Header(models.Hierarchy{}, "csv")
	if err != nil {
		return err
	}

	return p.getArchive(Hierarchy, func(parse func(v interface{}) error) error {
		model := &models.Hierarchy{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetShapes(handler func(*models.Shape) error) error {
	return p.getArchive(Shapes, func(parse func(v interface{}) error) error {
		model := &models.Shape{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	})
}

func (p Parser) GetUserTags(handler func(*models.UserTag) error) error {
	headers, err := csvutil.Header(models.UserTag{}, "csv")
	if err != nil {
		return err
	}

	return p.getArchive(UserTags, func(parse func(v interface{}) error) error {
		model := &models.UserTag{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetAdminDivisions(handler func(*models.AdminDivision) error) error {
	headers, err := csvutil.Header(models.AdminDivision{}, "csv")
	if err != nil {
		return err
	}

	return p.getFile(AdminDivisions, func(parse func(v interface{}) error) error {
		model := &models.AdminDivision{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetAdminSubdivisions(handler func(*models.AdminSubdivision) error) error {
	headers, err := csvutil.Header(models.AdminSubdivision{}, "csv")
	if err != nil {
		return err
	}

	return p.getFile(AdminSubDivisions, func(parse func(v interface{}) error) error {
		model := &models.AdminSubdivision{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetAdminCodes5(handler func(*models.AdminCode5) error) error {
	headers, err := csvutil.Header(models.AdminCode5{}, "csv")
	if err != nil {
		return err
	}

	return p.getArchive(AdminCode5, func(parse func(v interface{}) error) error {
		model := &models.AdminCode5{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetAlternameDeletes(handler func(*models.AlternameDelete) error) error {
	headers, err := csvutil.Header(models.AlternameDelete{}, "csv")
	if err != nil {
		return err
	}

	return p.getFile(FormatLastDate(AlternateNamesDeletes), func(parse func(v interface{}) error) error {
		model := &models.AlternameDelete{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetAlternameModifications(handler func(*models.AlternameModification) error) error {
	headers, err := csvutil.Header(models.AlternameModification{}, "csv")
	if err != nil {
		return err
	}

	return p.getFile(FormatLastDate(AlternateNamesModifications), func(parse func(v interface{}) error) error {
		model := &models.AlternameModification{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}

func (p Parser) GetDeletes(handler func(*models.GeonameDelete) error) error {
	headers, err := csvutil.Header(models.GeonameDelete{}, "csv")
	if err != nil {
		return err
	}

	return p.getFile(FormatLastDate(Deletes), func(parse func(v interface{}) error) error {
		model := &models.GeonameDelete{}
		if err := parse(model); err != nil {
			return err
		}

		return handler(model)
	}, headers...)
}
