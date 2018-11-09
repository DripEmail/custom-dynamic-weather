package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/DripEmail/custom-dynamic-weather/zipcode"
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

// NewLocation instantiates a Location given a lat and long
func NewLocation(lat, lng float64) Location {
	return Location{
		Latitude:  darksky.Measurement(lat),
		Longitude: darksky.Measurement(lng),
	}
}

func getForecast(location Location) (darksky.ForecastResponse, error) {
	client := darksky.New(os.Getenv("DARK_SKY_API_KEY"))
	request := darksky.ForecastRequest{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Options:   darksky.ForecastRequestOptions{Exclude: "hourly,minutely"},
	}
	forecast, err := client.Forecast(request)
	if err != nil {
		return darksky.ForecastResponse{}, err
	}

	return forecast, nil
}

func getLocation(zip zipcode.ZipCode) (Location, error) {
	lat, lng, err := geocoder.Geocode(string(zip))
	if err != nil {
		return Location{}, err
	}

	return NewLocation(lat, lng), nil
}

func getForecastResponse(zip zipcode.ZipCode) ([]byte, error) {
	loc, err := getLocation(zip)
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
	if len(params["subscriber[zipcode]"]) > 0 {
		zip := zipcode.ZipCode(params["subscriber[zipcode]"][0])
		writableForecast, err := getForecastResponse(zip)
		if err != nil {
			log.Printf("Error when getting forecast: %s", err.Error())
			http.Error(w, "Internal server error", 500)
			return
		}
		if _, err := w.Write(writableForecast); err != nil {
			log.Printf("Error when sending successful response: %s", err.Error())
			http.Error(w, "Internal server error", 500)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		jData, err := json.Marshal(ResponseError{"Needs subscriber.zipcode param"})
		if err != nil {
			log.Fatal(err)
		}
		if _, err := w.Write(jData); err != nil {
			log.Printf("Error when sending bad request response: %s", err.Error())
			http.Error(w, "Internal server error", 500)
		}
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Yup, it's working...")); err != nil {
		log.Printf("Error when sending bad request response: %s", err.Error())
		http.Error(w, "Internal server error", 500)
	}
}

func main() {
	if os.Getenv("DARK_SKY_API_KEY") == "" {
		log.Fatal("DARK_SKY_API_KEY environment variable is required. Get one at https://darksky.net/dev/account")
	}
	http.HandleFunc("/api", zipHandler)
	http.HandleFunc("/status", statusHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
