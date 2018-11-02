package response

import "github.com/nirmalvp/amadeusgo/api/request"

type AmadeusResponse struct {
	StatusCode int
	Request    request.AmadeusRequestData
	Parsed     bool
}
