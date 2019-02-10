package geonames

import (
	"github.com/mkrou/geonames/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIntegrationParser_GetAlternames(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When alternames is parsed", func() {
			err := p.GetAlternames(AlternateNames, func(_ *models.Altername) error {
				return nil
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
func TestIntegrationParser_GetAlphabeticalAlternames(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When alternames is parsed", func() {
			err := p.GetAlternames("alternatenames/AD.zip", func(_ *models.Altername) error {
				return nil
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func cityTest(t *testing.T, archive GeoNameFile, msg string) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey(msg, func() {
			err := p.GetGeonames(archive, func(geoname *models.Geoname) error {
				return nil
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetGeonames500(t *testing.T) {
	cityTest(t, Cities500, "When all cities with a population > 500 are parsed")
}

func TestIntegrationParser_GetGeonames1000(t *testing.T) {
	cityTest(t, Cities1000, "When all cities with a population > 1000 are parsed")
}

func TestIntegrationParser_GetGeonames5000(t *testing.T) {
	cityTest(t, Cities5000, "When all cities with a population > 5000 are parsed")
}

func TestIntegrationParser_GetGeonames15000(t *testing.T) {
	cityTest(t, Cities15000, "When all cities with a population > 15000 are parsed")
}

func TestIntegrationParser_GetGeonamesAll(t *testing.T) {
	cityTest(t, AllCountries, "When all cities in all countries are parsed")
}

func TestIntegrationParser_GetGeonamesWithoutCountry(t *testing.T) {
	cityTest(t, NoCountry, "When all cities without countries are parsed")
}

func TestIntegrationParser_GetGeonamesAlphabetical(t *testing.T) {
	cityTest(t, "AD.zip", "When all cities that start with AD are parsed")
}

func TestIntegrationParser_GetLanguages(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When languages is parsed", func() {
			err := p.GetLanguages(func(_ *models.Language) error {
				return nil
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetTimeZones(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When time zones is parsed", func() {
			err := p.GetTimeZones(func(_ *models.TimeZone) error {
				return nil
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetCountries(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When countries is parsed", func() {
			err := p.GetCountries(func(_ *models.Country) error {
				return nil
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetFeatureCodes(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When feature codes is parsed", func() {
			err := p.GetFeatureCodes(FeatureCodeRu, func(_ *models.FeatureCode) error {
				return nil
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
