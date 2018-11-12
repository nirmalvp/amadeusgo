package request

import (
	"github.com/nirmalvp/amadeusgo/api/params"
)

type Verb string
type Header map[string]string

const (
	GET  Verb = "GET"
	POST Verb = "POST"
)

// Amadeus Request data contains all the required information for the amadeus
// Rest client to make an HTTP request.
type AmadeusRequestData struct {
	Verb        Verb
	URI         string
	Params      params.Params
	BearerToken *string
	Headers     Header
}
