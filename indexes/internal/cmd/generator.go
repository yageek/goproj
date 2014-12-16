package main

import (
	"encoding/json"
	"flag"
	"github.com/go-gis/datex/handlers/ellipsoid"
	"github.com/go-gis/datex/handlers/meridian"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

var templatePath string
var outputPath string
var dataType string

const DatexURI = "http://go-datex.appsdeck.eu"

type GeneratorData struct {
	Collection interface{}
	Date       time.Time
	DataURL    string
}

func init() {
	flag.StringVar(&templatePath, "tmpl", "", "The filepath to the template to use")
	flag.StringVar(&outputPath, "out", "", "The output filename")
	flag.StringVar(&dataType, "type", "", "Either Ellipsoid, Meridian, Unit or Datum")
}

func main() {

	flag.Parse()
	if templatePath == "" || outputPath == "" || dataType == "" {
		log.Fatalln("No template string, type or output filename are not provided")
	}
	client := http.Client{}
	log.Println("Generate indexes for type:", dataType)
	resp, err := client.Get(DatexURI + "/" + strings.ToLower(dataType) + "/all")

	if err != nil {
		log.Fatalln("Could not get ellipsoids definitions from the datex:", err)
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	generatorData := &GeneratorData{Date: time.Now(), DataURL: DatexURI}

	var collection interface{}
	switch dataType {
	default:
	case "Meridian":
		conteneur := make([]meridian.Meridian, 1)
		err = json.Unmarshal(data, &conteneur)
		collection = conteneur
	case "Ellipsoid":
		conteneur := make([]ellipsoid.Ellipsoid, 1)
		err = json.Unmarshal(data, &conteneur)
		collection = conteneur
	}
	if err != nil {
		log.Fatalln("Could not unmarshal ellipsoids:", err)
	}

	generatorData.Collection = collection

	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln("Could not open file:", err)
	}

	defer file.Close()

	goTmpl, err := template.ParseFiles(templatePath)

	if err != nil {
		log.Fatalln("Invalid template:", err)
	}

	err = goTmpl.Execute(file, generatorData)

	if err != nil {
		log.Fatalln("Error on template:", err)
	}
}
