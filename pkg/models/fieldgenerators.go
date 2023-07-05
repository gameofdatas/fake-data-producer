package models

import (
	"github.com/brianvoe/gofakeit/v6"
)

type FieldGenerator func() interface{}

var fieldGenerators = map[string]FieldGenerator{
	"Address":          func() interface{} { return gofakeit.Address().Address },
	"Age":              func() interface{} { return gofakeit.Number },
	"City":             func() interface{} { return gofakeit.City() },
	"CompanyName":      func() interface{} { return gofakeit.Company() },
	"ContactNumber":    func() interface{} { return gofakeit.Phone() },
	"Country":          func() interface{} { return gofakeit.Country() },
	"CreditCardNumber": func() interface{} { return gofakeit.CreditCardNumber(&gofakeit.CreditCardOptions{}) },
	"CurrencyCode":     func() interface{} { return gofakeit.CurrencyShort() },
	"CustomerID":       func() interface{} { return gofakeit.UUID() },
	"DateOfBirth":      func() interface{} { return gofakeit.Date() },
	"Email":            func() interface{} { return gofakeit.Email() },
	"FirstName":        func() interface{} { return gofakeit.FirstName() },
	"Gender":           func() interface{} { return gofakeit.Gender() },
	"Height":           func() interface{} { return gofakeit.Number },
	"IPAddress":        func() interface{} { return gofakeit.IPv4Address() },
	"JobTitle":         func() interface{} { return gofakeit.JobTitle() },
	"LastName":         func() interface{} { return gofakeit.LastName() },
	"Latitude":         func() interface{} { return gofakeit.Latitude() },
	"Longitude":        func() interface{} { return gofakeit.Longitude() },
	"MiddleName":       func() interface{} { return gofakeit.FirstName() },
	"OrderID":          func() interface{} { return gofakeit.UUID() },
	"PhoneNumber":      func() interface{} { return gofakeit.Phone() },
	"PostalCode":       func() interface{} { return gofakeit.Zip() },
	"ProductCode":      func() interface{} { return gofakeit.UUID() },
	"Quantity":         func() interface{} { return gofakeit.Number },
	"Region":           func() interface{} { return gofakeit.State() },
	"SSN":              func() interface{} { return gofakeit.SSN() },
	"Salary":           func() interface{} { return gofakeit.Float64Range },
	"SKU":              func() interface{} { return gofakeit.UUID() },
	"State":            func() interface{} { return gofakeit.State() },
	"StreetAddress":    func() interface{} { return gofakeit.Address() },
	"TaxID":            func() interface{} { return gofakeit.SSN() },
	"TrackingNumber":   func() interface{} { return gofakeit.UUID() },
	"TransactionID":    func() interface{} { return gofakeit.UUID() },
	"UnitPrice":        func() interface{} { return gofakeit.Float64Range },
	"URL":              func() interface{} { return gofakeit.URL() },
	"Weight":           func() interface{} { return gofakeit.Float64Range },
	"Website":          func() interface{} { return gofakeit.URL() },
	"ZIPCode":          func() interface{} { return gofakeit.Zip() },
}
