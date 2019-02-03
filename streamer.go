package geonames

import (
	"encoding/csv"
	"fmt"
	"github.com/krolaw/zipstream"
	"io"
)

func Stream(r io.Reader, filename string, handler func(columns []string) error) error {
	archive := zipstream.NewReader(r)
	file, err := archive.Next()
	if err != nil && err != io.EOF {
		return err
	}

	for err != io.EOF {
		if file.Name == filename {
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

				if err := handler(columns); err != nil {
					return err
				}
			}

			return nil
		}

		file, err = archive.Next()
	}

	return fmt.Errorf("Archive doesnt contain the file %s", filename)
}
