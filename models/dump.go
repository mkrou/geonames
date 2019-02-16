package models

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type (
	DumpFile        string
	GeoNameFile     DumpFile
	AltNameFile     DumpFile
	FeatureCodeFile DumpFile
)

func (d DumpFile) TextFilename() string {
	if d.IsArchive() {
		return strings.Replace(filepath.Base(d.String()), ".zip", ".txt", 1)
	} else {
		return d.String()
	}
}

func (d DumpFile) IsArchive() bool {
	return filepath.Ext(d.String()) == ".zip"
}

func (d DumpFile) String() string {
	return string(d)
}

func (d DumpFile) WithLastDate() DumpFile {
	loc, _ := time.LoadLocation("CET")
	t := time.Now().In(loc).AddDate(0, 0, -1)
	return DumpFile(fmt.Sprintf(d.String(), t.Format("2006-01-02")))
}
