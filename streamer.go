package geonames

import (
	"encoding/csv"
	"github.com/krolaw/zipstream"
	"io"
)

func Stream(r io.Reader, handler func(filename string, columns []string) error) error {
	archive := zipstream.NewReader(r)
	file, err := archive.Next()
	if err != nil && err != io.EOF {
		return err
	}

	for err != io.EOF {
		r := csv.NewReader(archive)
		r.Comma = '\t'
		r.Comment = '#'
		r.LazyQuotes = true

		for {
			columns, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil && err != io.EOF {
				return err
			}

			if err := handler(file.Name, columns); err != nil {
				return err
			}
		}

		file, err = archive.Next()
	}

	return nil
}
