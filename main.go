package main

// https://medium.com/@elliotchance/working-with-json-in-go-d39c91d2464f
// https://medium.com/@nate510/dynamic-json-umarshalling-in-go-88095561d6a0
// https://ukiahsmith.com/blog/go-marshal-and-unmarshal-json-with-time-and-url-data/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const URI = "https://opendata.maryland.gov/resource/nigh-m2sg.json"

// {
// 	  "county": "Montgomery County",
// 	  "incident": "I-270 SOUTH AT EXIT 15B-A MD 118 GERMANTOWN RD (SB)",
// 	  "direction": "South",
// 	  "lanes": "1 of 2 Southbound shoulders closed.",
// 	  "link": "https://www.chart.maryland.gov/?id=105",
// 	  "created": "Created:  7/31/2019 9:33:16 AM by SOC.",
// 	  "lat": "{\"x\":-77.261034,\"y\":39.194238,\"spatialReference\":{\"wkid\":4326}}",
// 	  "long": "{\"x\":-77.261034,\"y\":39.194238,\"spatialReference\":{\"wkid\":4326}}",
// 	  "updated": "Fri, 06 Sep 2019 12:00:31 GMT"
// }

type Closure struct {
	County    string            `json:"county"`
	Incident  string            `json:"incident"`
	Direction string            `json:"direction"`
	Lanes     string            `json:"lanes"`
	Link      string            `json:"link"`
	Created   string            `json:"created"`
	Lat       CoordinateWrapper `json:"lat"`
	Long      CoordinateWrapper `json:"long"`
	Updated   TimeWrapper       `json:"updated"`
}

type Time struct {
	Updated time.Time
}

type TimeWrapper struct {
	Time
}

func (tw *TimeWrapper) UnmarshalJSON(b []byte) (err error) {
	// the 'updated' member is a string and ends up being double quoted
	// strip the quotes
	s := string(b)[1 : len(string(b))-1]
	t, err := time.Parse(time.RFC1123, s)
	if nil != err {
		return err
	}

	// set Closure.Updated to the parsed time.Time
	tw.Updated = t
	return nil
}

type Coordinate struct {
	X                float64 `json:"x"`
	Y                float64 `json:"y"`
	SpatialReference struct {
		Wkid int `json:"wkid"`
	} `json:"spatialReference"`
}

type CoordinateWrapper struct {
	Coordinate
}

func (cw *CoordinateWrapper) UnmarshalJSON(b []byte) (err error) {
	// The lat and long is invalid JSON.
	// "lat": "{\"x\":-77.261034,\"y\":39.194238,\"spatialReference\":{\"wkid\":4326}}"
	// here we unquote the string and unmarshal it
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(unquoted), &cw.Coordinate)
}

func main() {
	resp, err := http.Get(URI)
	if err != nil {
		fmt.Print("error downloading")
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("error reading")
		os.Exit(1)
	}

	var closures []Closure
	err = json.Unmarshal(body, &closures)
	if err != nil {
		// fmt.Print(fmt.Errorf("unmarshal body: %v", err))
		log.Fatal(err)
	}

	// for i, _ := range closures {

	// 	// Latitude
	// 	err := json.Unmarshal([]byte(closures[i].Lat), &closures[i].Latitude)
	// 	if err != nil {
	// 		fmt.Print(fmt.Errorf("unmarshal coord: %w", err))
	// 		os.Exit(1)
	// 	}

	// 	// Longitude
	// 	err = json.Unmarshal([]byte(closures[i].Long), &closures[i].Longitude)
	// 	if err != nil {
	// 		fmt.Print(fmt.Errorf("unmarshal coord: %w", err))
	// 		os.Exit(1)
	// 	}
	// }
	for i, _ := range closures {
		spew.Dump(closures[i])
	}
}
