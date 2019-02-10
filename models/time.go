package models

import "time"

type Time struct {
	time.Time
}

func (t Time) MarshalCSV() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *Time) UnmarshalCSV(data []byte) error {
	date := string(data)
	if date == "" {
		return nil
	}

	var err error
	if t.Time, err = time.Parse("2006-01-02", date); err == nil {
		return nil
	}

	if t.Time, err = time.Parse("02 January 2006", date); err == nil {
		return nil
	}

	if t.Time, err = time.Parse("2006", date); err == nil {
		return nil
	}

	if t.Time, err = time.Parse("200601", date); err == nil {
		return nil
	}

	if t.Time, err = time.Parse("20060102", date); err == nil {
		return nil
	}

	if t.Time, err = time.Parse("02-01-2006", date); err == nil {
		return nil
	}

	return nil
}
