package geolocationclient

import (
	"errors"
	"fmt"

	config "github.com/alexnorgaard/eventsapp"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/google"
)

func GetGeolocation(address string) (*geo.Location, error) {
	fmt.Printf("Getting geolocation for address: %v\n", address)
	api_key := config.GetConf().Google_geocoding_api.Api_key
	geocoder := google.Geocoder(api_key)
	location, err := geocoder.Geocode(address)
	if err != nil {
		return nil, err
	}
	if location == nil {
		return nil, errors.New("No location found")
	}
	fmt.Printf("Location is: %v\n", location)
	return location, nil
}

func GetAddress(lat float64, lng float64) (*geo.Address, error) {
	api_key := config.GetConf().Google_geocoding_api.Api_key
	geocoder := google.Geocoder(api_key)
	address, err := geocoder.ReverseGeocode(lat, lng)
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, errors.New("No address found")
	}
	return address, nil
}
