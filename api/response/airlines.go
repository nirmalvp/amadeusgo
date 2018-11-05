package response

import "github.com/nirmalvp/amadeusgo/api/request"

// Structure of the Data part of the AMadeus API response.
type AirlineData struct {
	Type         string
	IataCode     string
	BusinessName string
	CommonName   string
}

// The response that the Rest API returns
type AirlineRest struct {
	Meta Meta
	Data []AirlineData
}

// Final Result format as per the SDK spec that the user expects
type Airlines struct {
	AmadeusResponse
	Result AirlineRest
	Data   []AirlineData
}

func NewAirlineResponse(statusCode int, airlineRestResp AirlineRest, request request.AmadeusRequestData, isParsed bool) Airlines {
	return Airlines{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Request:    request,
			Parsed:     isParsed,
		},
		Result: airlineRestResp,
		Data:   airlineRestResp.Data,
	}

}
