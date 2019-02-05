package stream

import (
	"fmt"
	"github.com/krolaw/zipstream"
	"github.com/mkrou/geonames/csv"
	"io"
)

func StreamArchive(r io.Reader, filename string, handler func(columns []string) error, skipHeaders bool) error {
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
			return StreamFile(archive, handler, skipHeaders)
		}

		file, err = archive.Next()
	}

	return fmt.Errorf("Archive doesnt contain the file %s", filename)
}

func StreamFile(reader io.Reader, handler func(columns []string) error, skipHeaders bool) error {
	r := csv.NewReader(reader)
	r.Comma = '\t'
	r.Comment = '#'
	r.ReuseRecord = true
	if skipHeaders {
		if _, err := r.Read(); err != nil && err != io.EOF {
			return err
		}
	}

	for {
		columns, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if err := handler(columns); err != nil {
			return err
		}
	}

	return nil
}
