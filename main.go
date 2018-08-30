package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jasonwinn/geocoder"
	"github.com/shawntoffel/darksky"
)

// Location contains coordinates
type Location struct {
	Latitude  darksky.Measurement
	Longitude darksky.Measurement
}

// ResponseError defines the json response for an error.
type ResponseError struct {
	error string
}

func getForecast(location Location) darksky.ForecastResponse {
	client := darksky.New(os.Getenv("DARK_SKY_API_KEY"))
	request := darksky.ForecastRequest{}
	request.Latitude = location.Latitude
	request.Longitude = location.Longitude
	request.Options = darksky.ForecastRequestOptions{Exclude: "hourly,minutely"}
	forecast, err := client.Forecast(request)
	if err != nil {
		// TODO: better error handling.
		// Return 500?
		panic(err)
	}

	return forecast
}

func getLocation(zip ZipCode) Location {
	lat, lng, err := geocoder.Geocode(string(zip))
	if err != nil {
		// TODO: This may happen if the zipcode is not in the right format.
		panic(err)
	}

	return Location{darksky.Measurement(lat), darksky.Measurement(lng)}
}

func zipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	var zipcode ZipCode
	if len(params["subscriber.zipcode"]) > 0 {
		zipcode = ZipCode(params["subscriber.zipcode"][0])
		loc := getLocation(zipcode)
		forecast := getForecast(loc)
		writableForecast, err := json.Marshal(forecast.Currently)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(writableForecast)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		jData, err := json.Marshal(ResponseError{"Needs subscriber.zipcode param"})
		if err != nil {
			log.Fatal(err)
		}
		w.Write(jData)
	}
}

func main() {
	http.HandleFunc("/api", zipHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
