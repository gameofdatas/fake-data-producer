package models

import (
	"github.com/brianvoe/gofakeit/v6"
)

type Address struct {
	Id        string  `json:"id" xml:"id" fake:"{number:1001,1100}"`
	Street    string  `json:"street" xml:"street" fake:"{street}"`
	City      string  `json:"city" xml:"city" fake:"{city}"`
	State     string  `json:"state" xml:"state" fake:"{state}"`
	Zip       string  `json:"zip" xml:"zip" fake:"{zip}"`
	Country   string  `json:"country" xml:"country" fake:"{country}"`
	Latitude  float64 `json:"latitude" xml:"latitude" fake:"{latitude}"`
	Longitude float64 `json:"longitude" xml:"longitude" fake:"{longitude}"`
}

func createAddressData() Address {
	var a Address
	gofakeit.Struct(&a)
	return a
}
