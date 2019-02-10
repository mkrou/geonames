package stream

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"github.com/krolaw/zipstream"
	"github.com/mkrou/geonames/csv"
	"io"
)

func StreamArchive(r io.Reader, filename string, handler func(f func(v interface{}) error) error, missedHeaders []string) error {
	archive := zipstream.NewReader(r)
	file, err := archive.Next()
	if err != nil && err != io.EOF {
		return err
	}

	for err != io.EOF {
		if err != nil {
			return err
		}

		if file.Name == filename {
			return StreamFile(archive, handler, missedHeaders)
		}

		file, err = archive.Next()
	}

	return fmt.Errorf("Archive doesnt contain the file %s", filename)
}

func StreamFile(reader io.Reader, handler func(f func(v interface{}) error) error, missedHeaders []string) error {
	r := csv.NewReader(reader)
	r.Comma = '\t'
	r.Comment = '#'
	r.ReuseRecord = true

	dec, err := csvutil.NewDecoder(r, missedHeaders...)
	if err != nil {
		return err
	}

	for {
		err := handler(dec.Decode)
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

	}

	return nil
}
