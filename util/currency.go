package util

import (
	"reflect"
	"strings"
)

type currencies struct {
	EUR        string
	GBP        string
	USD        string
	CAD        string
	NOK        string
	DKK        string
	YEN        string
	INR        string
	BTC        string
	Currencies map[string]string
}

func (c currencies) ExistsCurrency(s string) bool {
	s = strings.ToUpper(s)

	val := reflect.Indirect(reflect.ValueOf(c)).Type()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Name == s {
			return true
		}
	}

	return false
}

var Currency = currencies{
	EUR: "€",
	GBP: "£",
	USD: "$",
	CAD: "$",
	NOK: "kr",
	DKK: "kr",
	YEN: "¥",
	INR: "₹",
	BTC: "₿",
}

func init() {
	Currency.Currencies = make(map[string]string)

	val := reflect.ValueOf(Currency)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Type().Name() != "string" {
			continue
		}

		Currency.Currencies[val.Type().Field(i).Name] = val.Field(i).Interface().(string)
	}
}
