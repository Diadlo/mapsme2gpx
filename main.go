package main

import (
	"encoding/xml"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type MapsMe struct {
	Document struct {
		Placemark struct {
			Name       string `xml:"name"`
			LineString struct {
				Coordinates string `xml:"coordinates"`
			} `xml:"LineString"`
			ExtendedData struct {
				TrainInfo struct {
					Point []struct {
						Timestamp string `xml:"timestamp"`
					} `xml:"point"`
				} `xml:"trainInfo"`
			} `xml:"ExtendedData"`
		} `xml:"Placemark"`
	} `xml:"Document"`
}

type Point struct {
	Lat        float64
	Lon        float64
	Elevation  float64
	IsoTimeStr string
}

type RenderInfo struct {
	Name        string
	StartTime   string
	TrackPoints []Point
}

func main() {
	kmlStr, err := os.ReadFile("/tmp/input.kml")
	if err != nil {
		panic(err)
	}

	var mapsMe MapsMe
	err = xml.Unmarshal([]byte(kmlStr), &mapsMe)
	if err != nil {
		panic(err)
	}

	coordWords := strings.Split(mapsMe.Document.Placemark.LineString.Coordinates, " ")
	points := make([]Point, len(coordWords))
	for i, coordWord := range coordWords {
		coord := strings.Split(coordWord, ",")

		points[i].Lon, err = strconv.ParseFloat(coord[0], 64)
		if err != nil {
			panic(err)
		}

		points[i].Lat, err = strconv.ParseFloat(coord[1], 64)
		if err != nil {
			panic(err)
		}

		points[i].Elevation, err = strconv.ParseFloat(coord[2], 64)
		if err != nil {
			panic(err)
		}
	}

	timestamps := mapsMe.Document.Placemark.ExtendedData.TrainInfo.Point
	if len(coordWords) != len(timestamps) {
		panic("Number of coordinates and timestamps doesn't match")
	}

	for i, timestamp := range timestamps {
		points[i].IsoTimeStr = timestamp.Timestamp
	}

	tmpl, err := template.New("gpx template").Parse(GpxTemplate)
	if err != nil {
		panic(err)
	}

	renderInfo := RenderInfo{
		Name:        mapsMe.Document.Placemark.Name,
		StartTime:   mapsMe.Document.Placemark.Name,
		TrackPoints: points,
	}

	err = tmpl.Execute(os.Stdout, renderInfo)
	if err != nil {
		panic(err)
	}
}
