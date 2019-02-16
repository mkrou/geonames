package geonames

import (
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/mkrou/geonames/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIntegrationParser_GetAlternames(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When alternames is parsed", func() {
			err := p.GetAlternames(AlternateNames, func(x *models.Altername) error {
				return v.ValidateStruct(x,
					v.Field(&x.Id, v.Required),
					v.Field(&x.Name, v.Required),
					v.Field(&x.GeonameId, v.Required),
				)
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
			err := p.GetAlternames("alternatenames/AD.zip", func(x *models.Altername) error {
				return v.ValidateStruct(x,
					v.Field(&x.Id, v.Required),
					v.Field(&x.Name, v.Required),
					v.Field(&x.GeonameId, v.Required),
				)
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
			err := p.GetGeonames(archive, func(x *models.Geoname) error {
				return v.ValidateStruct(x,
					v.Field(&x.Id, v.Required),
					v.Field(&x.Name, v.Required),
					v.Field(&x.ModificationDate, v.Required),
				)
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
			err := p.GetLanguages(func(x *models.Language) error {
				return v.ValidateStruct(x,
					v.Field(&x.Iso639_3, v.Required),
					v.Field(&x.Name, v.Required),
				)
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
			err := p.GetTimeZones(func(x *models.TimeZone) error {
				return v.ValidateStruct(x,
					v.Field(&x.Id, v.Required),
					v.Field(&x.CountryCode, v.Required),
				)
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
			err := p.GetCountries(func(x *models.Country) error {
				return v.ValidateStruct(x,
					v.Field(&x.Iso2Code, v.Required),
					v.Field(&x.Iso3Code, v.Required),
					v.Field(&x.IsoNumeric, v.Required),
					v.Field(&x.Name, v.Required),
					v.Field(&x.Continent, v.Required),
					v.Field(&x.GeonameID, v.Required),
				)
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
			err := p.GetFeatureCodes(FeatureCodeRu, func(x *models.FeatureCode) error {
				return v.ValidateStruct(x,
					v.Field(&x.Code, v.Required),
					v.Field(&x.Name, v.Required),
				)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetHierarchy(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When hierarchy is parsed", func() {
			err := p.GetHierarchy(func(x *models.Hierarchy) error {
				return v.ValidateStruct(x,
					v.Field(&x.Parent, v.Required),
					v.Field(&x.Child, v.Required),
				)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetShapes(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When shapes is parsed", func() {
			err := p.GetShapes(func(x *models.Shape) error {
				return v.ValidateStruct(x,
					v.Field(&x.GeonameId, v.Required),
					v.Field(&x.GeoJson, v.Required),
				)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetUserTags(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When user tags is parsed", func() {
			err := p.GetUserTags(func(x *models.UserTag) error {
				return v.Validate(x.Name, v.Required)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetAdminDivisions(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When admin divisions is parsed", func() {
			err := p.GetAdminDivisions(func(x *models.AdminDivision) error {
				return v.ValidateStruct(x,
					v.Field(&x.Code, v.Required),
					v.Field(&x.AsciiName, v.Required),
					v.Field(&x.GeonameId, v.Required),
				)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetAdminSubDivisions(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When admin sub divisions is parsed", func() {
			err := p.GetAdminSubdivisions(func(x *models.AdminSubdivision) error {
				return v.ValidateStruct(x,
					v.Field(&x.Code, v.Required),
					v.Field(&x.Name, v.Required),
					v.Field(&x.AsciiName, v.Required),
					v.Field(&x.GeonameId, v.Required),
				)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetAdminCode5(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When admin codes is parsed", func() {
			err := p.GetAdminCodes5(func(x *models.AdminCode5) error {
				return v.ValidateStruct(x,
					v.Field(&x.GeonameId, v.Required),
					v.Field(&x.AdminCode5, v.Required),
				)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetAlternameDeletes(t *testing.T) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey("When altername deletes is parsed", func() {
			err := p.GetAlternameDeletes(func(x *models.AlternameDeletes) error {
				return v.ValidateStruct(x,
					v.Field(&x.Id, v.Required),
					v.Field(&x.GeonameId, v.Required),
					v.Field(&x.Name, v.Required),
				)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
