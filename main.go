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

func getForecast(location Location) (darksky.ForecastResponse, error) {
	client := darksky.New(os.Getenv("DARK_SKY_API_KEY"))
	request := darksky.ForecastRequest{}
	request.Latitude = location.Latitude
	request.Longitude = location.Longitude
	request.Options = darksky.ForecastRequestOptions{Exclude: "hourly,minutely"}
	forecast, err := client.Forecast(request)
	if err != nil {
		return darksky.ForecastResponse{}, err
	}

	return forecast, nil
}

func getLocation(zip ZipCode) (Location, error) {
	lat, lng, err := geocoder.Geocode(string(zip))
	if err != nil {
		return Location{}, err
	}

	return Location{darksky.Measurement(lat), darksky.Measurement(lng)}, nil
}

func getForecastResponse(zipcode ZipCode) ([]byte, error) {
	loc, err := getLocation(zipcode)
	if err != nil {
		return nil, err
	}
	forecast, err := getForecast(loc)
	if err != nil {
		return nil, err
	}
	writableForecast, err := json.Marshal(forecast.Currently)
	if err != nil {
		return nil, err
	}
	return writableForecast, nil
}

func zipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	if len(params["subscriber.zipcode"]) > 0 {
		zipcode := ZipCode(params["subscriber.zipcode"][0])
		writableForecast, err := getForecastResponse(zipcode)
		if err != nil {
			log.Printf("Error when getting forecat: %s", err.Error())
			http.Error(w, "Internal server error", 500)
			return
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
