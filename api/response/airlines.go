package response

import "github.com/nirmalvp/amadeusgo/api/request"

type AirlineData struct {
	Type         string
	IataCode     string
	BusinessName string
	CommonName   string
}

type AirlineRest struct {
	Meta Meta
	Data []AirlineData
}

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
