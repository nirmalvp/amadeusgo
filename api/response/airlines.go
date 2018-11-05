package response

import (
	"encoding/json"
	"github.com/nirmalvp/amadeusgo/api/request"
)

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

func NewAirlineResponse(statusCode int, responseBody []byte, request request.AmadeusRequestData) Airlines {
	var formatedRestResponse AirlineRest
	parseError := json.Unmarshal(responseBody, &formatedRestResponse)
	return Airlines{
		AmadeusResponse: AmadeusResponse{
			StatusCode: statusCode,
			Body:       string(responseBody),
			Request:    request,
			Parsed:     parseError == nil,
		},
		Result: formatedRestResponse,
		Data:   formatedRestResponse.Data,
	}
}
