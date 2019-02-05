package models

//ISO 639-3
//ISO 639-2
//ISO 639-1
//Language Name

const LanguageFields = 4

type Language struct {
	Iso639_1 string
	Iso639_2 string
	Iso639_3 string
	Name     string
}

func ParseLanguage(parts []string) (*Language, error) {
	return &Language{
		Iso639_1: parts[2],
		Iso639_2: parts[1],
		Iso639_3: parts[0],
		Name:     parts[3],
	}, nil
}
