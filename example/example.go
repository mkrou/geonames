package main

import (
	"fmt"
	"geonames"
	"geonames/models"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"os"
	"runtime"
	"time"
)

func main() {
	p := geonames.NewParser()

	w := wow.New(os.Stdout, spin.Spinner{Frames: []string{"⚙️"}}, "  Parsing all cities with a population > 500...")
	w.Persist()
	w.Text("").Spinner(spin.Get(spin.Earth)).Start()

	count := 0
	since := time.Now()
	err := p.GetGeonames(geonames.Cities500, func(geoname *models.Geoname) error {
		count++
		w.Text(fmt.Sprintf("%d: %s", count, geoname.Name))

		return nil
	})
	if err != nil {
		w.PersistWith(spin.Spinner{Frames: []string{"🔥"}}, fmt.Sprintf(" Error: %s", err.Error()))
		return
	}
	duration := time.Since(since)

	w.PersistWith(spin.Spinner{Frames: []string{"✅"}}, fmt.Sprintf(" Done!"))
	w.PersistWith(spin.Spinner{Frames: []string{"⛩"}}, fmt.Sprintf("  Cities: %d", count))
	w.PersistWith(spin.Spinner{Frames: []string{"⏱"}}, fmt.Sprintf("  Duration: %d sec", duration/time.Second))
	w.PersistWith(spin.Spinner{Frames: []string{"💾️‍"}}, fmt.Sprintf(" Memory: %d Mb", GetMemUsage()))
}

func GetMemUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.TotalAlloc / 1024 / 1024
}
