package geonames

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/mkrou/geonames/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func validate(x interface{}) error {
	_, err := govalidator.ValidateStruct(x)
	return err
}

func integration(t *testing.T, kind string, f func(p Parser) error) {
	Convey("Given a default parser", t, func() {
		p := NewParser()

		Convey(fmt.Sprintf("When %s are parsed", kind), func() {
			err := f(p)

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestIntegrationParser_GetAlternames(t *testing.T) {
	integration(t, "alternames", func(p Parser) error {
		return p.GetAlternateNames(AlternateNames, func(x *models.AlternateName) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetAlphabeticalAlternames(t *testing.T) {
	integration(t, "alternames", func(p Parser) error {
		return p.GetAlternateNames("alternatenames/AD.zip", func(x *models.AlternateName) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetGeonames500(t *testing.T) {
	integration(t, "all cities with a population > 500", func(p Parser) error {
		return p.GetGeonames(Cities500, func(x *models.Geoname) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetGeonames1000(t *testing.T) {
	integration(t, "all cities with a population > 1000", func(p Parser) error {
		return p.GetGeonames(Cities1000, func(x *models.Geoname) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetGeonames5000(t *testing.T) {
	integration(t, "all cities with a population > 5000", func(p Parser) error {
		return p.GetGeonames(Cities5000, func(x *models.Geoname) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetGeonames15000(t *testing.T) {
	integration(t, "all cities with a population > 15000", func(p Parser) error {
		return p.GetGeonames(Cities15000, func(x *models.Geoname) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetGeonamesAll(t *testing.T) {
	integration(t, "all cities in all countries", func(p Parser) error {
		return p.GetGeonames(AllCountries, func(x *models.Geoname) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetGeonamesWithoutCountry(t *testing.T) {
	integration(t, "all cities without countries", func(p Parser) error {
		return p.GetGeonames(NoCountry, func(x *models.Geoname) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetGeonamesAlphabetical(t *testing.T) {
	integration(t, "all cities that start with AD", func(p Parser) error {
		return p.GetGeonames("AD.zip", func(x *models.Geoname) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetLanguages(t *testing.T) {
	integration(t, "languages", func(p Parser) error {
		return p.GetLanguages(func(x *models.Language) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetTimeZones(t *testing.T) {
	integration(t, "time zones", func(p Parser) error {
		return p.GetTimeZones(func(x *models.TimeZone) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetCountries(t *testing.T) {
	integration(t, "countries", func(p Parser) error {
		return p.GetCountries(func(x *models.Country) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetFeatureCodes(t *testing.T) {
	integration(t, "abc", func(p Parser) error {
		return p.GetFeatureCodes(FeatureCodeRu, func(x *models.FeatureCode) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetHierarchy(t *testing.T) {
	integration(t, "hierarchies", func(p Parser) error {
		return p.GetHierarchy(func(x *models.Hierarchy) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetShapes(t *testing.T) {
	integration(t, "shapes", func(p Parser) error {
		return p.GetShapes(func(x *models.Shape) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetUserTags(t *testing.T) {
	integration(t, "user tags", func(p Parser) error {
		return p.GetUserTags(func(x *models.UserTag) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetAdminDivisions(t *testing.T) {
	integration(t, "admin divisions", func(p Parser) error {
		return p.GetAdminDivisions(func(x *models.AdminDivision) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetAdminSubDivisions(t *testing.T) {
	integration(t, "admin sub divisions", func(p Parser) error {
		return p.GetAdminSubdivisions(func(x *models.AdminSubdivision) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetAdminCode5(t *testing.T) {
	integration(t, "new admin codes", func(p Parser) error {
		return p.GetAdminCodes5(func(x *models.AdminCode5) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetAlternameDeletes(t *testing.T) {
	integration(t, "altername deletes", func(p Parser) error {
		return p.GetAlternateNameDeletes(func(x *models.AlternateNameDelete) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetAlternameModifications(t *testing.T) {
	integration(t, "altername modifications", func(p Parser) error {
		return p.GetAlternateNameModifications(func(x *models.AlternateNameModification) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetDeletes(t *testing.T) {
	integration(t, "deletes", func(p Parser) error {
		return p.GetDeletes(func(x *models.GeonameDelete) error {
			return validate(x)
		})
	})
}

func TestIntegrationParser_GetModifications(t *testing.T) {
	integration(t, "modifications", func(p Parser) error {
		return p.GetModifications(func(x *models.Geoname) error {
			return validate(x)
		})
	})
}
