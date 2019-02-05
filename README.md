# Geonames

#### Golang parsing library for the [geonames.org](http://www.geonames.org) database [dump](http://download.geonames.org/export/dump/).

![](static/example.gif)

## Features
- Parse data directly without downloading and unzipping
- Read line by line with low memory consumption

## Implemented data

|status|archive|comment|
|---|---|---|
|✅|xx.zip|Parser.GetGeonames; See [readme](#parsing-alphabetical-list-of-archives)|
|🚫|admin1CodesASCII.txt||
|🚫|admin2Codes.txt||
|🚫|adminCode5.zip||
|✅|allCountries.zip|Parser.GetGeonames|
|🚫|alternateNames.zip|depricated, use alternateNamesV2.zip instead|
|🚫|alternateNamesDeletes-xxxx-xx-xx.txt||
|🚫|alternateNamesModifications-xxxx-xx-xx.txt||
|✅|alternateNamesV2.zip|Parser.GetAlternames|
|✅|alternatenames/xx.zip|Parser.GetAlternames; See [readme](#parsing-alphabetical-list-of-archives)|
|✅|cities1000.zip|Parser.GetGeonames|
|✅|cities15000.zip|Parser.GetGeonames|
|✅|cities500.zip|Parser.GetGeonames|
|✅|cities5000.zip|Parser.GetGeonames|
|🚫|countryInfo.txt|Parser.GetGeonames|
|🚫|deletes-xxxx-xx-xx.txt||
|🚫|featureCodes_bg.txt||
|🚫|featureCodes_en.txt||
|🚫|featureCodes_nb.txt||
|🚫|featureCodes_nn.txt||
|🚫|featureCodes_no.txt||
|🚫|featureCodes_ru.txt||
|🚫|featureCodes_sv.txt||
|🚫|hierarchy.zip||
|✅|iso-languagecodes.txt|Parser.GetLanguages|
|🚫|modifications-xxxx-xx-xx.txt||
|✅|no-country.zip|Parser.GetGeonames|
|🚫|shapes_all_low.zip||
|🚫|shapes_simplified_low.json.zip||
|🚫|timeZones.txt||
|🚫|userTags.zip||

## Installation

```bash 
$ go get github.com/mkrou/geonames
```

## Quick start

#### Parsing cities
```go

package main

import (
    "fmt"
    "github.com/mkrou/geonames"
    "github.com/mkrou/geonames/models"
    "log"
)

func main() {
    p := geonames.NewParser()
    
    //print all cities with a population greater then 5000
    err := p.GetGeonames(geonames.Cities5000, func(geoname *models.Geoname) error {
    fmt.Println(geoname.Name)
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
}

```
#### Parsing alternames

```go
package main

import (
    "fmt"
    "github.com/mkrou/geonames"
    "github.com/mkrou/geonames/models"
    "log"
)
func main() {
    p := geonames.NewParser()
    
    err := p.GetAlternames(geonames.AlternateNames, func(geoname *models.Altername) error {
        fmt.Println(geoname.Name)
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

#### Parsing alphabetical list of archives

```go
package main

import (
    "fmt"
    "github.com/mkrou/geonames"
    "github.com/mkrou/geonames/models"
    "log"
)
func main() {
    p := geonames.NewParser()
    
    err := p.GetGeonames("AD.zip", func(geoname *models.Geoname) error {
        fmt.Println(geoname.Name)
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
    
    err = p.GetAlternames("alternames/AD.zip", func(geoname *models.Altername) error {
        fmt.Println(geoname.Name)
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
}
```
