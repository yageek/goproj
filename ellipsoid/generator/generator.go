package main

import (
	"encoding/json"
	datex "github.com/go-gis/datex/handlers/ellipsoid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const templateString = `// Auto-generated data - DO NOT EDIT
// Data retrieved from {{.DataURL}} on {{.Date}}
package ellipsoid

var datexCollection =[...]*EpsgEllipsoid{
{{range .Collection}}
 {{if eq .SemiMinorAxis 0.0}}
 	NewEpsgEllipsoidWithFlattening("{{.Name}}", {{.EPSG}}, {{printf "%f" .SemiMajorAxis}}, {{printf "%f" .InverseFlattening}} ),
 {{else}}
 	 NewEpsgEllipsoidWithSemiAxis( "{{.Name}}", {{.EPSG}}, {{printf "%f" .SemiMajorAxis}}, {{printf "%f" .SemiMinorAxis}} ) ,
 {{end}}
{{end}}
}`

const DatexURI = "http://go-datex.appsdeck.eu"

type EllipsoidCollection []datex.Ellipsoid

type EllipsoidGeneratorData struct {
	Collection EllipsoidCollection
	Date       time.Time
	DataURL    string
}

func main() {

	client := http.Client{}

	resp, err := client.Get(DatexURI + "/ellipsoid/all")

	if err != nil {
		log.Fatalln("Could not get ellipsoids definitions from the datex:", err)
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	log.Println("Data:", string(data))
	generatorData := &EllipsoidGeneratorData{Date: time.Now(), DataURL: DatexURI}

	var collection []datex.Ellipsoid
	err = json.Unmarshal(data, &collection)

	log.Println("Data:", collection)
	if err != nil {
		log.Fatalln("Could not unmarshal ellipsoids:", err)
	}

	file, err := os.Create("ellipsoid_index.go")
	if err != nil {
		log.Fatalln("Could not open file:", err)
	}

	defer file.Close()

	goTmpl, err := template.New("index").Parse(templateString)

	if err != nil {
		log.Fatalln("Invalid template:", err)
	}

	err = goTmpl.Execute(file, generatorData)

	if err != nil {
		log.Fatalln("Error on template:", err)
	}
}
