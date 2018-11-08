# Drip Injectable Content Weather Demo

This is a demonstration of how a backend API might look to integrate with Drip's Injectable Content product feature.

This API assumes a subscriber with a custom field called `zipcode`, which is then translated into a response containing the weather for that location.

## DarkSky API Key

You will need to obtain a DarkSky API key here: https://darksky.net/dev/account

This will be passed into the server via an environment variable called `DARK_SKY_API_KEY`.

## Running via Docker

TODO

## Running directly

Download the code with Go:

```
go get -u github.com/DripEmail/drip-injectable-weather
```

Assuming you have Go and the associated tools installed locally, run `DARK_SKY_API_KEY=abc123 go run main.go`.

## Building Docker container

You can build a Docker container with `docker build .`. Then run `docker run -e "DARK_SKY_API_KEY=abc123" -P 1234567` where `1234567` is the hash returned from the build.
