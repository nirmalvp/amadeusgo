package response

import (
	"encoding/json"
	"github.com/nirmalvp/amadeusgo/api/request"
)

// Structure of the Data part of the AMadeus API response.
type AirportData struct {
	Type           string
	SubType        string
	Name           string
	DetailedName   string
	TimeZoneOffset string
	IataCode       string
	GeoCode        struct {
		Latitude  float64
		Longitude float64
	}
	Address struct {
		CityName    string
		CityCode    string
		CountryName string
		CountryCode string
		RegionCode  string
	}
	Distance struct {
		Value float64
		Unit  string
	}
	Analytics struct {
		Flights struct {
			Score int
		}
		Travelers struct {
			Score int
		}
	}
	Relevance float64
}

// The response that the Rest API returns
type AirportsRest struct {
	Meta Meta
	Data []AirportData
}

// Final Result format as per the SDK spec that the user expects
type Airports struct {
	AmadeusResponse
	Result AirportsRest
	Data   []AirportData
}

func NewAirportsResponse(statusCode int, responseBody []byte, request request.AmadeusRequestData) Airports {
	var formatedRestResponse AirportsRest
	parseError := json.Unmarshal(responseBody, &formatedRestResponse)
	return Airports{
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
