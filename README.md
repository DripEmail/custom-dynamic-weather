# Drip Custom Dynamic Content Weather Demo

[![Build Status](https://travis-ci.org/DripEmail/custom-dynamic-weather.svg?branch=master)](https://travis-ci.org/DripEmail/custom-dynamic-weather)

This is a demonstration of how a backend API might look to integrate with Drip's [Custom Dynamic Content](http://developer.drip.com/#background) product feature.

This API assumes a subscriber with a custom field called `zipcode`, which is then translated into a response containing the weather for that location.

## Dark Sky<sup>&copy;</sup> API Key

You will need to obtain a Dark Sky<sup>&copy;</sup> API key here: https://darksky.net/dev/account

This will be passed into the server via an environment variable called `DARK_SKY_API_KEY`.

## Running via Docker

```bash
docker pull getdrip/custom-dynamic-weather
docker run -e "DARK_SKY_API_KEY=abc123" -p 8080:8080 getdrip/custom-dynamic-weather
```

## Running directly

Download the code with Go:

```bash
go get -u github.com/DripEmail/custom-dynamic-weather
```

Assuming you have Go and the associated tools installed locally, run `DARK_SKY_API_KEY=abc123 go run main.go`.

## Building Docker container

You can build a Docker container with `docker build .`. Then run `docker run -e "DARK_SKY_API_KEY=abc123" -p 8080:8080 1234567` where `1234567` is the hash returned from the build.

## Usage

Via cURL:

```bash
curl -v "http://localhost:8080/api?subscriber%5Bzipcode%5D=55401"
```

Example response:

```json
{
  "apparentTemperature": 51.57,
  "cloudCover": 0.48,
  "dewPoint": 43.6,
  "humidity": 0.74,
  "icon": "partly-cloudy-night",
  "ozone": 295.5,
  "pressure": 1020.55,
  "summary": "Partly Cloudy",
  "temperature": 51.57,
  "time": 1541697515,
  "visibility": 8.52,
  "windBearing": 67,
  "windGust": 3.78,
  "windSpeed": 3.58
}
```
