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

const DatexURI = "http://go-datex.appsdeck.eu"

type EllipsoidCollection []datex.Ellipsoid

type EllipsoidGeneratorData struct {
	Collection *EllipsoidCollection
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

	generatorData := &EllipsoidGeneratorData{Collection: &EllipsoidCollection{}, Date: time.Now(), DataURL: DatexURI}

	err = json.Unmarshal(data, generatorData.Collection)

	if err != nil {
		log.Fatalln("Could not unmarshal ellipsoids:", err)
	}

	file, err := os.Create("../ellipsoid_index.go")
	if err != nil {
		log.Fatalln("Could not open file:", err)
	}

	defer file.Close()

	goTmpl, err := template.ParseFiles("ellipsoid.tmpl")

	if err != nil {
		log.Fatalln("Invalid template:", err)
	}

	err = goTmpl.Execute(file, generatorData)

	if err != nil {
		log.Fatalln("Error on template:", err)
	}
}
