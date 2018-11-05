package response

import "github.com/nirmalvp/amadeusgo/api/request"

// Amadeus response represents the common element that
// should be present in every response returned to
// the user from the SDK
type AmadeusResponse struct {
	StatusCode int
	Request    request.AmadeusRequestData
	Parsed     bool
}
