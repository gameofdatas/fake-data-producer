package generators

import (
	"github.com/brianvoe/gofakeit/v6"
)

type FieldGenerator func() interface{}

var FieldGenerators = map[string]FieldGenerator{
	"Address":          func() interface{} { return gofakeit.Address().Address },
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
	"Region":           func() interface{} { return gofakeit.State() },
	"SSN":              func() interface{} { return gofakeit.SSN() },
	"SKU":              func() interface{} { return gofakeit.UUID() },
	"State":            func() interface{} { return gofakeit.State() },
	"StreetAddress":    func() interface{} { return gofakeit.Address() },
	"TaxID":            func() interface{} { return gofakeit.SSN() },
	"TrackingNumber":   func() interface{} { return gofakeit.UUID() },
	"TransactionID":    func() interface{} { return gofakeit.UUID() },
	"URL":              func() interface{} { return gofakeit.URL() },
	"Website":          func() interface{} { return gofakeit.URL() },
	"ZIPCode":          func() interface{} { return gofakeit.Zip() },
}
